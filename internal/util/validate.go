package util

import (
	termError "github.com/roysti10/termCI/internal/errors"
	structs "github.com/roysti10/termCI/internal/structs"
	yaml "gopkg.in/yaml.v2"
	"os"
)

func Validate() (*structs.Pipeline, error) {
	if err := validateFileExists(); err != nil {
		return nil, err
	}
	pipeline, err := validateYaml()
	if err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func validateFileExists() error {
	_, err := os.Stat(".term/term.yml")
	if os.IsNotExist(err) {
		return termError.TermError{Msg: err.Error()}
	}
	if err != nil {
		return termError.TermError{Msg: err.Error()}
	}
	return nil
}

func validateYaml() (structs.Pipeline, error) {
	var pipeline structs.Pipeline
	data, err := os.ReadFile(".term/term.yml")
	if err != nil {
		return pipeline, err
	}
	if err := yaml.Unmarshal(data, &pipeline); err != nil {
		return pipeline, err
	}
	return pipeline, nil
}
