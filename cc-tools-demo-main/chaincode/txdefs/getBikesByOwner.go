package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Return the all bikes from an specific Owner
// GET method
var GetBikeByOwner = tx.Transaction{
	Tag:         "getBikesByOwner",
	Label:       "Get Bikes by the Owner Name",
	Description: "Return all the bikes from an Owner",
	Method:      "GET",
	Callers:     []string{"$org1MSP", "$org2MSP", "$orgMSP"}, // Only org1 and org2 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "Owner",
			Label:       "Owner",
			Description: "Owner",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "limit",
			Label:       "Limit",
			Description: "Limit",
			DataType:    "number",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		Owner, _ := req["Owner"].(string)
		limit, hasLimit := req["limit"].(float64)

		if hasLimit && limit <= 0 {
			return nil, errors.NewCCError("limit must be greater than 0", 400)
		}

		// Prepare couchdb query
		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType": "bike",
				"Owner":     Owner,
			},
		}

		if hasLimit {
			query["limit"] = limit
		}

		var err error
		response, err := assets.Search(stub, query, "", true)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error searching for bike's owner", 500)
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error marshaling response", 500)
		}

		return responseJSON, nil
	},
}
