package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Response map[string]any

func WriteResponse(w http.ResponseWriter, statusCode int, data Response) error {
	dataResponse, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	dataResponse = append(dataResponse, '\n')
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataResponse)
	return nil

}

func GetIdFromParams(r *http.Request) (uint, error) {
	param := chi.URLParam(r, "id")
	if param == "" {
		return 0, fmt.Errorf("must pass a valid Id")
	}
	parsed, err := strconv.ParseUint(param, 10, 10)
	if err != nil {
		return 0, err
	}
	return uint(parsed), nil

}
