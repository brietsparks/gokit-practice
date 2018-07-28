package abilities

import (
	"context"
	"net/http"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/go-kit/kit/log"
)

func MakeHTTPHandler(s Service, l log.Logger) http.Handler {
	r := mux.NewRouter()

	getAbilitiesByOwnerIdEndpoint := MakeGetAbilitiesByOwnerIdEndpoint(s)
	getAbilitiesByOwnerIdEndpoint = loggingMiddleware(log.With(l, "foo", "bar"))(getAbilitiesByOwnerIdEndpoint)

	createAbilityEndpoint := MakeCreateAbilityEndpoint(s)
	updateAbilityEndpoint := MakeUpdateAbilityEndpoint(s)
	deleteAbilityEndpoint := MakeDeleteAbilityEndpoint(s)

	r.Methods("GET").Path("/owner/{ownerId}/abilities").Handler(httptransport.NewServer(
		getAbilitiesByOwnerIdEndpoint,
		decodeAbilitiesReadRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/abilities").Handler(httptransport.NewServer(
		createAbilityEndpoint,
		decodeAbilityWriteRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/abilities").Handler(httptransport.NewServer(
		updateAbilityEndpoint,
		decodeAbilityWriteRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/abilities").Handler(httptransport.NewServer(
		deleteAbilityEndpoint,
		decodeAbilityWriteRequest,
		encodeResponse,
	))

	return r
}

func decodeAbilityWriteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req abilityWriteRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Ability); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeAbilitiesReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return abilitiesReadRequest{
		OwnerId: mux.Vars(r)["ownerId"],
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
