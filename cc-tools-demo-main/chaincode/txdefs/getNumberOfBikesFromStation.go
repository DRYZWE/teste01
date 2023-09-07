package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Return the number of bikes of a Station
// GET method
var GetNumberOfBikesFromStation = tx.Transaction{
	Tag:         "getNumberOfBikesFromStation",
	Label:       "Get Number Of Bikes From Station",
	Description: "Return the number of bikes of a station",
	Method:      "GET",
	Callers:     []string{"$org2MSP", "$orgMSP"}, // Only org2 can call this transactions

	Args: []tx.Argument{
		{
			Tag:         "station",
			Label:       "Station",
			Description: "Station",
			DataType:    "->station",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		stationKey, _ := req["station"].(assets.Key)

		// Returns Station from channel
		stationMap, err := stationKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}

		numberOfBikes := 0
		bikes, ok := stationMap["bikes"].([]interface{})
		if ok {
			numberOfBikes = len(bikes)
		}

		returnMap := make(map[string]interface{})
		returnMap["numberOfBikes"] = numberOfBikes

		// Marshal asset back to JSON format
		returnJSON, nerr := json.Marshal(returnMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return returnJSON, nil
	},
}
