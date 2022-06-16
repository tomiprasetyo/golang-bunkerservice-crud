package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// unmarshal json yang diterima dari request agar bisa digunakan di controller
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
