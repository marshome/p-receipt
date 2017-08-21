package handler

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/marshome/p-receipt/api/receipt/server/models"
	"github.com/marshome/p-receipt/api/receipt/server/restapi/operations"
	"github.com/marshome/p-receipt/services/receipt"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

type Options struct {
	ProxyUrl     string
	CacheDir     string
	GoogleApiKey string
}

type ReceiptExtractHandler struct {
	logger         *zap.Logger
	Options        *Options
	ReceiptService *receipt.Service
}

func NewReceiptExtractHandler() *ReceiptExtractHandler {
	h := &ReceiptExtractHandler{}
	h.logger = zap.L().Named(reflect.TypeOf(*h).String())

	return h
}

func (h *ReceiptExtractHandler) Init(options *Options) (err error) {
	h.logger.Info("Init", zap.Any("options", options))

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
		return operations.NewReceiptsExtractBadRequest().WithPayload("请求数据错误")
	}

	result, err := h.ReceiptService.Extract(*params.Body.Image.ContentBase64)
	if err != nil {
		return operations.NewReceiptsExtractInternalServerError().WithPayload(strings.Replace(err.Error(), "DqrNp9RdNECDM2LIszWwp", "", -1))
	}

	response := &models.ReceiptExtractResponse{}
	response.ReceiptInfo = fromReceiptInfo(result.ReceiptInfo)
	response.FullTextAnnotation = fromTextAnnotation(result.TextAnnotation)

	return operations.NewReceiptsExtractOK().
		WithAccessControlAllowOrigin("*").
		WithPayload(response)
}

func (h *ReceiptExtractHandler) Report(params operations.ReceiptsReportParams) middleware.Responder {
	return nil
}
