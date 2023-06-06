package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseForm(dst interface{}, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return parseBody(dst, r.Body)
}

func parseBody(dst interface{}, b io.ReadCloser) error {

	body, err := io.ReadAll(b)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &dst)
	if err != nil {
		return err
	}

	return nil
}

func httpResponse(w http.ResponseWriter, r *http.Request, data interface{}, status int) error {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	logger.Infof("[HANDLER] %v %v Request Processed Successfully", r.Method, r.URL)

	return nil
}
