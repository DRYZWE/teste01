package assettypes

import "github.com/goledgerdev/cc-tools/assets"

// Description of a bike
var Bike = assets.AssetType{
	Tag:         "bike",
	Label:       "Bike",
	Description: "Bike in UFF domain",

	Props: []assets.AssetProp{
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "Name",
			Label:    "Name Bike1",
			DataType: "string",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			// Composite Key
			Required: true,
			IsKey:    true,
			Tag:      "Owner",
			Label:    "Bike Owner",
			DataType: "string",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},
		{
			/// Reference to another asset
			Tag:      "currentTenant",
			Label:    "Current Tenant",
			DataType: "->person",
		},
		{
			// Date property
			Tag:      "Created",
			Label:    "Created Date",
			DataType: "datetime",
		},
		{
			// Custom data type
			Tag:      "bikeType",
			Label:    "Bike Type",
			DataType: "bikeType",
		},
	},
}
