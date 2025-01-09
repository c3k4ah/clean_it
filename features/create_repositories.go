package features

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/c3k4ah/cleanit/contents"
	"github.com/c3k4ah/cleanit/models"
	"github.com/c3k4ah/cleanit/utils"
	"gopkg.in/yaml.v3"
)

func CreateRepositories() {
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
		// Create domain repository
		tmpl, err := template.New("domain_repository").Parse(contents.DomainRepositoryTemplate)
		if err != nil {
			log.Fatalf("error initializing template : %v", err)
		}

		fileName := fmt.Sprintf("lib/features/%s/domain/repository/repository.dart", featureName)
		utils.CreateFileIfNotExists(fileName, tmpl, feature)

		// Create data repository
		tmpl, err = template.New("data_repository").Parse(contents.DataRepositoryTemplate)
		if err != nil {
			log.Fatalf("error initializing template : %v", err)
		}

		fileName = fmt.Sprintf("lib/features/%s/data/repositories/repository_impl.dart", featureName)
		utils.CreateFileIfNotExists(fileName, tmpl, feature)

		// Create data source
		tmpl, err = template.New("data_source").Parse(contents.DataSourceTemplate)
		if err != nil {
			log.Fatalf("error initializing template : %v", err)
		}

		fileName = fmt.Sprintf("lib/features/%s/data/datasource/remote/%s_datasource.dart", featureName, featureName)
		utils.CreateFileIfNotExists(fileName, tmpl, feature)
	}
}
