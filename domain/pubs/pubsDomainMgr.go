package pubsDomain

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/m4tty/pubcloud/web/resources"
	"io"
	"net/http"
	"os"
)

type PubDataRetriever interface {
	GetPubData(transformer PubDataTransformer) []resources.Pubs
}

type PubDataTransformer interface {
	PubDataTransform(v interface{}) ([]resources.Pubs, error)
}

type ODCPubDataRetriever struct {
}

type ODCDataTransformer struct {
}

func (retriever ODCPubDataRetriever) GetPubData(transformer PubDataTransformer) []resources.Pubs {
	//pubs := make([]resources.Pubs, 3000)
	var reader io.Reader
	var fileName = "./pubs.cache.csv"
	if isCached, _ := exists(fileName); isCached == true {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			os.Exit(1)
		}
		reader = bufio.NewReader(f)
		pubs, _ := transformer.PubDataTransform(reader)
		return pubs
	} else {
		fmt.Print("Going out to get the pub data...")
		response, err := http.Get("http://data.denvergov.org/download/gis/liquor_licenses/csv/liquor_licenses.csv")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
			return nil
		} else {
			defer response.Body.Close()
			out, _ := os.Create("pubs.cache.csv")
			defer out.Close()
			fmt.Print("Got the pub data...")
			reader2 := io.TeeReader(response.Body, out)
			pubs, _ := transformer.PubDataTransform(reader2)
			io.Copy(out, response.Body)
			return pubs
		}
	}

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
