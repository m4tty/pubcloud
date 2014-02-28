package pubsDomain

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPRawGetter struct {
}

func (getter HTTPRawGetter) RawDataGet(path string) (io.ReadCloser, error) {
	//"http://data.denvergov.org/download/gis/liquor_licenses/csv/liquor_licenses.csv"
	fmt.Print("Going out to get the pub data...")
	response, err := http.Get(path)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	} else {

		fmt.Print("Got the pub data...")
		return response.Body, nil
	}
}
