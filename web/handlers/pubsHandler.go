package handlers

import "encoding/json"
import "net/http"

import "github.com/m4tty/pubcloud/domain/pubs"

import ()

func GetPubsHandler(w http.ResponseWriter, r *http.Request) {
	dataTransformer := pubsDomain.ODCDataTransformer{}
	rawGetter := pubsDomain.HTTPRawGetter{}
	dataRetriever := pubsDomain.NewODCPubDataRetriever(rawGetter, dataTransformer)

	result, _ := dataRetriever.GetPubData()
	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}
