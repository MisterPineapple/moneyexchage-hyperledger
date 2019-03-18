package main

import (
	"encoding/json"
	"fmt"
	//"bytes"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ExchangeRequestTransaction struct {
	ObjectType		string 			`json:"docType" valid:"required"`
	Id				string			`json:"id" valid:"required"`
	RequestId 		string 			`json:"requestId" valid:"required"`
	Origin 			string 			`json:"origin"`
	Target 			string 			`json:"target"`
	OriginAmount 	string 			`json:"originamount"`
	TargetAmount 	string 			`json:"targetamount"`
}

const schemaExchangeRequestTransaction = `{
	"definitions": {},
	"$schema": "http://json-schema.org/draft-07/schema#",
	"$id": "http://example.com/root.json",
	"type": "object",
	"title": "The ExchangeRequest Schema",
	"required": [
		"id"
	],
	"properties": {
		"id": {
			"$id": "#/properties/id",
			"type": "string",
			"title": "The Id Schema"
		},
		"requestId": {
			"$id": "#/properties/requestId",
			"type": "string",
			"title": "The requestId Schema"
		},
		"origin": {
			"$id": "#/properties/origin",
			"type": "string",
			"title": "The Origin Schema"
		},
		"target": {
			"$id": "#/properties/target",
			"type": "string",
			"title": "The Target Schema"
		},
		"originamount": {
			"$id": "#/properties/originamount",
			"type": "string",
			"title": "The OriginAmount Schema"
		},
		"targetamount": {
			"$id": "#/properties/targetamount",
			"type": "string",
			"title": "The TargetAmount Schema"
		}
	}
}`

func (t *Account) ExchangeRequestTx(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	methodName := "[ExchangeRequest]"
	enrollmentID, err := getEnrollmentID(stub)
	if err != nil {
	  return shim.Error(methodName + " " + err.Error())
	}
	logger.Infof(methodName + " Creator enrollmentID: " + enrollmentID)

	s := args[0]
    var data Account
	err = json.Unmarshal([]byte(s), &data)
	key := data.Id

	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(methodName + " " + jsonResp)
	}

	var requestaccount Account
	err = ParseInput(valAsbytes, &requestaccount, schemaDepositAccount)
	if err != nil {
		return shim.Error(methodName + ": " + err.Error())
	}

	var exchangeRequestTx ExchangeRequestTransaction
	exchangeRequestTx.ObjectType = "request_exchange_tx"
	err = ParseInput([]byte(args[0]), &exchangeRequestTx, schemaExchangeRequestTransaction)
	if err != nil {
		return shim.Error(methodName + ": " + err.Error())
	}

	err = t.exchangeRequest(stub, exchangeRequestTx, requestaccount)
	if err != nil {
		return shim.Error((methodName) + ": " + err.Error())
	}

	_, AllRequest := t.queryAllRequest(stub)
	logger.Infof("queryRequest Result : " + AllRequest)

	// t.GetAllExchangeRequests(stub)

	// str := fmt.Sprintf("%s", result)
	// logger.Infof("GetAllExchangeRequests : " + str)

	return shim.Success([]byte(args[0]))
}

func (t *Account) queryAllRequest(stub shim.ChaincodeStubInterface) (pb.Response, string) {
	methodName := "[queryRequest]"
	
		//queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"marble\",\"owner\":\"%s\"}}", owner)
		queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"request_exchange_tx\"}}")
	
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(methodName + ": " + err.Error()), "error: on getQueryResultForQueryString"
		}
		return shim.Success(nil), string(queryResults)
}

func (t *Account) GetAllExchangeRequests(stub shim.ChaincodeStubInterface) pb.Response {
	methodName := "[GetAllExchangeRequests]"
	
		//queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"marble\",\"owner\":\"%s\"}}", owner)
		queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"request_exchange_tx\"}}")
		
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(methodName + ": " + err.Error())
		}
		logger.Infof("GetAllExchangeRequests queryResults: " + string(queryResults))

		return shim.Success(queryResults)
}