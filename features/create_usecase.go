package features

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/c3k4ah/cleanit/contents"
	"github.com/c3k4ah/cleanit/models"
	"github.com/c3k4ah/cleanit/utils"
	"gopkg.in/yaml.v3"
)

func CreateUseCases() {
	yamlFile, err := os.ReadFile("clean_it_config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config models.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error1: %v", err)
	}

	for featureName, feature := range config.Features {
		for _, usecase := range feature.UseCases {
			tmpl, err := template.New(usecase.Name).Parse(contents.UseCaseTemplate)
			fileName := fmt.Sprintf("lib/features/%s/domain/usecases/%s.dart", featureName, strings.ToLower(usecase.Name))

			if err != nil {
				log.Fatalf("error initializing template : %v", err)
			}

			usecaseData := struct {
				Name       string
				Method     string
				ReturnType string
				ParamsType string
				Repository string
			}{
				Name:       usecase.Name,
				Method:     usecase.Method,
				ReturnType: usecase.ReturnType,
				ParamsType: usecase.ParamsType,
				Repository: feature.DomainRepository,
			}
			utils.CreateFileIfNotExists(fileName, tmpl, usecaseData)
		}
	}
}
