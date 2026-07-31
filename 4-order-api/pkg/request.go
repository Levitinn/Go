package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ParseID(r *http.Request) (uint, error) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func DecodeJSON[T any](r *http.Request, dest *T) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return err
	}
	return validator.New().Struct(dest)
}
