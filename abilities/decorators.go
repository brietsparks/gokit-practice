package abilities

import (
	"github.com/go-kit/kit/log"
	"time"
	"gokit-practice/util"
)

type serviceWithLogger struct {
	logger log.Logger
	Service
}

func WithLogger(logger log.Logger, s Service) Service {
	return &serviceWithLogger{logger, s}
}

func (s *serviceWithLogger) CreateAbility(a Ability) (Ability, error) {
	return writeAbilityWithLogger(s.logger, s.Service.CreateAbility, a, "CreateAbility")
}

func (s *serviceWithLogger) UpdateAbility(a Ability) (Ability, error) {
	return writeAbilityWithLogger(s.logger, s.Service.UpdateAbility, a, "UpdateAbility")
}

func (s *serviceWithLogger) DeleteAbility(a Ability) (Ability, error) {
	return writeAbilityWithLogger(s.logger, s.Service.DeleteAbility, a, "DeleteAbility")
}

func writeAbilityWithLogger(logger log.Logger, method serviceWriteMethod, input Ability, methodName string) (Ability, error) {
	begin := time.Now()
	output, err := method(input)
	end := time.Now()

	util.LogFunctionCall(logger, util.LogEntry{
		Function: methodName,
		Input: input,
		Output: output,
		Err: err,
		Begin: begin,
		End: end,
	})

	return output, err
}


func (s *serviceWithLogger) GetAbilitiesByOwnerId(input string) ([]Ability, error) {
	begin := time.Now()
	output, err := s.Service.GetAbilitiesByOwnerId(input)
	end := time.Now()

	util.LogFunctionCall(s.logger, util.LogEntry{
		Function: "GetAbilitiesByOwnerId",
		Input: input,
		Output: output,
		Err: err,
		Begin: begin,
		End: end,
	})

	return output, err
}
