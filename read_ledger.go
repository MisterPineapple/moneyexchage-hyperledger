package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func read(stub shim.ChaincodeStubInterface, args []string ) pb.Response {
	var key,jsonResp string

	fmt.Println("-- Reading --")

	if len(args) != 1 {
		return shim.Error("Error args length is input incorrectly")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil{
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("- end read")
	return shim.Success(valAsbytes) 
}