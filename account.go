package main

import (

	"errors"
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// Asset Definitions 
// ============================================================================================================================

type Account struct {
	ObjectType		string			`json:"docType" valid:"required"`
	Id 				string 			`json:"id"`
	Username 		string 			`json:"username"`
	Age 			string 			`json:"age"`
	Gender 			string 			`json:"gender"`
	Nationality 	string 			`json:"nationality"`
	Role			string			`json:"role"`
	CreatorId		string			`json:"creatorId"`
	Enabled 		bool 			`json:"enabled"`
	Balance			[]Balance       `json:"balance"`
	CertID 			string 			`json:"cert" valid:"required"`
}

type Balance struct {
	CurrencyCode 	string 			`json:"currencyCode"`
	BalanceAmount 	string 			`json:"balanceAmount"`
}

//Init initializes the Parcel model/smart contract
func (t *Account) Init(stub shim.ChaincodeStubInterface) pb.Response {

	return shim.Success(nil)

}

func (t *Account) create(stub shim.ChaincodeStubInterface, createAccountTx CreateAccountTransaction) error {
	methodName := "[Account.create()]"
	objectType := "Account"

	var account Account
	account = Account{
		ObjectType:		createAccountTx.ObjectType,
		Id:				createAccountTx.Id,
		Username:		createAccountTx.Username,
		Age:			createAccountTx.Age,
		Gender:			createAccountTx.Gender,
		Nationality:	createAccountTx.Nationality,
		Role:			"user",
		CreatorId:		"creatorAccount.Id",
		Enabled:		true,
		Balance:		createAccountTx.Balance,
		CertID: 		"",
	}

	// Get Enrollment ID  to check if it's admin, need to create admin account
	enrollmentID, err := getEnrollmentID(stub)
	if err != nil {
		return errors.New(methodName + " " + err.Error())
	}

	// 1. Check if requested account already exists. If yes, and if creator is not admin@bullionist.com, reject TX.
	// antiErr := account.checkExists(stub)

	// if antiErr != nil {
	// 	if !(enrollmentID == BULLIONISTADMIN && createAccountTx.OrgRole == "BULLIONIST") {
	// 		logger.Errorf(methodName + " " + objectType + " with ID: " + createAccountTx.AccountID + " already exists")
	// 		return errors.New(methodName + " " + objectType + " with ID: " + createAccountTx.AccountID + " already exists")
	// 	}
	// }

	// 2. Extract certificate ID from the user's cert createAccountTx.Cert
	account.CertID, err = getIDFromCert(createAccountTx.Cert)
	if err != nil {
		logger.Errorf(methodName + " " + err.Error())
		return errors.New(methodName + " " + err.Error())
	}

	// 6.1 Check if account with current certID axists. If yes, and if creator is not admin@bullionist.com, reject TX.
	_, _, antiErr := account.getByCertID(stub, account.CertID)

	if antiErr == nil {
		if !(enrollmentID == ADMIN && createAccountTx.Role == "ADMIN") {
			logger.Errorf(methodName + " " + objectType + " with certId: " + account.CertID + " already exists")
			return errors.New(methodName + " " + objectType + " with certId: " + account.CertID + " already exists")
		}
	}

	// 11. Put the record into the world state
	key := account.Id

	accountJSONBytes, _ := json.Marshal(account)
	err = stub.PutState(key, accountJSONBytes)
	if err != nil {
		logger.Errorf(methodName + " PutState error: " + err.Error())
	}

	return nil
}

//GetAccountByOwner - getting account
func (t *Account) getByCertID(stub shim.ChaincodeStubInterface, ID string) (Account, []byte, error) {
	methodName := "[Account.getByCertID]"
	tmpStruct := Account{}
	tmpStructs := []Account{}
	tmpJSONStringsArray := []string{}

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Account\",\"certId\":\"%s\"}, \"use_index\":[\"_design/AccountsByCertID\", \"AccountsByCertID\"]}", ID)

	resultsIterator, err := stub.GetQueryResult(queryString)

	if err != nil {
		return tmpStruct, nil, errors.New(methodName + " Error query results:" + err.Error())
	}

	//defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	//count := 0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return tmpStruct, nil, errors.New(methodName + " " + err.Error())
		}

		tmpAccount := Account{}
		err = json.Unmarshal(queryResponse.Value, &tmpAccount)
		if err != nil {
			//error unmarshaling
			return tmpStruct, nil, errors.New(methodName + " Error unmarshaling Account:" + err.Error())
		}
		tmpStructs = append(tmpStructs, tmpAccount)
		// Record is a JSON object, so we write as-is
		normalizedJSON, err := json.Marshal(tmpAccount)
		if err != nil {
			//error marshaling
			return tmpStruct, nil, errors.New(methodName + " Error marshaling:" + err.Error())
		}
		tmpJSONStringsArray = append(tmpJSONStringsArray, string(normalizedJSON))
		//count++
	}

	if len(tmpJSONStringsArray) < 1 {
		return tmpStruct, nil, errors.New(methodName + " Looks like account with certificate ID: " + ID + " doesn't exist.")
	}

	//TODO: Delete me
	//logger.Debugf(methodName + " Found Account: '" + tmpJSONStringsArray[0] + "'")

	// if count == 0 {
	// 	return tmpStruct, nil, errors.New(methodName + " Looks like account with certificate ID: " + ID + " doesn't exist.")
	// }

	if len(tmpJSONStringsArray) > 1 {
		return tmpStruct, nil, errors.New(methodName + " Looks like there are more than one account with certificate ID: " + ID)
	}

	if []byte(tmpJSONStringsArray[0]) == nil {
		//logger.Errorf(methodName + " No Account with CertID: " + ID + " is found")
		return tmpStruct, nil, errors.New(methodName + " No Account with " + ID + " is found")
	}

	tmpStruct = tmpStructs[0]

	//tmpJSONBytes := []byte("[" + strings.Join(tmpJSONStringsArray, ",") + "]")

	return tmpStruct, []byte(tmpJSONStringsArray[0]), nil
}