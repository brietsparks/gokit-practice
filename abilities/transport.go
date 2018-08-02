package abilities

import (
	"context"
	"net/http"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/go-kit/kit/log"
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"gokit-practice/util"
	"github.com/go-kit/kit/endpoint"
)

func MakeHTTPHandler(
	s Service,
	logger log.Logger,
	jwtChecker *jwtmiddleware.JWTMiddleware,
) http.Handler {
	r := mux.NewRouter()

	getAbilitiesByOwnerIdEndpoint := MakeGetAbilitiesByOwnerIdEndpoint(s)
	getAbilitiesByOwnerIdEndpoint = loggingMiddleware(log.With(logger))(getAbilitiesByOwnerIdEndpoint, "getAbilitiesByOwnerIdEndpoint")

	createAbilityEndpoint := MakeCreateAbilityEndpoint(s)
	createAbilityEndpoint = loggingMiddleware(log.With(logger))(createAbilityEndpoint, "createAbilityEndpoint")

	updateAbilityEndpoint := MakeUpdateAbilityEndpoint(s)
	deleteAbilityEndpoint := MakeDeleteAbilityEndpoint(s)

	getAbilitiesByOwnerIdHandler := httptransport.NewServer(
		getAbilitiesByOwnerIdEndpoint,
		decodeAbilitiesReadRequest,
		util.EncodeResponse,
	)

	makeWriteHandler := func(endpoint endpoint.Endpoint) http.Handler {
		writeHandler := httptransport.NewServer(endpoint, decodeAbilityWriteRequest, util.EncodeResponse)
		return negroni.New(
			negroni.HandlerFunc(jwtChecker.HandlerWithNext),
			negroni.Wrap(writeHandler),
		)
	}

	createAbilityHandler := makeWriteHandler(createAbilityEndpoint)
	updateAbilityHandler := makeWriteHandler(updateAbilityEndpoint)
	deleteAbilityHandler := makeWriteHandler(deleteAbilityEndpoint)

	r.Methods("GET").Path("/owner/{ownerId}/abilities").Handler(getAbilitiesByOwnerIdHandler)
	r.Methods("POST").Path("/abilities").Handler(createAbilityHandler)
	r.Methods("PUT").Path("/abilities").Handler(updateAbilityHandler)
	r.Methods("DELETE").Path("/abilities").Handler(deleteAbilityHandler)

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
