package router

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func BindJsonBody(rawBody io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(rawBody)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &v); err != nil {
		return err
	}

	return nil
}

func WriteJson(w http.ResponseWriter, v interface{}) {
	res, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
