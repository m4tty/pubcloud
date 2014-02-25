package handlers

import "encoding/json"
import "net/http"

import "github.com/m4tty/pubcloud/web/resources"

func GetPubsHandler(w http.ResponseWriter, r *http.Request) {

	//	dataMgr := pubsDomain.NewPubsMgr(dataManager)
	// result, err := dataMgr.GetPubs()

	// if err != nil {
	// 	return
	// }
	result := new(resources.PubsResource)

	js, error := json.MarshalIndent(result, "", "  ")
	if error != nil {
		return
	}
	w.Write(js)
	return
}

//PubDataGetter
// - PubDataGet

//PubDataGetter implementation depends on Formatter, formatter takes a IOReader

//Formatter
// -

//func NewFormatter(r io.Reader) *Formatter

//dec := json.NewFormatter(os.Stdin)

//dec.Decode
