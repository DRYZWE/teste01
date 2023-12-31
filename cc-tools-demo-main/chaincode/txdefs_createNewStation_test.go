package main_test

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

	cc "github.com/goledgerdev/cc-tools-demo/chaincode"
	"github.com/goledgerdev/cc-tools/mock"
)

func TestCreateNewStation(t *testing.T) {
	stub := mock.NewMockStub("org3MSP", new(cc.CCDemo))

	expectedResponse := map[string]interface{}{
		"@key":         "station:3cab201f-9e2b-579d-b7b2-72297ed17f49",
		"@lastTouchBy": "org3MSP",
		"@lastTx":      "createNewStation",
		"@assetType":   "station",
		"name":         "Praia Vermelha Campus",
	}
	req := map[string]interface{}{
		"name": "Praia Vermelha Campus",
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		t.FailNow()
	}

	res := stub.MockInvoke("createNewStation", [][]byte{
		[]byte("createNewStation"),
		reqBytes,
	})

	expectedResponse["@lastUpdated"] = stub.TxTimestamp.AsTime().Format(time.RFC3339)

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

	var state map[string]interface{}
	stateBytes := stub.State["station:3cab201f-9e2b-579d-b7b2-72297ed17f49"]
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
