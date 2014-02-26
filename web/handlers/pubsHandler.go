package handlers

import "encoding/json"
import "net/http"

import "github.com/m4tty/pubcloud/web/resources"

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func GetPubsHandler(w http.ResponseWriter, r *http.Request) {
	//	dataMgr := pubsDomain.NewPubsMgr(dataManager)
	// result, err := dataMgr.GetPubs()

	// if err != nil {
	// 	return
	// }
	//result := new(resources.PubsResource)
	result := GetData()
	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

func GetData() interface{} {
	//pubs := make([]resources.PubsResource, 3000)
	var reader io.Reader
	var fileName = "./pubs.cache.csv"
	if isCached, _ := exists(fileName); isCached == true {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			os.Exit(1)
		}
		reader = bufio.NewReader(f)
		pubs := Format(reader)
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
			reader2 := io.TeeReader(response.Body, out)
			pubs := Format(reader2)
			io.Copy(out, response.Body)
			return pubs
		}
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
func Format(r io.Reader) interface{} {
	reader := csv.NewReader(r)
	lines, _ := reader.ReadAll()

	pubs := make([]resources.PubsResource, len(lines))

	for i, line := range lines {
		if i != 0 {
			pub := resources.PubsResource{
				Name: line[3],
			}
			pubs[i-1] = pub
		}
	}
	return pubs
}

//PubDataGetter
// - PubDataGet

//PubDataGetter implementation depends on Formatter, formatter takes a IOReader

//Formatter
// -

//func NewFormatter(r io.Reader) *Formatter

//dec := json.NewFormatter(os.Stdin)

//dec.Decode(v interface{}) error
