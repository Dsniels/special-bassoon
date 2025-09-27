package utils

import (
	"encoding/json"
	"fmt"
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


func ToStruct[T any]( r *http.Request, s *T){ 
	val := reflect.ValueOf(s)
	el := val.Elem()
	t:= val.Type()
	for i:= 0; i < t.NumField() ; i++ { 
		field:= t.Field(i)
		value := r.FormValue(field.Name)
		if el.Field(i).CanSet() && el.Field(i).Kind() != reflect.Uint{
			el.Field(i).SetString(value)
		}
	}

}