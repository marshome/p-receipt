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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
)

func main() {
	var bind_addr string

	var options = &handler.Options{
		CacheDir: "./google_vision_cache",
	}

	cmd := cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			middleware.Debug = true

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
					panic(errors.New("aaaa"))
					return h.Extract(params)
				})

			api.Logger = func(format string, args ...interface{}) {
				logrus.WithField("_api_", "MarsReceiptAPI").Infof(format, args)
			}

			logrus.Infoln("addr=", bind_addr)

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

	err := cmd.Execute()
	if err != nil {
		logrus.Errorln(err)
	}
}
