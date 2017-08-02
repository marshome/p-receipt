package google_vision

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/marshome/x/filesystem"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

type Cache struct {
	Dir string
}

func NewCache(dir string) *Cache {
	return &Cache{Dir: dir}
}

func (c *Cache) getPath(fileName string) (filePath string) {
	fileName = strings.Replace(fileName, "/", "-", -1)
	fileName = strings.Replace(fileName, "\\", "_", -1)

	return c.Dir + "/" + fileName + ".json"
}

func (c *Cache) CalcFileName(imageBase64 string) (fileName string) {
	sig := sha256.Sum256([]byte(imageBase64))
	return base64.StdEncoding.EncodeToString(sig[:])
}

func (c *Cache) Load(fileName string) (data []byte) {
	logrus.WithField("_method_", "Load").WithField("fileName", fileName).Infoln()

	if c.Dir == "" {
		return nil
	}

	data, err := ioutil.ReadFile(c.getPath(fileName))
	if err != nil {
		logrus.Warningln(errors.WithMessage(err, "Load failed"))
		return nil
	}

	return data
}

func (c *Cache) Save(fileName string, data []byte) error {
	logrus.WithField("_method_", "Save").WithField("fileName", fileName).Infoln()
	if c.Dir == "" {
		return nil
	}

	return filesystem.NewFile(c.getPath(fileName), data)
}

func (c *Cache) Remove(fileName string) (err error) {
	logrus.WithField("_method_", "Remove").WithField("fileName", fileName).Infoln()
	if c.Dir == "" {
		return nil
	}

	return errors.Wrap(os.Remove(c.getPath(fileName)), "cache remove failed")
}
