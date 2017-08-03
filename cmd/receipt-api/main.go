package main

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/marshome/p-vision/api/receipt/server/restapi"
	"github.com/marshome/p-vision/api/receipt/server/restapi/operations"
	"github.com/marshome/p-vision/cmd/receipt-api/handler"
	"github.com/marshome/x/jsonhelper"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
)

func main() {
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
					logrus.WithField("_path_", "ReceiptsReport").WithField("params", jsonhelper.SimpleJson(params)).Infoln()
					response := h.Report(params)
					logrus.WithField("_path_", "ReceiptsReport").WithField("response", jsonhelper.SimpleJson(response)).Infoln()

					return response
				})
			api.ReceiptsExtractHandler = operations.ReceiptsExtractHandlerFunc(
				func(params operations.ReceiptsExtractParams) middleware.Responder {
					logrus.WithField("_path_", "ReceiptsExtract").WithField("params", jsonhelper.SimpleJson(params)).Infoln()
					response := h.Extract(params)
					//logrus.WithField("_path_", "ReceiptsExtract").WithField("response", jsonhelper.SimpleJson(response)).Infoln()

					return response
				})

			api.Logger = func(format string, args ...interface{}) {
				logrus.WithField("_api_", "MarsReceiptAPI").Infof(format, args)
			}

			addr := ":8080"
			logrus.Infoln("addr=", addr)

			err = http.ListenAndServe(addr, cors.Default().Handler(api.Serve(nil)))
			if err != nil {
				return errors.WithStack(err)
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&options.GoogleApiKey, "google-api-key", "", "googleApiKey")
	cmd.PersistentFlags().StringVar(&options.ProxyUrl, "proxy-url", "", "proxy")

	err := cmd.Execute()
	if err != nil {
		logrus.Fatalln(err)
	}
}
