package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const MAX_BODY_SIZE = 1 << 20

func BindJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	if r.Header.Get("Content-Type") != "" &&
		r.Header.Get("Content-Type") != "application/json" &&
		r.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		return errors.New("content type must be application/json")
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_BODY_SIZE)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("request body is empty")
		}
		var syntaxErr *json.SyntaxError
		if errors.As(err, &syntaxErr) {
			return errors.New("malformed JSON at position")
		}
		var unmarshalTypeErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeErr) {
			return errors.New("invalid type for field: " + unmarshalTypeErr.Field)
		}
		if strings.HasPrefix(err.Error(), "json: unknown field") {
			return errors.New("unknown field in request")
		}
		if errors.Is(err, io.ErrUnexpectedEOF) {
			return errors.New("badly formed JSON")
		}
		if err == io.EOF {
			return errors.New("empty body")
		}
		return err
	}

	if dec.More() {
		return errors.New("only one JSON object allowed")
	}

	return nil
}
