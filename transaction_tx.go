package main

import (
	"encoding/json"
	//"strconv"
	//"reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type DepositMoneyTransaction struct {
	Id 				string 			`json:"id" valid:"required"`
	Balance 		[]Balance 		`json:"balance"`
}

const schemaDepositMoneyTransaction = `{
	"definitions": {},
	"$schema": "http://json-schema.org/draft-07/schema#",
	"$id": "http://example.com/root.json",
	"type": "object",
	"title": "The Root Schema",
	"required": [
		"id"
	],
	"properties": {
		"id": {
			"$id": "#/properties/id",
			"type": "string",
			"title": "The Id Schema"
		},
		"balance": {
			"$id": "#/properties/balance",
			"type": "array",
			"title": "The Balance Schema",
			"items": {
				"$id": "#/properties/balance/items",
				"type": "object",
				"title": "The Items Schema",
				"properties": {
					"currencyCode": {
						"$id": "#/properties/balance/items/properties/currencyCode",
						"type": "string",
						"title": "The CurrecnyCode Schema"
					},
					"balanceAmount": {
						"$id": "#/properties/balance/items/properties/balanceAmount",
						"type": "string",
						"title": "The BalanceAmount Schema"
					}
				}
			}
		}
	}
}`

const schemaDepositAccount = `{
	"definitions": {},
	"$schema": "http://json-schema.org/draft-07/schema#",
	"$id": "http://example.com/root.json",
	"type": "object",
	"title": "The Root Schema",
	"required": [
		"cert"
	],
	"properties": {
			"docType": {
				"$id": "#/properties/docType",
				"type": "string",
				"title": "The docType Schema"
			},
			"id": {
				"$id": "#/properties/id",
				"type": "string",
				"title": "The Id Schema"
			},
			"username": {
				"$id": "#/properties/username",
				"type": "string",
				"title": "The Username Schema"
			},
			"age": {
				"$id": "#/properties/age",
				"type": "string",
				"title": "The Age Schema"
			},
			"gender": {
				"$id": "#/properties/age",
				"type": "string",
				"title": "The Gender Schema"
			},
			"nationality": {
				"$id": "#/properties/nationality",
				"type": "string",
				"title": "The Nationality Schema"
			},
			"role": {
				"$id": "#/properties/role",
				"type": "string",
				"title": "The Role Schema"
			},
			"cert": {
				"$id": "#/properties/cert",
				"type": "string",
				"title": "The PEM encoded users x.509 certtificate"
			},
			"balance": {
				"$id": "#/properties/balance",
				"type": "array",
				"title": "The Balance Schema",
				"items": {
					"$id": "#/properties/balance/items",
					"type": "object",
					"title": "The Items Schema",
					"properties": {
						"currencyCode": {
							"$id": "#/properties/balance/items/properties/currencyCode",
							"type": "string",
							"title": "The CurrecnyCode Schema"
						},
						"balanceAmount": {
							"$id": "#/properties/balance/items/properties/balanceAmount",
							"type": "string",
							"title": "The BalanceAmount Schema"
						}
					}
				}
			}
	}
}`

func (t *Account) DepositMoneyTx(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	methodName := "[DepositMoneyTx]"
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
	//logger.Infof("valAsbytes: " + string(valAsbytes))

	var depositaccount Account
	err = ParseInput(valAsbytes, &depositaccount, schemaDepositAccount)
	if err != nil {
		return shim.Error(methodName + ": " + err.Error())
	}

	var depositMoneyTx DepositMoneyTransaction
	err = ParseInput([]byte(args[0]), &depositMoneyTx, schemaDepositMoneyTransaction)
	if err != nil {
		return shim.Error(methodName + ": " + err.Error())
	}
	
	err = t.deposit(stub, depositMoneyTx, depositaccount)
	if err != nil {
		return shim.Error((methodName) + ": " + err.Error())
	}

	// logger.Infof("depositMoneyTx length value : " + strconv.Itoa(len(depositMoneyTx.Balance))) 
	// for x := 0; x < len(depositMoneyTx.Balance); x++ {
	// 	v := reflect.ValueOf(depositMoneyTx.Balance[x])
	// 	values := make([]string, v.NumField())
	// 	for i := 0; i < v.NumField(); i=i+1 {
	// 		values[i] = v.Field(i).Interface().(string)
	// 	}
	// 	for _, j := range values {
	// 		logger.Infof(j)
	// 	}
	// 	logger.Infof("xxxxxx")
	// }

	// t.getAllAccounts(stub)
	// logger.Infof("xxxxxxxxxxxx")


	return shim.Success([]byte(args[0]))
}