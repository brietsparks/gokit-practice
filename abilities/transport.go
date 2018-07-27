package abilities

import (
	"context"
	"net/http"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	//"github.com/satori/go.uuid"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := mux.NewRouter()

	createAbilityEndpoint := MakeCreateAbilityEndpoint(s)
	updateAbilityEndpoint := MakeUpdateAbilityEndpoint(s)
	deleteAbilityEndpoint := MakeDeleteAbilityEndpoint(s)
	//getAbilityEndpoint := MakeGetAbilityEndpoint(s)

	r.Methods("POST").Path("/abilities").Handler(httptransport.NewServer(
		createAbilityEndpoint,
		decodeCreateAbilityRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/abilities").Handler(httptransport.NewServer(
		updateAbilityEndpoint,
		decodeUpdateAbilityRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/abilities").Handler(httptransport.NewServer(
		deleteAbilityEndpoint,
		decodeDeleteAbilityRequest,
		encodeResponse,
	))

	//r.Methods("GET").Path("/test").Handler(httptransport.NewServer(
	//	getAbilityEndpoint,
	//	decodeGetAbilityRequest,
	//	encodeResponse,
	//))

	return r
}

func decodeCreateAbilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createAbilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Ability); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateAbilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req updateAbilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Ability); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteAbilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req deleteAbilityRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Ability); err != nil {
		return nil, err
	}
	return req, nil
}



//func decodeGetAbilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
//	var req GetAbilityRequest
//	if err := json.NewDecoder(r.Body).Decode(&req.Ability); err != nil {
//		return nil, err
//	}
//	return req, nil
//}


func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
