package main

import (
	"bytes"
	"net/http"
	"strings"
)

func static(r *http.Request, w http.ResponseWriter) {
	r.ParseForm()
	uri := strings.Replace(r.RequestURI, "", "", -1)
	data, err := Asset("static" + uri)
	if err != nil {
		return
	} else {
		http.ServeContent(w, r, uri, getTimeStamp(), bytes.NewReader(data))
	}
}
