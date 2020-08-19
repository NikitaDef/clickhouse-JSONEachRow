package main

import (
	"github.com/NikitaDef/clickhouse-JSONEachRow"
	"log"
	"net/http"
)

// http://localhost:8123/?query=select * from logs FORMAT JSONEachRow
func getValue(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://localhost:8123/?query=select%20*%20from%20logs%20FORMAT%20JSONEachRow")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if res.StatusCode != 200 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer res.Body.Close()
	_, err = clickhouseJSONEachRow.Copy(w, res.Body, 300)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/get", getValue)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
