package google_vision

import (
	"encoding/json"
	"fmt"
	"github.com/marshome/p-vision/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/api/vision/v1"
	"net/http"
	"net/url"
	"reflect"
)

const MAX_TEXT_DETECT_RESULT_DEFAULT = 20

type Options struct {
	ProxyUrl            string
	MaxTextDetectResult int64
	CacheDir            string
	ApiKey              string
}

type CallOptionKey struct {
	ApiKey string
}

func (op *CallOptionKey) Get() (k, v string) {
	return "key", op.ApiKey
}

type Client struct {
	logger  *zap.Logger
	options *Options
	service *vision.Service
	cache   *Cache
}

func NewClient() *Client {
	c := &Client{}
	c.logger = zap.L().Named(reflect.TypeOf(*c).Name())

	return c
}

func (c *Client) Init(options *Options) (err error) {
	c.logger.Info("Init", zap.Any("options", options))

	if options == nil {
		return fmt.Errorf("invalid param: opts")
	}

	c.options = options

	if c.options.ApiKey == "" {
		return errors.WithStack(errors.New("need google ApiKey"))
	}

	if c.options.MaxTextDetectResult == 0 {
		c.options.MaxTextDetectResult = MAX_TEXT_DETECT_RESULT_DEFAULT
	}

	c.cache = NewCache(c.options.CacheDir)

	var client *http.Client
	if options.ProxyUrl != "" {
		proxyUrl, err := url.Parse(options.ProxyUrl)
		if err != nil {
			return err
		}

		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	} else {
		client = &http.Client{}
	}

	c.service, err = vision.New(client)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CallGoogleVision(request *vision.BatchAnnotateImagesRequest) (response *vision.BatchAnnotateImagesResponse, err error) {
	response, err = c.service.Images.Annotate(request).Do(&CallOptionKey{ApiKey: c.options.ApiKey})
	if err != nil {
		return nil, errors.WithMessage(err, "Images.Annotate failed")
	}

	return response, nil
}

func (c *Client) cacheResponse(cacheFileName string, response *vision.BatchAnnotateImagesResponse) {
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		c.logger.Error("cacheResponse marshal", zap.Error(err))
	} else {
		err = c.cache.Save(cacheFileName, jsonData)
		if err != nil {
			c.logger.Error("cacheResponse save failed", zap.Error(err))
		}
	}
}

func (c *Client) TextDetection(imageBase64 string) (textAnnotation *models.TextAnnotation, err error) {
	cacheFileName := c.cache.CalcFileName(imageBase64)
	cacheData := c.cache.Load(cacheFileName)
	if cacheData != nil {
		response := &vision.BatchAnnotateImagesResponse{}
		err = json.Unmarshal(cacheData, response)
		if err != nil {
			c.logger.Error("TextDetection unmarshal cached data failed", zap.Error(err))
			err = c.cache.Remove(cacheFileName)
			if err != nil {
				c.logger.Error("TextDetection cache remove failed", zap.Error(err))
			}
		} else {
			if response.Responses[0].FullTextAnnotation == nil {
				return nil, fmt.Errorf("未能识别图片中的文字，请选择包含清晰文字的图片")
			}

			return fromTextAnnotation(response.Responses[0].FullTextAnnotation), nil
		}
	}

	request := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{
			{
				Image: &vision.Image{
					Content: imageBase64,
				},
				Features: []*vision.Feature{{
					Type:       "TEXT_DETECTION",
					MaxResults: c.options.MaxTextDetectResult,
				}},
			},
		},
	}
	response, err := c.CallGoogleVision(request)
	if err != nil {
		return nil, errors.WithMessage(err, "后端服务失败：TextDetection failed")
	}

	if response.Responses == nil || len(response.Responses) != 1 {
		return nil, fmt.Errorf("后端服务失败：Images.Annotate response.Responses count failed")
	}

	if response.Responses[0].FullTextAnnotation == nil {
		c.cacheResponse(cacheFileName, response)

		return nil, fmt.Errorf("未能识别图片中的文字，请选择包含清晰文字的图片")
	}

	c.cacheResponse(cacheFileName, response)

	return fromTextAnnotation(response.Responses[0].FullTextAnnotation), nil
}
