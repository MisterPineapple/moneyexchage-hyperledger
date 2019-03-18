package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("moneyexchange", func() {

	var scc = new(SimpleChaincode)
	var stub = shim.NewCustomMockStub("sample", scc, adminCert)

	It("Init chaincode", func() {
		initChaincode(stub)
	})
})