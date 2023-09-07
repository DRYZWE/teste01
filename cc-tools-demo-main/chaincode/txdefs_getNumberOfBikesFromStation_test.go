package main

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/goledgerdev/cc-tools/mock"
)

func TestGetNumberOfBikesFromStation(t *testing.T) {
	stub := mock.NewMockStub("org2MSP", new(CCDemo))

	// Setup state
	setupBike := map[string]interface{}{
		"@key":         "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
		"@lastTouchBy": "org2MSP",
		"@lastTx":      "createAsset",
		"@assetType":   "bike",
		"title":        "Bike 000",
		"author":       "UFF",
		"published":    "2019-05-06T22:12:41Z",
	}
	setupStation := map[string]interface{}{
		"@key":         "station:3cab201f-9e2b-579d-b7b2-72297ed17f49",
		"@lastTouchBy": "org3MSP",
		"@lastTx":      "createNewStation",
		"@assetType":   "station",
		"name":         "Praia Vermelha Campus",
		"bikes": []map[string]interface{}{
			{
				"@assetType": "bike",
				"@key":       "bike:a36a2920-c405-51c3-b584-dcd758338cb5",
			},
		},
	}
	setupBikeJSON, _ := json.Marshal(setupBike)
	setupStationJSON, _ := json.Marshal(setupStation)

	stub.MockTransactionStart("setupGetNumberOfBikesFromStation")
	stub.PutState("bike:a36a2920-c405-51c3-b584-dcd758338cb5", setupBikeJSON)
	stub.PutState("station:3cab201f-9e2b-579d-b7b2-72297ed17f49", setupStationJSON)
	refIdx, err := stub.CreateCompositeKey("bike:a36a2920-c405-51c3-b584-dcd758338cb5", []string{"station:3cab201f-9e2b-579d-b7b2-72297ed17f49"})
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	stub.PutState(refIdx, []byte{0x00})
	stub.MockTransactionEnd("setupGetNumberOfBikesFromStation")

	expectedResponse := map[string]interface{}{
		"numberOfBikes": 1.0,
	}
	req := map[string]interface{}{
		"station": map[string]interface{}{
			"name": "Praia Vermelha Campus",
		},
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		t.FailNow()
	}

	res := stub.MockInvoke("getNumberOfBikesFromStation", [][]byte{
		[]byte("getNumberOfBikesFromStation"),
		reqBytes,
	})

	if res.GetStatus() != 200 {
		log.Println(res)
		t.FailNow()
	}

	var resPayload map[string]interface{}
	err = json.Unmarshal(res.GetPayload(), &resPayload)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(resPayload, expectedResponse) {
		log.Println("these should be equal")
		log.Printf("%#v\n", resPayload)
		log.Printf("%#v\n", expectedResponse)
		t.FailNow()
	}
}
