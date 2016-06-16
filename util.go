package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func AssetHtml(fn string) ([]byte, error) {
	return Asset("static/html/" + fn + ".html")
}

func AssetImages(fn string) ([]byte, error) {
	return Asset("static/images/" + fn)
}

func AssetJs(fn string) ([]byte, error) {
	return Asset("static/js/" + fn + ".js")
}

func AssetCss(fn string) ([]byte, error) {
	return Asset("static/css/" + fn + ".css")
}

func parseForm(r *http.Request) {
	if strings.EqualFold(http.MethodPost, r.Method) {
		body, _ := ioutil.ReadAll(r.Body)
		var data map[string]string
		json.Unmarshal(body, &data)
		for k, v := range data {
			r.Form.Set(k, v)
		}
	}
}
