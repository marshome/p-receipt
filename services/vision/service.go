package vision

import (
	"google.golang.org/api/vision/v1"
	"net/http"
	"io/ioutil"
	"encoding/base64"
	"fmt"
	"net/url"
)

type VisionService struct {

}

func NewVisionService() *VisionService {
	return &VisionService{}
}

func (v *VisionService)Annotate(imageFile string) (err error) {
	proxyUrl, err := url.Parse("https://127.0.0.1:49916")
	if err != nil {
		return err
	}

	client := &http.Client{
		Transport:&http.Transport{
			Proxy:http.ProxyURL(proxyUrl),
		},
	}

	service, err := vision.New(client)
	if err != nil {
		return err
	}

	imageData, err := ioutil.ReadFile(imageFile)
	if err != nil {
		return err
	}

	request := &vision.BatchAnnotateImagesRequest{
		Requests:[]*vision.AnnotateImageRequest{
			{
				Image:&vision.Image{
					Content:base64.StdEncoding.EncodeToString(imageData),
				},
				Features:[]*vision.Feature{{
					Type:"TEXT_DETECTION",
					MaxResults:20,
				}},
			},
		},
	}
	call := service.Images.Annotate(request)
	response, err := call.Do()
	if err != nil {
		return err
	}

	if response.Responses != nil {
		for _, v := range response.Responses {
			fmt.Println(v)
		}
	}

	return nil
}
