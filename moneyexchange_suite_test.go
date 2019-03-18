package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Jeffail/gabs"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// set of certs
var adminCert = readMockFile("./mock/cert/admin")

const defaultTxID = "169e1e83e930853391bc6f35f605c6754cfead57cf8387639d3b4096c54f18f4"

var createAcctByte, _ = ioutil.ReadFile("./mock/dummy-create-account.json")
var createAcctObj, _ = gabs.ParseJSON(createAcctByte)

var depositMoneyByte, _ = ioutil.ReadFile("./mock/deposit-money.json")
var depositMoneyObj, _ = gabs.ParseJSON(depositMoneyByte)

var exchangeRequestByte, _ = ioutil.ReadFile("./mock/exchange-request.json")
var exchangeRequestObj, _ = gabs.ParseJSON(exchangeRequestByte)

type TestCase struct {
	Creator      []byte
	Key          string
	Negative     bool
	ErrorMessage string
}


func TestMoneyexchange(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moneyexchange Suite")
}

func initChaincode(stub *shim.CustomMockStub) {
	res := stub.MockInit("1", [][]byte{})

	Expect(res.Status).To(Equal(int32(shim.OK)))
	Expect(res.Payload).To(Equal([]byte(nil)))
}

func invokeTransaction(stub *shim.CustomMockStub, txID string, creator []byte, args [][]byte, outputJSONStr string, negative bool, errorMsg string) {
	stub.SetCreator(creator)
	res := stub.MockInvoke(txID, args)

	// fmt.Printf("Test Case: args: %s Negative: %t res.Message: %s\n", args, negative, res.Message)
	if negative {
		// fmt.Println("res.Message", res.Message)
		Expect(res.Status).To(Equal(int32(shim.ERROR)))
		Expect(strings.ToLower(res.Message)).To(ContainSubstring(strings.ToLower(errorMsg)))
	} else {
		payloadJSONParsed, _ := gabs.ParseJSON([]byte(res.Payload))
		fmt.Println("Payload      ", payloadJSONParsed.String())
		fmt.Println("outputJSONStr", outputJSONStr)
		Expect(res.Status).To(Equal(int32(shim.OK)))
		Expect(payloadJSONParsed.String()).To(Equal(outputJSONStr))
	}
}

func readMockFile(fileName string) []byte {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return raw
}