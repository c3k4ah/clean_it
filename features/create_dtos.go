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

func CreateDTOs() {
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
		// Create DTO model
		tmpl, err := template.New("dto_model").Parse(contents.DtoModelTemplate)
		if err != nil {
			log.Fatalf("error initializing template : %v", err)
		}

		fileName := fmt.Sprintf("lib/core/DTO/models/%s_model.dart", featureName)
		utils.CreateFileIfNotExists(fileName, tmpl, feature.DTO)

		// Create DTO entity
		tmpl, err = template.New("dto_entity").Parse(contents.DtoEntityTemplate)
		if err != nil {
			log.Fatalf("error initializing template : %v", err)
		}

		fileName = fmt.Sprintf("lib/core/DTO/entities/%s_entity.dart", featureName)
		utils.CreateFileIfNotExists(fileName, tmpl, feature.DTO)
	}
}
