package main

import (
	"encoding/json"
	"strconv"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type QueryResult struct {
	Key			string 								`json:"Key"`
	Record		ExchangeRequestTransaction			`json:"Record"`
}


func (t *Account) exchangeRequest(stub shim.ChaincodeStubInterface, exchangeRequest ExchangeRequestTransaction, requestAccount Account) error {
	methodName := "[exchangeRequest]"

	USDSGD := 1.35
	USDTHB := 31.70
	USDCNY := 6.70
	USDJPY := 111.85

	SGDUSD := 0.73
	SGDTHB := 23.43
	SGDCNY := 4.95
	SGDJPY := 82.70

	THBUSD := 0.03
	THBSGD := 0.04
	THBCNY := 0.21
	THBJPY := 3.52

	CNYUSD := 0.15
	CNYSGD := 0.20
	CNYTHB := 4.73
	CNYJPY := 16.68

	JPYUSD := 0.01
	JPYSGD := 0.01
	JPYTHB := 0.28
	JPYCNY := 0.06

	var hasTarget bool = false

	var requestaccount Account 
	var targetaccount Account

	requestaccount = Account{		
		ObjectType:		requestAccount.ObjectType, 
		Id:				requestAccount.Id,
		Username:		requestAccount.Username,
		Age:			requestAccount.Age,
		Gender:			requestAccount.Gender,
		Nationality:	requestAccount.Nationality,
		Role:			requestAccount.Role,
		CreatorId:		requestAccount.CreatorId,
		Enabled:		requestAccount.Enabled,
		Balance:		requestAccount.Balance,
		CertID: 		requestAccount.CertID,
	}
	//retrieve latest request
	_, AllRequest := t.queryAllRequest(stub)
	// logger.Infof("queryRequest Result : " + AllRequest)

	//check each request if match with this request
	allRequest := AllRequest
    var dataAsStruct []QueryResult
	err := json.Unmarshal([]byte(allRequest), &dataAsStruct)

	exchangeRequest.ObjectType = "request_exchange_tx"
	mock_key := exchangeRequest.RequestId
	if mock_key == "MOCK000" {

		exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
		err = stub.PutState(mock_key, exchangeRequestJSONBytes)
		if err != nil {
			logger.Errorf(methodName + " PutState error: " + err.Error())
		}
	}

	for _, eachRequest := range dataAsStruct {
		switch exchangeRequest.Origin {
		case "USD":
			switch exchangeRequest.Target {
			case "SGD":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "SGD" && eachRequest.Record.Target == "USD" && exchangeRequest_OriginAmount * USDSGD == eachRequest_OriginAmount) {

					origin_index := 0
					target_index := 1

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			case "THB":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "THB" && eachRequest.Record.Target == "USD" && exchangeRequest_OriginAmount * USDTHB == eachRequest_OriginAmount) {

					origin_index := 0
					target_index := 2

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			case "CNY":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "CNY" && eachRequest.Record.Target == "USD" && exchangeRequest_OriginAmount * USDCNY == eachRequest_OriginAmount) {

					origin_index := 0
					target_index := 3

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			case "JPY":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "JPY" && eachRequest.Record.Target == "USD" && exchangeRequest_OriginAmount * USDJPY == eachRequest_OriginAmount) {

					origin_index := 0
					target_index := 4

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			}
		case "SGD":
			switch exchangeRequest.Target {
			 case "USD":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "USD" && eachRequest.Record.Target == "SGD" && exchangeRequest_OriginAmount * SGDUSD == eachRequest_OriginAmount) {

					origin_index := 1
					target_index := 0

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "THB":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "THB" && eachRequest.Record.Target == "SGD" && exchangeRequest_OriginAmount * SGDTHB == eachRequest_OriginAmount) {

					origin_index := 1
					target_index := 2

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "CNY":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "CNY" && eachRequest.Record.Target == "SGD" && exchangeRequest_OriginAmount * SGDCNY == eachRequest_OriginAmount) {

					origin_index := 1
					target_index := 3

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "JPY":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "JPY" && eachRequest.Record.Target == "SGD" && exchangeRequest_OriginAmount * SGDJPY == eachRequest_OriginAmount) {

					origin_index := 1
					target_index := 4

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			}
		case "THB":
		 	switch exchangeRequest.Target {
			 case "USD":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "USD" && eachRequest.Record.Target == "THB" && exchangeRequest_OriginAmount * THBUSD == eachRequest_OriginAmount) {

					origin_index := 2
					target_index := 0

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "SGD":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "SGD" && eachRequest.Record.Target == "THB" && exchangeRequest_OriginAmount * THBSGD == eachRequest_OriginAmount) {

					origin_index := 2
					target_index := 1

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "CNY":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "CNY" && eachRequest.Record.Target == "THB" && exchangeRequest_OriginAmount * THBCNY == eachRequest_OriginAmount) {

					origin_index := 2
					target_index := 3

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "JPY":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "JPY" && eachRequest.Record.Target == "THB" && exchangeRequest_OriginAmount * THBJPY == eachRequest_OriginAmount) {

					origin_index := 2
					target_index := 4

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			}
		case "CNY":
		 	switch exchangeRequest.Target {
			 case "USD":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "USD" && eachRequest.Record.Target == "CNY" && exchangeRequest_OriginAmount * CNYUSD == eachRequest_OriginAmount) {

					origin_index := 3
					target_index := 0

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "SGD":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "SGD" && eachRequest.Record.Target == "CNY" && exchangeRequest_OriginAmount * CNYSGD == eachRequest_OriginAmount) {

					origin_index := 3
					target_index := 1

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			case "THB":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "THB" && eachRequest.Record.Target == "CNY" && exchangeRequest_OriginAmount * CNYTHB == eachRequest_OriginAmount) {

					origin_index := 3
					target_index := 2

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "JPY":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "JPY" && eachRequest.Record.Target == "CNY" && exchangeRequest_OriginAmount * CNYJPY == eachRequest_OriginAmount) {

					origin_index := 3
					target_index := 4

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			}
		case "JPY":
		 	switch exchangeRequest.Target {
			 case "USD":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "USD" && eachRequest.Record.Target == "JPY" && exchangeRequest_OriginAmount * JPYUSD == eachRequest_OriginAmount) {

					origin_index := 4
					target_index := 0

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "SGD":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "SGD" && eachRequest.Record.Target == "JPY" && exchangeRequest_OriginAmount * JPYSGD == eachRequest_OriginAmount) {

					origin_index := 4
					target_index := 1

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "THB":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "THB" && eachRequest.Record.Target == "JPY" && exchangeRequest_OriginAmount * JPYTHB == eachRequest_OriginAmount) {

					origin_index := 4
					target_index := 2

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			 case "CNY":
				exchangeRequest_OriginAmount, _ := strconv.ParseFloat(exchangeRequest.OriginAmount, 64)
				eachRequest_OriginAmount, _ := strconv.ParseFloat(eachRequest.Record.OriginAmount, 64)
				if (eachRequest.Record.Origin == "CNY" && eachRequest.Record.Target == "JPY" && exchangeRequest_OriginAmount * JPYCNY == eachRequest_OriginAmount) {

					origin_index := 4
					target_index := 3

					targetaccount = getTargetAccount(stub, eachRequest)
					if err != nil {
						logger.Errorf(methodName + " getTargetAccount error: " + err.Error())
					}

					err = requestMatch(stub, requestaccount, targetaccount, eachRequest, exchangeRequest, origin_index, target_index, exchangeRequest_OriginAmount, eachRequest_OriginAmount)
					if err != nil {
						logger.Errorf(methodName + " RequestMatch error: " + err.Error())
					}

					hasTarget = true
				} else {
					requestexchange_key := exchangeRequest.RequestId

					exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
					err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
					if err != nil {
						logger.Errorf(methodName + " PutState error: " + err.Error())
					}
				}
			}
		default: 
			// requestexchange_key := exchangeRequest.RequestId

			// exchangeRequestJSONBytes, _ := json.Marshal(exchangeRequest)
			// err = stub.PutState(requestexchange_key, exchangeRequestJSONBytes)
			// if err != nil {
			// 	logger.Errorf(methodName + " PutState error: " + err.Error())
			// }
			logger.Errorf(methodName + " Currency rather than USD, SGD, THB, CNY, JPY is not allow")
			return errors.New(methodName + " Currency rather than USD, SGD, THB, CNY, JPY is not allow")
		}
	}

	//putstate of this account and target account
	if hasTarget {
		requestaccount_key := requestaccount.Id
		logger.Infof("requestaccount.Id : " + requestaccount.Id)

		requestaccountJSONBytes, _ := json.Marshal(requestaccount)
		err = stub.PutState(requestaccount_key, requestaccountJSONBytes)
		if err != nil {
			logger.Errorf(methodName + " PutState error: " + err.Error())
		}
	
		targetaccount_key := targetaccount.Id
		logger.Infof("targetaccount.Id : " + targetaccount.Id)

		targetaccountJSONBytes, _ := json.Marshal(targetaccount)
		err = stub.PutState(targetaccount_key, targetaccountJSONBytes)
		if err != nil {
			logger.Errorf(methodName + " PutState error: " + err.Error())
		}
		/////////////
		valAsbytes, err := stub.GetState(targetaccount_key)
		if err != nil {
			logger.Infof("GetState Error")
		}

		err = ParseInput(valAsbytes, &targetaccount, schemaDepositAccount)
		if err != nil {
			//return shim.Error(methodName + ": " + err.Error())
		}
		// logger.Infof("----------requestaccount info----------")
		// logger.Infof("requestaccount info (Id) : " + requestaccount.Id)
		// logger.Infof("requestaccount info (Username) : " + requestaccount.Username)
		// for _, v := range requestaccount.Balance {
		// 	logger.Infof("requestaccount info (Balance[CurrencyCode]) : " + v.CurrencyCode)
		// 	logger.Infof("requestaccount info (Balance[BalanceAmount]) : " + v.BalanceAmount)
		// }
	}

	// logger.Infof("exchangerequest info (ObjectType) : " + exchangeRequest.ObjectType)
	// logger.Infof("exchangerequest info (Id) : " + exchangeRequest.Id)
	// logger.Infof("exchangerequest info (ResquestId) : " + exchangeRequest.RequestId)
	// logger.Infof("exchangerequest info (Origin) : " + exchangeRequest.Origin)
	// logger.Infof("exchangerequest info (Target) : " + exchangeRequest.Target)
	// logger.Infof("exchangerequest info (OriginAmount) : " + exchangeRequest.OriginAmount)
	// logger.Infof("exchangerequest info (TargetAmount) : " + exchangeRequest.TargetAmount)


	return nil
}

func requestMatch(stub shim.ChaincodeStubInterface, 
	RequestAccount Account, 
	TargetAccount Account, 
	EachRequest QueryResult, 
	ExchangeRequest ExchangeRequestTransaction, 
	origin_index int, target_index int, 
	exchangeRequest_OriginAmount float64, 
	eachRequest_OriginAmount float64) error {
	methodName := "[RequestMatch]"

	////////// update Request account
	valAsbytes, _ := stub.GetState(ExchangeRequest.Id)
	err := ParseInput(valAsbytes, &TargetAccount, schemaDepositAccount)
	if err != nil {
		return errors.New(methodName + "Parse Input error")
	}
	RequestAccount_OriginBalance, _ := strconv.ParseFloat(RequestAccount.Balance[origin_index].BalanceAmount, 64)
	newRequestAccount_OriginBalance := fmt.Sprintf("%.2f", RequestAccount_OriginBalance - exchangeRequest_OriginAmount)

	RequestAccount_TargetBalance, _ := strconv.ParseFloat(RequestAccount.Balance[target_index].BalanceAmount, 64)
	newRequestAccount_TargetBalance := fmt.Sprintf("%.2f", RequestAccount_TargetBalance + eachRequest_OriginAmount)

	RequestAccount.Balance[origin_index].BalanceAmount = newRequestAccount_OriginBalance
	RequestAccount.Balance[target_index].BalanceAmount = newRequestAccount_TargetBalance
	////////// update Target account
	valAsbytes, _ = stub.GetState(EachRequest.Record.Id)
	err = ParseInput(valAsbytes, &TargetAccount, schemaDepositAccount)
	if err != nil {
		return errors.New(methodName + "Parse Input error")
	}
	TargetAccount_TargetBalance, _ := strconv.ParseFloat(TargetAccount.Balance[origin_index].BalanceAmount, 64)
	newTargetAccount_TargetBalance := fmt.Sprintf("%.2f", TargetAccount_TargetBalance + exchangeRequest_OriginAmount)

	TargetAccount_OriginBalance, _ := strconv.ParseFloat(TargetAccount.Balance[target_index].BalanceAmount, 64)
	newTargetAccount_OriginBalance := fmt.Sprintf("%.2f", TargetAccount_OriginBalance - eachRequest_OriginAmount)

	TargetAccount.Balance[origin_index].BalanceAmount = newTargetAccount_TargetBalance
	TargetAccount.Balance[target_index].BalanceAmount = newTargetAccount_OriginBalance

	//delete Target account's request
	eachRequest_key := EachRequest.Record.RequestId
	err = stub.DelState(eachRequest_key)
	if err != nil {
		return errors.New(methodName + "DelState error")
	}
	//delete Request account's request
	exchangeRequest_key := ExchangeRequest.RequestId
	err = stub.DelState(exchangeRequest_key)
	if err != nil {
		return errors.New(methodName + "DelState error")
	}

	return nil
}

func getTargetAccount(stub shim.ChaincodeStubInterface, EachRequest QueryResult) Account {
	methodName := "[getTargetAccount]"
	key := EachRequest.Record.Id

	valAsbytes, err := stub.GetState(key)
	if err != nil {
		logger.Infof(methodName + " valAsbytes error!")
	}

	var target_account Account
	err = ParseInput(valAsbytes, &target_account, schemaDepositAccount)
	if err != nil {
		logger.Infof(methodName + " ParseInput error!")
	}

	return target_account
}