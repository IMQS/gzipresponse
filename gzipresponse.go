// Package gzipresponse is a tiny helper to gzip an HTTP response
package gzipresponse

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// Use this as a replacement for w.Write()
func Write(w http.ResponseWriter, r *http.Request, data []byte) (int, error) {
	if strings.Index(r.Header.Get("Accept-Encoding"), "gzip") == -1 {
		return w.Write(data)
	}

	if _, ok := w.Header()["Content-Type"]; !ok {
		// If content type is not set, infer it from the uncompressed body.
		// The net/http package normally does this, but that wasn't written to
		// work on compressed data.
		w.Header().Set("Content-Type", http.DetectContentType(data))
	}

	w.Header().Set("Content-Encoding", "gzip")

	zipper := gzip.NewWriter(w)
	bytes, err := zipper.Write(data)
	errClose := zipper.Close()
	if err != nil {
		return bytes, err
	}
	return bytes, errClose
}
