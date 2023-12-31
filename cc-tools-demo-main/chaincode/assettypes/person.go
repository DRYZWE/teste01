package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

var Person = assets.AssetType{
	Tag:         "person",
	Label:       "Person",
	Description: "Personal data of someone",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "CPF (Brazilian ID)",
			DataType: "cpf",                         // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name of the person1",
			DataType: "string",
			// Validate funcion
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("Por Favor insira um nome Válido")
				}
				return nil
			},
		},
		{
			// Optional property
			Tag:      "dateOfBirth",
			Label:    "Date of Birth",
			DataType: "datetime",
		},
	},
}
