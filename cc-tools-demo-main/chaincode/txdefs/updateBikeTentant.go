package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Updates the tenant of a Bike
// POST Method
var UpdateBikeTenant = tx.Transaction{
	Tag:         "updateBikeTenant",
	Label:       "Update Bike Tenant",
	Description: "Change the tenant of a Bike",
	Method:      "PUT",
	Callers:     []string{`$org\dMSP`, "orgMSP"}, // Any orgs can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "bike",
			Label:       "Bike",
			Description: "Bike",
			DataType:    "->bike",
			Required:    true,
		},
		{
			Tag:         "tenant",
			Label:       "tenant",
			Description: "New tenant of the bike",
			DataType:    "->person",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		bikeKey, ok := req["bike"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter bike must be an asset")
		}
		tenantKey, ok := req["tenant"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter tenant must be an asset")
		}

		// Returns Bike from channel
		bikeAsset, err := bikeKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		bikeMap := (map[string]interface{})(*bikeAsset)

		// Returns person from channel
		tenantAsset, err := tenantKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		tenantMap := (map[string]interface{})(*tenantAsset)

		updatedTenantKey := make(map[string]interface{})
		updatedTenantKey["@assetType"] = "person"
		updatedTenantKey["@key"] = tenantMap["@key"]

		// Update data
		bikeMap["currentTenant"] = updatedTenantKey

		bikeMap, err = bikeAsset.Update(stub, bikeMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update asset")
		}

		// Marshal asset back to JSON format
		bikeJSON, nerr := json.Marshal(bikeMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return bikeJSON, nil
	},
}
