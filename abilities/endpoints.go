package abilities

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

type createAbilityRequest struct {
	Ability Ability
}

type createAbilityResponse struct {
	Ability Ability
	Err     error
}

func MakeCreateAbilityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createAbilityRequest)
		ability, err := s.CreateAbility(ctx, req.Ability)
		return createAbilityResponse{Ability: ability}, err
	}
}

type updateAbilityRequest struct {
	Ability Ability
}

type updateAbilityResponse struct {
	Ability Ability
	Err     error
}

func MakeUpdateAbilityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateAbilityRequest)
		ability, err := s.UpdateAbility(ctx, req.Ability)
		return updateAbilityResponse{Ability: ability}, err
	}
}

type deleteAbilityRequest struct {
	Ability Ability
}

type deleteAbilityResponse struct {
}

func MakeDeleteAbilityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteAbilityRequest)
		err := s.DeleteAbility(ctx, req.Ability)
		return deleteAbilityResponse{}, err
	}
}

//type GetAbilityRequest struct {
//	Ability Ability
//}
//
//type getAbilityResponse struct {
//	Ability Ability
//	Err     error
//}
//
//func MakeGetAbilityEndpoint(s Service) endpoint.Endpoint {
//	return func(ctx context.Context, request interface{}) (interface{}, error) {
//		req := request.(GetAbilityRequest)
//		ability, err := s.GetAbility(ctx, req.Ability)
//		return getAbilityResponse{Ability: ability}, err
//	}
//}


