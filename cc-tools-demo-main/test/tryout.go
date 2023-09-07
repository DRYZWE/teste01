package main

import (
	"fmt"
	"log"
)

func tryout() []error {
	var err error

	// Get Header
	fmt.Print("Get Header... ")
	err = GetAndVerify(
		"http://localhost:80/api/query/getHeader",
		200,
		map[string]interface{}{
			"ccToolsVersion": "v0.7.1",
			"colors": []interface{}{
				"#4267B2",
				"#34495E",
				"#ECF0F1",
			},
			"name":     "CC Tools Demo",
			"orgMSP":   "org1MSP",
			"orgTitle": "CC Tools Demo",
			"version":  "1.0.0",
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Get Transactions
	fmt.Print("Get Transactions... ")
	err = GetAndVerify(
		"http://localhost:80/api/query/getTx",
		200,
		[]interface{}{
			map[string]interface{}{
				"description": "",
				"label":       "Create Asset",
				"tag":         "createAsset",
			},
			map[string]interface{}{
				"description": "",
				"label":       "Update Asset",
				"tag":         "updateAsset",
			},
			map[string]interface{}{
				"description": "",
				"label":       "Delete Asset",
				"tag":         "deleteAsset",
			},
			map[string]interface{}{
				"callers":     []interface{}{"$org3MSP"},
				"description": "Create a New Station",
				"label":       "Create New Station",
				"tag":         "createNewStation",
			},
			map[string]interface{}{
				"callers":     []interface{}{"$org2MSP"},
				"description": "Return the number of bikes of a station",
				"label":       "Get Number Of Bikes From Station",
				"tag":         "getNumberOfBikesFromStation",
			},
			map[string]interface{}{
				"callers":     []interface{}{"$org\\dMSP"},
				"description": "Change the tenant of a bike",
				"label":       "Update Bike Tenant",
				"tag":         "updateBikeTenant",
			},
			map[string]interface{}{
				"description": "",
				"label":       "Get Tx",
				"tag":         "getTx",
			},
			map[string]interface{}{
				"description": "",
				"label":       "Get Header",
				"tag":         "getHeader",
			},
			map[string]interface{}{
				"description": "",
				"label":       "Get Schema",
				"tag":         "getSchema",
			},
			map[string]interface{}{
				"description": "GetDataTypes returns the primary data type map",
				"label":       "Get DataTypes",
				"tag":         "getDataTypes",
			},
			map[string]interface{}{
				"description": "",
				"label":       "Read Asset",
				"tag":         "readAsset",
			},
			map[string]interface{}{
				"description": "",
				"label":       "Read Asset History",
				"tag":         "readAssetHistory",
			},
			map[string]interface{}{
				"description": "",
				"label":       "Search World State",
				"tag":         "search",
			},
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Get CreateAsset definition
	fmt.Print("Get CreateAsset definition... ")
	err = PostAndVerify(
		"http://localhost:80/api/query/getTx",
		map[string]interface{}{
			"txName": "createAsset",
		},
		200,
		map[string]interface{}{
			"args": []interface{}{
				map[string]interface{}{
					"dataType":    "[]@asset",
					"description": "List of assets to be created.",
					"label":       "",
					"private":     false,
					"required":    true,
					"tag":         "asset",
				},
			},
			"description": "",
			"label":       "Create Asset",
			"metaTx":      true,
			"method":      "POST",
			"readOnly":    false,
			"tag":         "createAsset",
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Get Asset Types
	fmt.Print("Get Asset Types... ")
	err = GetAndVerify(
		"http://localhost:80/api/query/getSchema",
		200,
		[]interface{}{
			map[string]interface{}{
				"description": "Personal data of someone",
				"label":       "Person",
				"tag":         "person",
				"writers":     nil,
			},
			map[string]interface{}{
				"description": "Bike",
				"label":       "Bike",
				"tag":         "bike",
				"writers":     nil,
			},
			map[string]interface{}{
				"description": "Station as a collection of bikes",
				"label":       "Station",
				"tag":         "station",
				"writers":     nil,
			},
			map[string]interface{}{
				"description": "Secret between Org2 and Org3",
				"label":       "Secret",
				"readers": []interface{}{
					"org2MSP",
					"org3MSP",
				},
				"tag":     "secret",
				"writers": nil,
			},
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Get Person asset type definition
	fmt.Print("Get Person asset type definition... ")
	err = PostAndVerify(
		"http://localhost:80/api/query/getSchema",
		map[string]interface{}{
			"assetType": "person",
		},
		200,
		map[string]interface{}{
			"tag":         "person",
			"label":       "Person",
			"description": "Personal data of someone",
			"props": []interface{}{
				map[string]interface{}{
					"dataType":    "cpf",
					"description": "",
					"isKey":       true,
					"label":       "CPF (Brazilian ID)",
					"readOnly":    false,
					"required":    true,
					"tag":         "id",
					"writers": []interface{}{
						"org1MSP",
					},
				},
				map[string]interface{}{
					"dataType":    "string",
					"description": "",
					"isKey":       false,
					"label":       "Name of the person",
					"readOnly":    false,
					"required":    true,
					"tag":         "name",
					"writers":     nil,
				},
				map[string]interface{}{
					"dataType":    "datetime",
					"description": "",
					"isKey":       false,
					"label":       "Date of Birth",
					"readOnly":    false,
					"required":    false,
					"tag":         "dateOfBirth",
					"writers":     nil,
				},
				,
			},
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Create Person
	fmt.Print("Create Person... ")
	err = PostAndVerify(
		"http://localhost:80/api/invoke/createAsset",
		map[string]interface{}{
			"asset": []map[string]interface{}{
				{
					"@assetType": "person",
					"name":       "Daniel",
					"id":         "318.207.920-48",
				},
			},
		},
		200,
		[]interface{}{
			map[string]interface{}{
				"@assetType":   "person",
				"@key":         "person:47061146-c642-51a1-844a-bf0b17cb5e19",
				"@lastTouchBy": "org1MSP",
				"@lastTx":      "createAsset",
				"id":           "31820792048",
				"name":         "Daniel",
			},
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Create Bike
	fmt.Print("Create Bike... ")
	err = PostAndVerify(
		"http://localhost:980/api/invoke/createAsset",
		map[string]interface{}{
			"asset": []map[string]interface{}{
				{
					"@assetType": "bike",
					"Name":      "Bike 001",
					"Owner":     "UFF",
					"currentTenant": map[string]interface{}{
						"id": "318.207.920-48",
					},
					},
					"published": "2019-05-06T22:12:41Z",
				},
			},
		},
		200,
		[]interface{}{
			map[string]interface{}{
				"@assetType":   "bike",
				"@key":         "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
				"@lastTouchBy": "org2MSP",
				"@lastTx":      "createAsset",
				"Name":        "Bike 001",
				"Owner":       "UFF",
				"currentTenant": map[string]interface{}{
					"@assetType": "person",
					"@key":       "person:47061146-c642-51a1-844a-bf0b17cb5e19",
				},
				},
				"published": "2019-05-06T22:12:41Z",
			},
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Read Bike
	fmt.Print("Read Bike... ")
	err = PostAndVerify(
		"http://localhost:80/api/query/readAsset",
		map[string]interface{}{
			"key": map[string]interface{}{
				"@assetType": "bike",
				"Owner":     "UFF",
				"Name":      "Bike 001",
			},
			"resolve": true,
		},
		200,
		map[string]interface{}{
			"@assetType":   "bike",
			"@key":         "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
			"@lastTouchBy": "org2MSP",
			"@lastTx":      "createAsset",
			"Name":        "Bike 001",
			"Owner":       "UFF",
			"currentTenant": map[string]interface{}{
				"@assetType":   "person",
				"@key":         "person:47061146-c642-51a1-844a-bf0b17cb5e19",
				"@lastTouchBy": "org1MSP",
				"@lastTx":      "createAsset",
				"id":           "31820792048",
				"name":         "Daniel",
			},
			},
			"published": "2019-05-06T22:12:41Z",
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Update Person
	fmt.Print("Update Person... ")
	err = PostAndVerify(
		"http://localhost:80/api/invoke/updateAsset",
		map[string]interface{}{
			"update": map[string]interface{}{
				"@assetType": "person",
				"id":         "318.207.920-48",
				"name":       "Daniel",
			},
		},
		200,
		map[string]interface{}{
			"@assetType":   "person",
			"@key":         "person:47061146-c642-51a1-844a-bf0b17cb5e19",
			"@lastTouchBy": "org1MSP",
			"@lastTx":      "updateAsset",
			"id":           "31820792048",
			"name":         "Daniel",
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Check if person was updated
	fmt.Print("Check if person was updated... ")
	err = PostAndVerify(
		"http://localhost:80/api/query/readAsset",
		map[string]interface{}{
			"key": map[string]interface{}{
				"@assetType": "person",
				"id":         "318.207.920-48",
			},
		},
		200,
		map[string]interface{}{
			"@assetType":   "person",
			"@key":         "person:47061146-c642-51a1-844a-bf0b17cb5e19",
			"@lastTouchBy": "org1MSP",
			"@lastTx":      "updateAsset",
			"id":           "31820792048",
			"name":         "Daniel",
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Query all bikes using CouchDB
	fmt.Print("Query all bikes using CouchDB... ")
	err = PostAndVerify(
		"http://localhost:80/api/query/search",
		map[string]interface{}{
			"query": map[string]interface{}{
				"selector": map[string]interface{}{
					"@assetType": "bike",
				},
			},
			"resolve": true,
		},
		200,
		map[string]interface{}{
			"metadata": map[string]interface{}{},
			"result": []interface{}{
				map[string]interface{}{
					"@assetType":   "bike",
					"@key":         "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
					"@lastTouchBy": "org2MSP",
					"@lastTx":      "createAsset",
					"Owner":       "UFF",
					"currentTenant": map[string]interface{}{
						"@assetType":   "person",
						"@key":         "person:47061146-c642-51a1-844a-bf0b17cb5e19",
						"@lastTouchBy": "org1MSP",
						"@lastTx":      "updateAsset",
						"id":           "31820792048",
						"name":         "Daniel"},

						"published": "2019-05-06T22:12:41Z",
					"Name":     "Bike 001",
				},
			},
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Delete bike
	fmt.Print("Delete bike... ")
	err = PostAndVerify(
		"http://localhost:980/api/invoke/deleteAsset",
		map[string]interface{}{
			"key": map[string]interface{}{
				"@assetType": "bike",
				"@key":       "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
			},
		},
		200,
		map[string]interface{}{
			"@assetType":   "bike",
			"@key":         "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
			"@lastTouchBy": "org2MSP",
			"@lastTx":      "createAsset",
			"Name":        "Bike 001",
			"Owner":       "UFF",
			"currentTenant": map[string]interface{}{
				"@assetType": "person",
				"@key":       "person:47061146-c642-51a1-844a-bf0b17cb5e19",
			},
			"published": "2019-05-06T22:12:41Z",
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	// Delete person
	fmt.Print("Delete person... ")
	err = PostAndVerify(
		"http://localhost:80/api/invoke/deleteAsset",
		map[string]interface{}{
			"key": map[string]interface{}{
				"@assetType": "person",
				"@key":       "person:47061146-c642-51a1-844a-bf0b17cb5e19",
			},
		},
		200,
		map[string]interface{}{
			"@assetType":   "person",
			"@key":         "person:47061146-c642-51a1-844a-bf0b17cb5e19",
			"@lastTouchBy": "org1MSP",
			"@lastTx":      "updateAsset",
			"id":           "31820792048",
			"name":         "Daniel",
		},
	)
	if err != nil {
		fail()
		log.Fatalln(err)
	}
	pass()

	return nil
}
