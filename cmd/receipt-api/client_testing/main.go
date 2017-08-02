package main

import (
	"encoding/base64"
	runtime "github.com/go-openapi/runtime/client"
	"github.com/marshome/p-vision/api/receipt/client/client"
	"github.com/marshome/p-vision/api/receipt/client/client/operations"
	"github.com/marshome/p-vision/api/receipt/client/models"
	"github.com/marshome/x/jsonhelper"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func processFile(filePath string) {
	logrus.Infoln("processFile")
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.WithField("_func_", "processFile").Errorln(err)
		return
	}

	transport := runtime.New("127.0.0.1:8080", client.DefaultBasePath, nil)
	transport.Debug = true

	c := client.New(transport, nil)

	imageBase64 := base64.StdEncoding.EncodeToString(imageData)
	request := operations.NewReceiptsExtractParams().WithBody(&models.ReceiptExtractRequest{
		Image: &models.Image{
			ContentBase64: &imageBase64,
		},
	})
	response, err := c.Operations.ReceiptsExtract(request)
	if err != nil {
		logrus.Errorln(errors.WithMessage(err, "ReceiptExtract"))
		return
	}

	if response == nil {
		logrus.WithField("_func_", "processFile").Errorln(err)
		return
	}

	logrus.WithField("_func_", "processFile").Infoln(jsonhelper.SimpleJson(response.Payload.ReceiptInfo))
}

func main() {
	processFile("./testdata/5.jpg")
	//processFile("./testdata/2.jpeg")
	//processFile("./testdata/3.jpg")
	//processFile("./testdata/4.jpg")
	//processFile("./testdata/5.jpg")
}
