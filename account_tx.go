package main

import (
	 "fmt"
	 "encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CreateAccountTransaction - tx to create account, value are inputted
type CreateAccountTransaction struct {
	ObjectType	string 	  `json:"docType" valid:"required"`
	Id         	string    `json:"id"`
	Username   	string    `json:"username"`
	Age      	string    `json:"age"`
	Gender      string 	  `json:"gender"`
	Nationality string    `json:"nationality"`
	Role		string	  `json:"role"`
	Balance		[]Balance `json:"balance"`
	Cert 		string 	  `json:"cert" valid:"required"`
}

type GetAccountById struct{
	Id			string	  `json:"id"`
}

const schemaCreateAccountTransaction = `{
	"definitions": {},
	"$schema": "http://json-schema.org/draft-07/schema#",
	"$id": "http://example.com/root.json",
	"type": "object",
	"title": "The Root Schema",
	"required": [
		"cert"
	],
	"properties": {
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

func (t *Account) CreateAccountTx(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	methodName := "[CreateAccountTX]"

	enrollmentID, err := getEnrollmentID(stub)
	if err != nil {
	  return shim.Error(methodName + " " + err.Error())
	}
	logger.Debugf(methodName + " Creator enrollmentID: " + enrollmentID)

	var createAccountTx CreateAccountTransaction
	createAccountTx.ObjectType = "account"
	err = ParseInput([]byte(args[0]), &createAccountTx, schemaCreateAccountTransaction)
	if err != nil {
		return shim.Error(methodName + ": " + err.Error())
	}

	//Check ACL if we are not Super Admin
	// if !(enrollmentID == ADMIN && createAccountTx.Role == "admin") {
	// 	return shim.Error(methodName + " Access denied:" + err.Error())
	// }

	err = t.create(stub, createAccountTx)
	if err != nil {
		return shim.Error(methodName + ": " + err.Error())
	}

	return shim.Success([]byte(args[0]))
}

func (t *Account) GetAccountById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	methodName := "[Account.getAccountByID]"
	//tmpStruct := Account{}
	tmpStructs := []Account{}
	tmpJSONStringsArray := []string{}

	value := args[0]
	var getAccountByIdTx GetAccountById
	err := json.Unmarshal([]byte(value), &getAccountByIdTx)
	ID := getAccountByIdTx.Id

	//queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"marble\",\"owner\":\"%s\"}}", owner)
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"account\",\"id\":\"%s\"}}", ID)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(methodName + " Error query results:" + err.Error() + "  id: " + ID)
	}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(methodName + ": " + err.Error())
		}

		tmpAccount := Account{}
		err = json.Unmarshal(queryResponse.Value, &tmpAccount)
		if err != nil {
			//error unmarshaling
			return shim.Error(methodName + ": " + err.Error())
		}
		tmpStructs = append(tmpStructs, tmpAccount)
		// Record is a JSON object, so we write as-is
		normalizedJSON, err := json.Marshal(tmpAccount)
		if err != nil {
			//error marshaling
			return shim.Error(methodName + ": " + err.Error())
		}
		tmpJSONStringsArray = append(tmpJSONStringsArray, string(normalizedJSON))
		//count++
	}

	if len(tmpJSONStringsArray) < 1 {
		return shim.Error(methodName + " Looks like account with account ID: " + ID + " doesn't exist.")
	}

	if len(tmpJSONStringsArray) > 1 {
		return shim.Error(methodName + " Looks like there are more than one account with account ID: " + ID)
	}

	if []byte(tmpJSONStringsArray[0]) == nil {
		//logger.Errorf(methodName + " No Account with CertID: " + ID + " is found")
		return shim.Error(methodName + " No Account with " + ID + " is found")
	}

	return shim.Success([]byte(tmpJSONStringsArray[0]))
}

func (t *Account) getAllAccounts(stub shim.ChaincodeStubInterface) pb.Response {
	methodName := "[queryRequest]"
	
		//queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"marble\",\"owner\":\"%s\"}}", owner)
		queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"account\"}}")
	
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		if err != nil {
			return shim.Error(methodName + ": " + err.Error())
		}
		logger.Infof("queryResults: " + string(queryResults))

		return shim.Success(queryResults)
}