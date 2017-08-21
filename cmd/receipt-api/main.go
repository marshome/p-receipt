package main

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/marshome/p-vision/api/receipt/server/restapi"
	"github.com/marshome/p-vision/api/receipt/server/restapi/operations"
	"github.com/marshome/p-vision/cmd/receipt-api/handler"
	"github.com/marshome/x/httphelper"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
		return
	}
	zap.ReplaceGlobals(l)

	var bind_addr string

	var options = &handler.Options{
		CacheDir: "./google_vision_cache",
	}

	cmd := cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
			if err != nil {
				return errors.WithStack(err)
			}

			api := operations.NewMarsReceiptAPI(swaggerSpec)

			h := handler.NewReceiptExtractHandler()
			err = h.Init(options)
			if err != nil {
				return errors.WithStack(err)
			}

			api.ReceiptsReportHandler = operations.ReceiptsReportHandlerFunc(
				func(params operations.ReceiptsReportParams) middleware.Responder {
					return h.Report(params)
				})
			api.ReceiptsExtractHandler = operations.ReceiptsExtractHandlerFunc(
				func(params operations.ReceiptsExtractParams) middleware.Responder {
					return h.Extract(params)
				})

			zap.L().Info("Start server", zap.String("addr", bind_addr))
			err = http.ListenAndServe(bind_addr,
				httphelper.Recovery(cors.Default().Handler(api.Serve(nil))))
			if err != nil {
				return errors.WithStack(err)
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&bind_addr, "bind-addr", ":", "api server bind addr")
	cmd.PersistentFlags().StringVar(&options.GoogleApiKey, "google-api-key", "", "googleApiKey")
	cmd.PersistentFlags().StringVar(&options.ProxyUrl, "proxy-url", "", "proxy")

	err = cmd.Execute()
	if err != nil {
		zap.L().Fatal("Cmd execute failed ", zap.Error(err))
	}
}
