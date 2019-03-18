package main

import (
	"encoding/json"
	"reflect"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	//pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *Account) deposit(stub shim.ChaincodeStubInterface, depositMoneyTx DepositMoneyTransaction, depositAccount Account) error {
	methodName := "[deposit]"

	var depositaccount Account
	depositaccount = Account{
		ObjectType:		depositAccount.ObjectType,		 
		Id:				depositAccount.Id,
		Username:		depositAccount.Username,
		Age:			depositAccount.Age,
		Gender:			depositAccount.Gender,
		Nationality:	depositAccount.Nationality,
		Role:			depositAccount.Role,
		CreatorId:		depositAccount.CreatorId,
		Enabled:		depositAccount.Enabled,
		Balance:		depositAccount.Balance,
		CertID: 		depositAccount.CertID,
	}

	logger.Infof("depositaccount[deposit] : " + depositaccount.Id)
	logger.Infof("depositMoneyTx[deposit] : " + depositMoneyTx.Id)
	//logger.Infof("depositMoneyTx[Balance] : " + depositMoneyTx.Balance)
	for i := 0; i < len(depositMoneyTx.Balance); i++ {
		v := reflect.ValueOf(depositMoneyTx.Balance[i])
		values := make([]string, v.NumField())
		for x := 0; x < v.NumField(); x++ {
			values[x] = v.Field(x).Interface().(string)
		}
		currencyCode := values[0] 
		//logger.Infof("Currency Code : " + currencyCode)
			
		switch currencyCode {
		case "USD":
				index := 0
				depositAccountBalance, _ := strconv.Atoi(depositaccount.Balance[index].BalanceAmount)
				depositAmount, _ := strconv.Atoi(depositMoneyTx.Balance[i].BalanceAmount)
				newUSDamount :=  depositAccountBalance + depositAmount
				depositaccount.Balance[index].BalanceAmount = strconv.Itoa(newUSDamount)
				//logger.Infof("US : " + strconv.Itoa(newUSDamount))
		case "SGD":
				index := 1
				depositAccountBalance, _ := strconv.Atoi(depositaccount.Balance[index].BalanceAmount)
				depositAmount, _ := strconv.Atoi(depositMoneyTx.Balance[i].BalanceAmount)
				newSGDamount := depositAccountBalance + depositAmount
				depositaccount.Balance[index].BalanceAmount = strconv.Itoa(newSGDamount)
				//logger.Infof("SGD : " + strconv.Itoa(newSGDamount))
		case "THB":
				index := 2
				depositAccountBalance, _ := strconv.Atoi(depositaccount.Balance[index].BalanceAmount)
				depositAmount, _ := strconv.Atoi(depositMoneyTx.Balance[i].BalanceAmount)
				newTHBamount := depositAccountBalance + depositAmount
				depositaccount.Balance[index].BalanceAmount = strconv.Itoa(newTHBamount)
				//logger.Infof("THB : " + strconv.Itoa(newTHBamount))
		case "CNY":
				index := 3
				depositAccountBalance, _ := strconv.Atoi(depositaccount.Balance[index].BalanceAmount)
				depositAmount, _ := strconv.Atoi(depositMoneyTx.Balance[i].BalanceAmount)
				newCNYamount := depositAccountBalance + depositAmount
				depositaccount.Balance[index].BalanceAmount = strconv.Itoa(newCNYamount)
				//logger.Infof("CNY : " + strconv.Itoa(newCNYamount))
		case "JPY":
				index := 4
				depositAccountBalance, _ := strconv.Atoi(depositaccount.Balance[index].BalanceAmount)
				depositAmount, _ := strconv.Atoi(depositMoneyTx.Balance[i].BalanceAmount)
				newJPYamount := depositAccountBalance + depositAmount
				depositaccount.Balance[index].BalanceAmount = strconv.Itoa(newJPYamount)
				//logger.Infof("JPY : " + strconv.Itoa(newJPYamount))
		default:
			logger.Errorf(methodName + " Currency rather than USD, SGD, THB, CNY, JPY is not allow")
			return errors.New(methodName + " Currency rather than USD, SGD, THB, CNY, JPY is not allow")
		}	
	}

	depositaccount.ObjectType = "account"
	key := depositaccount.Id

	accountJSONBytes, _ := json.Marshal(depositaccount)
	err := stub.PutState(key, accountJSONBytes)
	if err != nil {
		logger.Errorf(methodName + " PutState error: " + err.Error())
	}

	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
		logger.Infof(methodName + " " + jsonResp)
	}

	logger.Infof("Get State :" + string(valAsbytes))

	// for x := 0; x < len(depositaccount.Balance); x++ {
	// 	v := reflect.ValueOf(depositaccount.Balance[x])
	// 	values := make([]string, v.NumField())
	// 	for i := 0; i < v.NumField(); i=i+1 {
	// 		values[i] = v.Field(i).Interface().(string)
	// 	}
	// 	for _, j := range values {
	// 		logger.Infof(j)
	// 	}
	// 	logger.Infof("xxxxxx")
	// }

	return nil

}