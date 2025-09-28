package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
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

func ToStruct[T any](r *http.Request, g *T) error {

	val := reflect.ValueOf(g)
	el := val.Elem()
	t := el.Type()

	if val.Kind() != reflect.Pointer {
		return fmt.Errorf("must pass a pointer to an struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := r.FormValue(field.Name)
		fieldVal := el.Field(i)
		if !fieldVal.CanSet() {
			continue
		}
		if err := setField(field, fieldVal, value); err != nil {
			slog.Error("setting field", slog.Any("message", err))
			continue
		}
	}

	return nil
}

func setField(field reflect.StructField, fieldElem reflect.Value, value string) error {
	switch field.Type.Kind() {
	case reflect.String:
		fieldElem.SetString(value)
	case reflect.Uint:
		num, err := strconv.ParseUint(value, 10, 60)
		if err != nil {
			return err
		}
		fieldElem.SetUint(num)
	}
	return nil
}
