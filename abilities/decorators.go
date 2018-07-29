package abilities

import (
	"github.com/go-kit/kit/log"
	"time"
	"encoding/json"
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

	logEntry := logEntry{
		method: methodName,
		input: input,
		output: output,
		err: err,
		begin: begin,
		end: end,
	}
	logMethodCall(logger, logEntry)

	return output, err
}

type logEntry struct {
	method string
	input interface{}
	output interface{}
	err error
	begin time.Time
	end time.Time
}


func logMethodCall(logger log.Logger, entry logEntry) {
	input, _ := json.Marshal(entry.input)
	output, _ := json.Marshal(entry.output)
	logger.Log(
		"method", entry.method,
		"input", input,
		"output", output,
		"err", entry.err,
		"began", entry.begin,
		"end", entry.end,
		"took", entry.end.Sub(entry.begin),
	)
}


func (s *serviceWithLogger) GetAbilitiesByOwnerId(input string) ([]Ability, error) {
	var output []Ability
	var err error

	defer func(begin time.Time) {
		output, _ := json.Marshal(output)
		s.logger.Log(
			"method", "GetAbilitiesByOwnerId",
			"input", input,
			"output", string(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = s.Service.GetAbilitiesByOwnerId(input)
	return output, err
}
