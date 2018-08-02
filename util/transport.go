package util

import (
	"encoding/json"
	"context"
	"net/http"
)

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
