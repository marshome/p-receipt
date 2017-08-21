package google_vision

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/marshome/x/filesystem"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type Cache struct {
	logger *zap.Logger
	Dir    string
}

func NewCache(dir string) *Cache {
	c := &Cache{
		Dir: dir,
	}
	c.logger = zap.L().Named(reflect.TypeOf(*c).Name())

	return c
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
	c.logger.Info("Load", zap.String("FileName", fileName))

	if c.Dir == "" {
		return nil
	}

	data, err := ioutil.ReadFile(c.getPath(fileName))
	if err != nil {
		c.logger.Info("Load failed", zap.Error(err))
		return nil
	}

	return data
}

func (c *Cache) Save(fileName string, data []byte) error {
	c.logger.Info("Save", zap.String("FileName", fileName))
	if c.Dir == "" {
		return nil
	}

	return filesystem.NewFile(c.getPath(fileName), data)
}

func (c *Cache) Remove(fileName string) (err error) {
	c.logger.Info("Remove", zap.String("FileName", fileName))
	if c.Dir == "" {
		return nil
	}

	return errors.Wrap(os.Remove(c.getPath(fileName)), "cache remove failed")
}
