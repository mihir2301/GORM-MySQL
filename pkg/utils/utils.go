package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	errs := json.Unmarshal([]byte(body), x)
	if errs != nil {
		return
	}
}
