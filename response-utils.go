package swiftgo

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
)

func (s *SwiftGO) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single json value")
	}

	return nil
}

// WriteJSON writes json from arbitrary data
func (s *SwiftGO) WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// WriteXML writes xml from arbitrary data
func (s *SwiftGO) WriteXML(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := xml.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// DownloadFile downloads a file
func (s *SwiftGO) DownloadFile(w http.ResponseWriter, r *http.Request, pathToFile, fileName string) error {
	fp := path.Join(pathToFile, fileName)
	fileToServe := filepath.Clean(fp)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; file=\"%s\"", fileName))
	http.ServeFile(w, r, fileToServe)
	return nil
}

// Error404 returns page not found response
func (s *SwiftGO) Error404(w http.ResponseWriter, r *http.Request) {
	s.ErrorStatus(w, http.StatusNotFound)
}

// Error500 returns internal server error response
func (s *SwiftGO) Error500(w http.ResponseWriter, r *http.Request) {
	s.ErrorStatus(w, http.StatusInternalServerError)
}

// ErrorUnauthorized sends an unauthorized status (client is not known)
func (s *SwiftGO) ErrorUnauthorized(w http.ResponseWriter, r *http.Request) {
	s.ErrorStatus(w, http.StatusUnauthorized)
}

// ErrorForbidden returns a forbidden status message (client is known)
func (s *SwiftGO) ErrorForbidden(w http.ResponseWriter, r *http.Request) {
	s.ErrorStatus(w, http.StatusForbidden)
}

// ErrorStatus returns a response with the supplied http status
func (s *SwiftGO) ErrorStatus(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}