package datatypes

import (
	"fmt"
	"strconv"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

// Example of a custom data type using enum-like structure (iota)
// This allows the use of verification by const values instead of float64, improving readability
// Example:
// 		if assetMap["bikekType"].(float64) == (float64)(BikekTypeHardcover)
// 			...

type BikeType float64

const (
	BikeTypeHardcover BikeType = iota
	BikeTypePaperback
	BikeTypeEbook
)

// CheckType checks if the given value is defined as valid BikeType consts
func (b BikeType) CheckType() errors.ICCError {
	switch b {
	case BikeTypeHardcover:
		return nil
	case BikeTypePaperback:
		return nil
	case BikeTypeEbike:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}

}

var bikeType = assets.DataType{
	AcceptedFormats: []string{"number"},
	DropDownValues: map[string]interface{}{
		"Hardcover": BikeTypeHardcover,
		"Paperback": BikeTypePaperback,
		"Ebike":     BikeTypeEbike,
	},
	Description: ``,

	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64
		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = (float64)(v)
		case BikeType:
			dataVal = (float64)(v)
		case string:
			var err error
			dataVal, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "asset property must be an integer, is %t", 400)
			}
		default:
			return "", nil, errors.NewCCError("asset property must be an integer, is %t", 400)
		}

		retVal := (BikeType)(dataVal)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}
