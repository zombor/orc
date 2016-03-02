package http

import (
	"encoding/json"
	"net/http"
)

func DecodeRegisterNodeRequest(r *http.Request) (interface{}, error) {
	var request registerNodeRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if request.Name == "" || request.Address == "" {
		var e errRegNodeFields
		if request.Name == "" {
			e.Name = "`name` is a required field"
		}
		if request.Address == "" {
			e.Address = "`address` is a required field"
		}
		return request, e
	}

	return request, err
}

func EncodeRegisterNodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type errRegNodeFields struct {
	Name, Address string
}

func (err errRegNodeFields) Error() string {
	var vndErr struct {
		Message  string `json:"message"`
		Embedded struct {
			Errors []vndError `json:"errors"`
		} `json:"_embedded"`
	}

	vndErr.Message = "Validation error"
	vndErr.Embedded.Errors = make([]vndError, 0)
	if err.Name != "" {
		vndErr.Embedded.Errors = append(vndErr.Embedded.Errors, vndError{
			Message: err.Name,
			LogRef:  1,
			Path:    "/name",
		})
	}
	if err.Address != "" {
		vndErr.Embedded.Errors = append(vndErr.Embedded.Errors, vndError{
			Message: err.Address,
			LogRef:  2,
			Path:    "/address",
		})
	}

	encErr, _ := json.Marshal(vndErr)
	return string(encErr)
}

type vndError struct {
	Message string      `json:"message"`
	LogRef  interface{} `json:"logref"`
	Path    string      `json:"path"`
}
