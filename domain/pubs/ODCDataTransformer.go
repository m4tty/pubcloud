package pubsDomain

import (
	"encoding/csv"
	"errors"
	"github.com/m4tty/pubcloud/web/resources"
	"io"
)

type ODCDataTransformer struct {
}

func (transformer ODCDataTransformer) PubDataTransform(r interface{}) ([]resources.Pubs, error) {
	if v, ok := r.(io.Reader); ok {
		reader := csv.NewReader(v)
		lines, _ := reader.ReadAll()

		pubs := make([]resources.Pubs, len(lines))

		for i, line := range lines {
			if i != 0 {
				pub := resources.Pubs{
					Name: line[3],
				}
				pubs[i-1] = pub
			}
		}
		return pubs, nil
	} else {
		return nil, errors.New("transform: input data type is not supported")
	}

}
