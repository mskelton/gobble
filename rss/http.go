package rss

import (
	"encoding/xml"
	"io"
	"net/http"
)

func GetXml(uri string, v any) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(body, &v)
	return err
}
