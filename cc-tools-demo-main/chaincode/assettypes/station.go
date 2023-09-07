package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a Station as a collection of books
var Station = assets.AssetType{
	Tag:         "Station",
	Label:       "Station",
	Description: "Station as a collection of bikes",

	Props: []assets.AssetProp{
		{
			// Primary Key
			Required: true,
			IsKey:    true,
			Tag:      "name",
			Label:    "Station Name",
			DataType: "string",
			Writers:  []string{`org3MSP`, "orgMSP"}, // This means only org3 can create the asset (others can edit)
		},
		{
			// Asset reference list
			Tag:      "bikes",
			Label:    "Bikes Collection",
			DataType: "[]->bike",
		},
		{
			// Asset reference list
			Tag:      "entranceCode",
			Label:    "Entrance Code for the Station",
			DataType: "->secret",
		},
	},
}
