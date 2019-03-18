/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package shim provides APIs for the chaincode to access its state
// variables, transaction context and call other chaincodes.
package shim

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	pb "github.com/hyperledger/fabric/protos/peer"
	queryParser "github.com/wearetheledger/go-couchdb-query-engine/query"
)

// CustomMockStub -
type CustomMockStub struct {
	stub         *MockStub
	creator      []byte
	filteredStub *MockStub
}

// GetTxID -
func (mock *CustomMockStub) GetTxID() string {
	/*Code to generate transaction id
	timestamp := time.Now().Unix()
	rand.Seed(timestamp)
	hash := sha256.Sum256([]byte(string(timestamp)))
	s := hex.EncodeToString(hash[:])
	*/
	return mock.stub.GetTxID()
}

func (mock *CustomMockStub) GetChannelID() string {
	return mock.stub.GetChannelID()
}

func (mock *CustomMockStub) GetArgs() [][]byte {
	return mock.stub.GetArgs()
}

func (mock *CustomMockStub) GetStringArgs() []string {
	return mock.stub.GetStringArgs()
}

func (mock *CustomMockStub) GetFunctionAndParameters() (function string, params []string) {
	return mock.stub.GetFunctionAndParameters()
}

// Used to indicate to a chaincode that it is part of a transaction.
// This is important when chaincodes invoke each other.
// MockStub doesn't support concurrent transactions at present.
func (mock *CustomMockStub) MockTransactionStart(txid string) {
	mock.stub.MockTransactionStart(txid)
}

// End a mocked transaction, clearing the UUID.
func (mock *CustomMockStub) MockTransactionEnd(uuid string) {
	mock.stub.MockTransactionEnd(uuid)
}

// Register a peer chaincode with this MockStub
// invokableChaincodeName is the name or hash of the peer
// otherStub is a MockStub of the peer, already intialised
func (mock *CustomMockStub) MockPeerChaincode(invokableChaincodeName string, otherStub *MockStub) {
	mock.stub.MockPeerChaincode(invokableChaincodeName, otherStub)
}

// Initialise this chaincode,  also starts and ends a transaction.
func (mock *CustomMockStub) MockInit(uuid string, args [][]byte) pb.Response {
	mock.stub.args = args
	mock.stub.MockTransactionStart(uuid)
	res := mock.stub.cc.Init(mock)
	mock.stub.MockTransactionEnd(uuid)
	return res
}

// Invoke this chaincode, also starts and ends a transaction.
func (mock *CustomMockStub) MockInvoke(uuid string, args [][]byte) pb.Response {
	mock.stub.args = args
	mock.MockTransactionStart(uuid)
	res := mock.stub.cc.Invoke(mock)
	mock.MockTransactionEnd(uuid)
	return res
}

func (stub *CustomMockStub) GetDecorations() map[string][]byte {
	return nil
}

// Invoke this chaincode, also starts and ends a transaction.
func (mock *CustomMockStub) MockInvokeWithSignedProposal(uuid string, args [][]byte, sp *pb.SignedProposal) pb.Response {
	return mock.stub.MockInvokeWithSignedProposal(uuid, args, sp)
}

// GetState retrieves the value for a given key from the ledger
func (mock *CustomMockStub) GetState(key string) ([]byte, error) {
	return mock.stub.GetState(key)
}

// PutState writes the specified `value` and `key` into the ledger.
func (mock *CustomMockStub) PutState(key string, value []byte) error {
	mock.stub.PutState(key, value)
	return nil
}

// DelState removes the specified `key` and its value from the ledger.
func (mock *CustomMockStub) DelState(key string) error {
	mock.stub.DelState(key)
	return nil
}

func (mock *CustomMockStub) GetStateByRange(startKey, endKey string) (StateQueryIteratorInterface, error) {
	return mock.stub.GetStateByRange(startKey, endKey)
}

// GetQueryResult function can be invoked by a chaincode to perform a
// rich query against state database.  Only supported by state database implementations
// that support rich query.  The query string is in the syntax of the underlying
// state database. An iterator is returned which can be used to iterate (next) over
// the query result set
// GetQueryResult function can be invoked by a chaincode to perform a
// rich query against state database.  Only supported by state database implementations
// that support rich query.  The query string is in the syntax of the underlying
// state database. An iterator is returned which can be used to iterate (next) over
// the query result set
func (mock *CustomMockStub) GetQueryResult(query string) (StateQueryIteratorInterface, error) {
	var interfaces = make(map[string]interface{})

	for k, v := range mock.stub.State {

		var asInterface interface{}

		json.Unmarshal(v, &asInterface)

		interfaces[k] = asInterface

	}

	result, err := queryParser.ParseCouchDBQueryString(interfaces, query)

	if err != nil {
		return nil, err
	}

	return StateQueryIteratorInterface(NewMockQueryIterator(mock.stub, &result)), nil
}

// GetHistoryForKey function can be invoked by a chaincode to return a history of
// key values across time. GetHistoryForKey is intended to be used for read-only queries.
func (mock *CustomMockStub) GetHistoryForKey(key string) (HistoryQueryIteratorInterface, error) {
	return nil, errors.New("Not Implemented")
}

//GetStateByPartialCompositeKey function can be invoked by a chaincode to query the
//state based on a given partial composite key. This function returns an
//iterator which can be used to iterate over all composite keys whose prefix
//matches the given partial composite key. This function should be used only for
//a partial composite key. For a full composite key, an iter with empty response
//would be returned.
func (mock *CustomMockStub) GetStateByPartialCompositeKey(objectType string, attributes []string) (StateQueryIteratorInterface, error) {
	return mock.stub.GetStateByPartialCompositeKey(objectType, attributes)
}

// CreateCompositeKey combines the list of attributes
//to form a composite key.
func (mock *CustomMockStub) CreateCompositeKey(objectType string, attributes []string) (string, error) {
	return mock.stub.CreateCompositeKey(objectType, attributes)
}

// SplitCompositeKey splits the composite key into attributes
// on which the composite key was formed.
func (mock *CustomMockStub) SplitCompositeKey(compositeKey string) (string, []string, error) {
	return mock.stub.SplitCompositeKey(compositeKey)
}

// InvokeChaincode calls a peered chaincode.
// E.g. stub1.InvokeChaincode("stub2Hash", funcArgs, channel)
// Before calling this make sure to create another MockStub stub2, call stub2.MockInit(uuid, func, args)
// and register it with stub1 by calling stub1.MockPeerChaincode("stub2Hash", stub2)
func (mock *CustomMockStub) InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response {
	// Internally we use chaincode name as a composite name
	return mock.stub.InvokeChaincode(chaincodeName, args, channel)
}

// Not implemented
func (mock *CustomMockStub) GetCreator() ([]byte, error) {
	return mock.creator, nil
}

// Not implemented
func (mock *CustomMockStub) GetTransient() (map[string][]byte, error) {
	return nil, nil
}

// Not implemented
func (mock *CustomMockStub) GetBinding() ([]byte, error) {
	return nil, nil
}

// Not implemented
func (mock *CustomMockStub) GetSignedProposal() (*pb.SignedProposal, error) {
	return mock.stub.GetSignedProposal()
}

func (mock *CustomMockStub) setSignedProposal(sp *pb.SignedProposal) {
	mock.stub.setSignedProposal(sp)
}

// Not implemented
func (mock *CustomMockStub) GetArgsSlice() ([]byte, error) {
	return nil, nil
}

func (mock *CustomMockStub) setTxTimestamp(time *timestamp.Timestamp) {
	mock.stub.setTxTimestamp(time)
}

func (mock *CustomMockStub) GetTxTimestamp() (*timestamp.Timestamp, error) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	t, _ := time.Parse(layout, str)
	// now := time.Now().UTC()
	secs := t.Unix()
	nanos := int32(t.UnixNano() - (secs * 1000000000))
	return &(timestamp.Timestamp{Seconds: secs, Nanos: nanos}), nil
}

// Not implemented
func (mock *CustomMockStub) SetEvent(name string, payload []byte) error {
	return nil
}

// Constructor to initialise the internal State map
func NewCustomMockStub(name string, cc Chaincode, creator []byte) *CustomMockStub {
	s := new(CustomMockStub)
	s.stub = NewMockStub(name, cc)
	s.creator = []byte(creator)
	s.filteredStub = NewMockStub(name, cc)
	return s
}

// SetCreator : To change Creator / Enrollment ID
func (mock *CustomMockStub) SetCreator(creator []byte) {
	mock.creator = []byte(creator)
}

/*****************************
Mock Query Iterator
*****************************/

type MockQueryIterator struct {
	Closed     bool
	Stub       *MockStub
	Slice      *queryParser.StateCouchDBQueryResult
	currentLoc int
}

func (iter *MockQueryIterator) Next() (*queryresult.KV, error) {
	iter.currentLoc++
	result := (*iter.Slice)[iter.currentLoc-1]

	valueAsBytes, err := json.Marshal(result.Value)

	if err != nil {
		return nil, err
	}

	return &queryresult.KV{
		Key:       result.Key,
		Value:     valueAsBytes,
		Namespace: iter.Stub.Name,
	}, nil
}

func (iter *MockQueryIterator) HasNext() bool {
	return iter.currentLoc < len(*iter.Slice)
}

// Close closes the range query iterator. This should be called when done
// reading from the iterator to free up resources.
func (iter *MockQueryIterator) Close() error {
	if iter.Closed == true {
		mockLogger.Error("MockQueryIterator.Close() called after Close()")
		return errors.New("MockQueryIterator.Close() called after Close()")
	}

	iter.Closed = true
	return nil
}

func (iter *MockQueryIterator) Print() {
	mockLogger.Debug("MockQueryIterator {")
	mockLogger.Debug("Closed?", iter.Closed)
	mockLogger.Debug("Stub", iter.Stub)
	mockLogger.Debug("Slice", iter.Slice)
	mockLogger.Debug("HasNext?", iter.HasNext())
	mockLogger.Debug("}")
}

func NewMockQueryIterator(stub *MockStub, results *queryParser.StateCouchDBQueryResult) *MockQueryIterator {
	mockLogger.Debug("NewMockQueryIterator(", stub, ")")
	iter := new(MockQueryIterator)
	iter.Closed = false
	iter.Stub = stub
	iter.Slice = results
	iter.currentLoc = 0

	iter.Print()

	return iter
}
