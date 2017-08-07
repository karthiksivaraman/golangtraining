package main

import (
	"strconv"
	"log"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/*
{  "disposition": "MALICIOUS",

    "malwareName": "string",

    "ampScore": 0,

    "txId": "string" }
 */

type AMPVerdict struct {
	Disposition string  `json:"disposition"`
	MalwareName string  `json:"malwareName"`
	AmpScore    int32   `json:"ampScore"`
	Txid        string  `json:"txId"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/{oid}/scan/{txId}", SimulateScan).Methods("POST")
	log.Fatal(http.ListenAndServeTLS(":4433", "server.crt", "server.key", router))
}

func SimulateScan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Request-Size", strconv.FormatInt(r.ContentLength, 10))
	w.Header().Set("X-Request-Type", r.Header.Get("Content-type"))
	w.Header().Set("X-Organization ID", vars["oid"])
	ampVerdict := AMPVerdict{"malware", "Eicar", 0, vars["txId"]}
	encodedResponse, error := json.Marshal(ampVerdict)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(encodedResponse)
}
