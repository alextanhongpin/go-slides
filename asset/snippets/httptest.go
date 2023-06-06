package main

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

func Dump(w *http.Response, r *http.Request) (string, error) {
	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		return "", err
	}

	res, err := httputil.DumpResponse(w, true)
	if err != nil {
		return "", err
	}

	output := make([]string, 3)
	output[0] = string(req)
	output[2] = string(res)

	return strings.Join(output, "\n"), nil
}
