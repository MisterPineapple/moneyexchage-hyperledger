package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Transaction", func() {

	var scc = new(SimpleChaincode)
	var stub = shim.NewCustomMockStub("sample", scc, adminCert)

	Describe("Exchange Request", func() {
		It("Init Chaincode", func() {
			initChaincode(stub)
		})
		It("Create Account", func() {
			var tcList = []TestCase{
				{Creator: adminCert, Key: "USER01", Negative: false},
				//{Creator: adminCert, Key: "USER02", Negative: false},
				//{Creator: adminCert, Key: "USER03", Negative: false},
				//{Creator: adminCert, Key: "USER04", Negative: false},
				//{Creator: adminCert, Key: "USER05", Negative: false},
			}

			for _, tc := range tcList {
				inputJSONStr := createAcctObj.Path(tc.Key).String()
				outputJSONStr := createAcctObj.Path(tc.Key).String()

				args := [][]byte{[]byte("create_account"), []byte(inputJSONStr)}
				invokeTransaction(stub, defaultTxID, tc.Creator, args, outputJSONStr, tc.Negative, tc.ErrorMessage)
			}
		})

		It("Mock Request Exchange", func(){
			var tcList = []TestCase{
				{Creator: adminCert, Key: "REQUEST01", Negative: false},
			}

			for _, tc := range tcList {
				inputJSONStr := exchangeRequestObj.Path(tc.Key).String()
				outputJSONStr := exchangeRequestObj.Path(tc.Key).String()

				args := [][]byte{[]byte("create_exchange_request"), []byte(inputJSONStr)}
				invokeTransaction(stub, defaultTxID, tc.Creator, args, outputJSONStr, tc.Negative, tc.ErrorMessage)
			}
		})

	})
})