package models

type UseCase struct {
	Name       string `yaml:"name"`
	Method     string `yaml:"method"`
	ReturnType string `yaml:"return_type"`
	ParamsType string `yaml:"params_type"`
}
