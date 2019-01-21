package secret

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/property"
)

func resolveSecretValue(encrypted string) (string, error) {
	decodedValue, err := GetSecretValueHandler().DecodeValue(encrypted)
	if err != nil {
		return "", err
	}
	return decodedValue, nil
}

func PropertyProcessor(properties property.Properties) error {

	for key, value := range properties {
		if value == nil || value.DataType != data.TypeSecret {
			continue
		}

		if strVal, ok := value.Value.(string); ok {
			// Resolve secret value
			newVal, err := resolveSecretValue(strVal)
			if err != nil {
				return err
			}
			properties[key].Value = newVal
		}
	}

	return nil
}
