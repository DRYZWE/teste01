package main

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/goledgerdev/cc-tools/mock"
)

func TestUpdateBikeTenant(t *testing.T) {
	stub := mock.NewMockStub("org1MSP", new(CCDemo))

	// State setup
	setupPerson := map[string]interface{}{
		"@key":         "person:47061146-c642-51a1-844a-bf0b17cb5e19",
		"@lastTouchBy": "org1MSP",
		"@lastTx":      "createAsset",
		"@assetType":   "person",
		"name":         "Daniel",
		"id":           "31820792048",
	}
	setupBike := map[string]interface{}{
		"@key":         "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
		"@lastTouchBy": "org2MSP",
		"@lastTx":      "createAsset",
		"@assetType":   "bike",
		"Name":        "Bike 0010",
		"Owner":       "UFF",
		// "currentTenant": map[string]interface{}{
		// 	"@assetType": "person",
		// 	"@key":       "person:47061146-c642-51a1-844a-bf0b17cb5e19",
		// },
		"published": "2019-05-06T22:12:41Z",
	}
	setupPersonJSON, _ := json.Marshal(setupPerson)
	setupBikeJSON, _ := json.Marshal(setupBike)

	stub.MockTransactionStart("setupUpdateBikeTenant")
	stub.PutState("person:47061146-c642-51a1-844a-bf0b17cb5e19", setupPersonJSON)
	stub.PutState("bike:a36a2920-c405-51c3-b584-dcd758338cb5", setupBikeJSON)
	stub.MockTransactionEnd("setupUpdateBikeTenant")

	req := map[string]interface{}{
		"bike": map[string]interface{}{
			"@key": "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
		},
		"tenant": map[string]interface{}{
			"@key": "person:47061146-c642-51a1-844a-bf0b17cb5e19",
		},
	}
	reqBytes, _ := json.Marshal(req)

	res := stub.MockInvoke("updateBikeTenant", [][]byte{
		[]byte("updateBikeTenant"),
		reqBytes,
	})

	if res.GetStatus() != 200 {
		log.Println(res)
		t.FailNow()
	}

	var resPayload map[string]interface{}
	err := json.Unmarshal(res.GetPayload(), &resPayload)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	expectedResponse := map[string]interface{}{
		"@key":         "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
		"@lastTouchBy": "org1MSP",
		"@lastTx":      "updateBikeTenant",
		"@assetType":   "bike",
		"Name":        "Bike 0010",
		"Owner":       "UFF",
		"currentTenant": map[string]interface{}{
			"@assetType": "person",
			"@key":       "person:47061146-c642-51a1-844a-bf0b17cb5e19",
		},
		"published": "2019-05-06T22:12:41Z",
	}

	expectedResponse["@lastUpdated"] = stub.TxTimestamp.AsTime().Format(time.RFC3339)

	if !reflect.DeepEqual(resPayload, expectedResponse) {
		log.Println("these should be equal")
		log.Printf("%#v\n", resPayload)
		log.Printf("%#v\n", expectedResponse)
		t.FailNow()
	}

	var state map[string]interface{}
	stateBytes := stub.State["bike:a36a2920-c405-51c3-b584-dcd758338cb5"]
	err = json.Unmarshal(stateBytes, &state)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(state, expectedResponse) {
		log.Println("these should be equal")
		log.Printf("%#v\n", state)
		log.Printf("%#v\n", expectedResponse)
		t.FailNow()
	}
}
