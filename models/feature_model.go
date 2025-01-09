package models

type Feature struct {
	DTO              DTO       `yaml:"DTO"`
	DomainRepository string    `yaml:"domain_repository"`
	DataRepository   string    `yaml:"data_repository"`
	Source           string    `yaml:"remote_source"`
	UseCases         []UseCase `yaml:"usecases"`
}
