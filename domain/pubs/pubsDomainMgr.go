package pubsDomain

import (
	"github.com/m4tty/pubcloud/web/resources"
	"io"
	"os"
)

type PubDataRetriever interface {
	GetPubData(transformer PubDataTransformer) ([]resources.Pubs, error)
}
type RawDataGetter interface {
	RawDataGet(path string) (io.ReadCloser, error)
}

type PubDataTransformer interface {
	PubDataTransform(v interface{}) ([]resources.Pubs, error)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
