package main

import (

	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("simple-chaincode")

// SimpleChaincode -
type SimpleChaincode struct {
	account		Account
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode - %s", err)
	}
}

// ============================================================================================================================
// Invoke
// ============================================================================================================================

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// methodName := "[Invoke]"
	// invokerAccount := Account{}

	function, args := stub.GetFunctionAndParameters()
	logger.Infof("Invoke is running " + function)

	if function == "init" {
		return t.Init(stub)
	} else if function == "read" {
		return read(stub, args)
	} else if function == "create_account" {
		return t.account.CreateAccountTx(stub, args)
	} else if function == "get_all_accounts" {
		return t.account.getAllAccounts(stub)
	} else if function == "get_account_by_id" {
		return t.account.GetAccountById(stub, args)
	}else if function == "deposit_money" {
		return t.account.DepositMoneyTx(stub, args)
	} else if function == "create_exchange_request" {
		return t.account.ExchangeRequestTx(stub, args)
	} else if function == "get_all_exchange_requests" {
		return t.account.GetAllExchangeRequests(stub)
	}

	return shim.Error("Unknown invoke function: " + function)
}

// ============================================================================================================================
// Initilize
// ============================================================================================================================

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	t.account.Init(stub)

	return shim.Success(nil)
}