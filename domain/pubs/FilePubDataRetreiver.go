package pubsDomain

import (
	"bufio"

	"fmt"
	"github.com/m4tty/pubcloud/web/resources"
	"io"
	"os"
)

type FilePubDataRetriever struct {
	transformer PubDataTransformer
	rawGetter   RawDataGetter
}

func (retriever FilePubDataRetriever) GetPubData(transformer PubDataTransformer) []resources.Pubs {
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
		fmt.Print("File is missing.")
		return nil
	}

}
