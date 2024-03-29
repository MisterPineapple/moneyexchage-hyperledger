package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Transaction", func() {

	var scc = new(SimpleChaincode)
	var stub = shim.NewCustomMockStub("sample", scc, adminCert)

	Describe("Deposit Money", func() {

		It("Init Chaincode", func() {
			initChaincode(stub)
		})
		It("Create Account", func() {
			var tcList = []TestCase{
				{Creator: adminCert, Key: "USER01", Negative: false},
				{Creator: adminCert, Key: "USER02", Negative: false},
				{Creator: adminCert, Key: "USER03", Negative: false},
				{Creator: adminCert, Key: "USER04", Negative: false},
				//{Creator: adminCert, Key: "USER05", Negative: false},
			}

			for _, tc := range tcList {
				inputJSONStr := createAcctObj.Path(tc.Key).String()
				outputJSONStr := createAcctObj.Path(tc.Key).String()

				args := [][]byte{[]byte("create_account"), []byte(inputJSONStr)}
				invokeTransaction(stub, defaultTxID, tc.Creator, args, outputJSONStr, tc.Negative, tc.ErrorMessage)
			}
		})

		It("Deposit Money", func() {
			var tcList = []TestCase{
				{Creator: adminCert, Key: "DEPOSIT01", Negative: false},
				{Creator: adminCert, Key: "DEPOSIT02", Negative: false},
				{Creator: adminCert, Key: "DEPOSIT03", Negative: false},
				{Creator: adminCert, Key: "DEPOSIT04", Negative: false},
				//{Creator: adminCert, Key: "DEPOSIT05", Negative: true, ErrorMessage: "Access denied"},
				//{Creator: adminCert, Key: "USER04", Negative: true, ErrorMessage: "Access denied"},
			}

			for _, tc := range tcList {
				inputJSONStr := depositMoneyObj.Path(tc.Key).String()
				outputJSONStr := depositMoneyObj.Path(tc.Key).String()

				args := [][]byte{[]byte("deposit_money"), []byte(inputJSONStr)}
				invokeTransaction(stub, defaultTxID, tc.Creator, args, outputJSONStr, tc.Negative, tc.ErrorMessage)
			}
		})
	})
})