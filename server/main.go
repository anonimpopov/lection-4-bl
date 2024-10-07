package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type requestBody struct {
	Test string `json:"test"`
}

func StartHandle(w http.ResponseWriter, r *http.Response) {
	b, _ := io.ReadAll(r.Body)

	var rb requestBody
	_ = json.Unmarshal(b, &rb)

	_, _ = w.Write([]byte("ok"))

}
