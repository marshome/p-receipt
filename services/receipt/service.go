package receipt

import (
	"fmt"
	"github.com/marshome/p-receipt/backends/google_vision"
	"github.com/marshome/p-receipt/models"
	"go.uber.org/zap"
	"reflect"
)

type Options struct {
	ProxyUrl     string
	CacheDir     string
	GoogleApiKey string
}

type Service struct {
	logger             *zap.Logger
	options            *Options
	googleVisionClient *google_vision.Client
}

func NewService() *Service {
	s := &Service{}
	s.logger = zap.L().Named(reflect.TypeOf(*s).String())

	return s
}

func (s *Service) Init(options *Options) (err error) {
	s.logger.Info("Receipt service init", zap.Any("options", options))

	if options == nil {
		return fmt.Errorf("options is nil")
	}

	s.options = options

	s.googleVisionClient = google_vision.NewClient()
	err = s.googleVisionClient.Init(&google_vision.Options{
		ProxyUrl:            options.ProxyUrl,
		MaxTextDetectResult: 20,
		CacheDir:            s.options.CacheDir,
		ApiKey:              s.options.GoogleApiKey,
	})
	if err != nil {
		return
	}

	return nil
}

func (s *Service) Extract(imageBase64 string) (result *models.ReceiptResult, err error) {
	textAnnotation, err := s.googleVisionClient.TextDetection(imageBase64)
	if err != nil {
		return nil, err
	}

	ctx := NewExtractContext(textAnnotation)
	ctx.Process()

	return &models.ReceiptResult{
		TextAnnotation: textAnnotation,
		ReceiptInfo:    ctx.Result,
	}, nil
}
