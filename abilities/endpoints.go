package abilities

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

type abilityWriteRequest struct {
	Ability Ability
}

type abilityWriteResponse struct {
	Ability Ability
	Err     error
}

func makeWriteEndpoint(method serviceWriteMethod) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(abilityWriteRequest)
		ability, err := method(req.Ability)
		return abilityWriteResponse{Ability: ability}, err // todo where should err be passed into
	}
}

func MakeCreateAbilityEndpoint(s Service) endpoint.Endpoint {
	return makeWriteEndpoint(s.CreateAbility)
}

func MakeUpdateAbilityEndpoint(s Service) endpoint.Endpoint {
	return makeWriteEndpoint(s.UpdateAbility)
}

func MakeDeleteAbilityEndpoint(s Service) endpoint.Endpoint {
	return makeWriteEndpoint(s.DeleteAbility)
}

type abilitiesReadRequest struct {
	OwnerId string
}

type abilitiesReadResponse struct {
	Abilities []Ability
	Err       error
}

func MakeGetAbilitiesByOwnerIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(abilitiesReadRequest)
		abilities, err := s.GetAbilitiesByOwnerId(req.OwnerId)
		return abilitiesReadResponse{Abilities: abilities}, err
	}
}
