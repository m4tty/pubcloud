package pubsDomain

import (
	"bufio"

	"fmt"
	"github.com/m4tty/pubcloud/web/resources"
	"io"
	"net/http"
	"os"
)

type ODCPubDataRetriever struct {
}

func (retriever ODCPubDataRetriever) GetPubData(transformer PubDataTransformer) []resources.Pubs {
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
