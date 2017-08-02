package handler

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/marshome/p-vision/api/receipt/server/models"
	"github.com/marshome/p-vision/api/receipt/server/restapi/operations"
	"github.com/marshome/p-vision/services/receipt"
	"github.com/sirupsen/logrus"
)

type Options struct {
	ProxyUrl     string
	CacheDir     string
	GoogleApiKey string
}

type ReceiptExtractHandler struct {
	Options        *Options
	ReceiptService *receipt.Service
}

func NewReceiptExtractHandler() *ReceiptExtractHandler {
	return &ReceiptExtractHandler{}
}

func (h *ReceiptExtractHandler) Init(options *Options) (err error) {
	logrus.Infoln("ReceiptExtractHandler Init options=", options)

	if options == nil {
		return fmt.Errorf("options is nil")
	}

	h.Options = options

	h.ReceiptService = receipt.NewService()
	err = h.ReceiptService.Init(&receipt.Options{
		ProxyUrl:     h.Options.ProxyUrl,
		CacheDir:     h.Options.CacheDir,
		GoogleApiKey: h.Options.GoogleApiKey,
	})

	if err != nil {
		return err
	}

	return nil
}

func (h *ReceiptExtractHandler) Extract(params operations.ReceiptsExtractParams) middleware.Responder {
	if params.Body == nil || params.Body.Image == nil || params.Body.Image.ContentBase64 == nil {
		logrus.WithField("_func_", "ReceiptExtractHandler.Extract").Errorln("params failed", params)
		return operations.NewReceiptsExtractBadRequest()
	}

	result, err := h.ReceiptService.Extract(*params.Body.Image.ContentBase64)
	if err != nil {
		logrus.WithField("func", "ReceiptExtractHandler.Extract").Errorln(err)
		return operations.NewReceiptsExtractBadRequest()
	}

	response := &models.ReceiptExtractResponse{}
	response.ReceiptInfo = fromReceiptInfo(result.ReceiptInfo)
	response.FullTextAnnotation = fromTextAnnotation(result.TextAnnotation)

	return operations.NewReceiptsExtractOK().WithPayload(response)
}

func (h *ReceiptExtractHandler) Report(params operations.ReceiptsReportParams) middleware.Responder {
	return nil
}
