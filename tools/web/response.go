package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// RespondJSON converts a Go value to JSON and sends it to the client.
// If v is nil or code is equal to http.StatusNoContent we avoid writing any content to w.
// HTTP response header with the provided status code is always set.
func RespondJSON(w http.ResponseWriter, v interface{}, code int) error {
	// According to https://tools.ietf.org/search/rfc2616#section-7.2.1:
	//
	// "Any HTTP/1.1 message containing an entity-body SHOULD include a Content-Type
	// header field defining the media type of that body"
	//
	// Since there is no content, then there is no reason to specify a Content-Type header
	if code == http.StatusNoContent || v == nil {
		w.WriteHeader(code)
		return nil
	}

	var jsonData []byte

	var err error
	switch v := v.(type) {
	case []byte:
		jsonData = v
	case io.Reader:
		jsonData, err = ioutil.ReadAll(v)
	default:
		jsonData, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}

	// Set the content type.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response and context.
	w.WriteHeader(code)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
