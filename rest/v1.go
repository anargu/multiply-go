package restserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"multiply"
	"net/http"
	"strconv"
)

type Response struct {
	Data  *float32 `json:"data"`
	Error *string  `json:"error"`
}

func extractValues(r *http.Request) (*float32, *float32, error) {
	xStr := r.URL.Query().Get("x")
	yStr := r.URL.Query().Get("y")
	if xStr == "" || yStr == "" {
		return nil, nil, errors.New("incorrect number params")
	}
	x64, err := strconv.ParseFloat(xStr, 32)
	if err != nil {
		return nil, nil, err
	}
	y64, err := strconv.ParseFloat(yStr, 32)
	if err != nil {
		return nil, nil, err
	}
	x := float32(x64)
	y := float32(y64)

	return &x, &y, nil
}

func MultiplyRestHandlerV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	x, y, err := extractValues(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errormessage := fmt.Sprintf("Incorrect parameters setted. Error: %v", err.Error())
		json.NewEncoder(w).Encode(Response{
			Data:  nil,
			Error: &errormessage,
		})
		return
	}
	z, err := multiply.Multiply(x, y)
	response := Response{
		Data:  z,
		Error: multiply.GetErrorMessage(err),
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(response)
}
