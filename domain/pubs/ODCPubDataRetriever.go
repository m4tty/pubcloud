package pubsDomain

import (
	"fmt"
	"github.com/m4tty/pubcloud/web/resources"
	"io"
	"os"
)

type ODCPubDataRetriever struct {
	transformer PubDataTransformer
	rawGetter   RawDataGetter
}

func NewODCPubDataRetriever(raw RawDataGetter, trans PubDataTransformer) *ODCPubDataRetriever {
	return &ODCPubDataRetriever{trans, raw}
}
func (retriever ODCPubDataRetriever) GetPubData() ([]resources.Pubs, error) {

	var location = "http://data.denvergov.org/download/gis/liquor_licenses/csv/liquor_licenses.csv"
	reader, err := retriever.rawGetter.RawDataGet(location)
	defer reader.Close()
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	} else {
		out, _ := os.Create("pubs.cache.csv")
		defer out.Close()
		reader2 := io.TeeReader(reader, out)
		pubs, _ := retriever.transformer.PubDataTransform(reader)
		io.Copy(out, reader2)
		return pubs, nil
	}

}
