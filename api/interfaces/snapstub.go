/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package interfaces

import (
	"errors"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const notImplemented string = "Required functionality was not implemented"

//SnapStub Implementation of the snap stub interface
type SnapStub struct {
	Args [][]byte
}

// GetArgs ...Get the arguments to the stub call as a 2D byte array
func (sc *SnapStub) GetArgs() [][]byte {
	return sc.Args
}

// GetStringArgs the arguments to the stub call as a string array
func (sc *SnapStub) GetStringArgs() []string {
	args := sc.GetArgs()
	strargs := make([]string, 0, len(args))
	for _, barg := range args {
		strargs = append(strargs, string(barg))
	}
	return strargs

}

//GetFunctionAndParameters ...
func (sc *SnapStub) GetFunctionAndParameters() (string, []string) {
	allargs := sc.GetStringArgs()
	function := ""
	params := []string{}
	if len(allargs) >= 1 {
		function = allargs[0]
		params = allargs[1:]
	}
	return function, params
}

//GetTxID ...
func (sc *SnapStub) GetTxID() string {
	return notImplemented
}

// GetState not supported for Snap
func (sc *SnapStub) GetState(key string) ([]byte, error) {
	return nil, errors.New(notImplemented)
}

// PutState not supported for Snap
func (sc *SnapStub) PutState(key string, value []byte) error {
	return errors.New(notImplemented)
}

// DelState not supported for Snap
func (sc *SnapStub) DelState(key string) error {
	return errors.New(notImplemented)
}

// GetStateByRange not supported for Snap
func (sc *SnapStub) GetStateByRange(startKey, endKey string) (shim.StateQueryIteratorInterface, error) {
	return nil, errors.New(notImplemented)
}

//GetStateByPartialCompositeKey not supported for Snap
func (sc *SnapStub) GetStateByPartialCompositeKey(objectType string, keys []string) (shim.StateQueryIteratorInterface, error) {
	return nil, errors.New(notImplemented)

}

//CreateCompositeKey not supported for Snap
func (sc *SnapStub) CreateCompositeKey(objectType string, attributes []string) (string, error) {
	return "", errors.New(notImplemented)
}

//SplitCompositeKey not supported for Snap
func (sc *SnapStub) SplitCompositeKey(compositeKey string) (string, []string, error) {
	return "", nil, errors.New(notImplemented)
}

//GetQueryResult not supported for Snap
func (sc *SnapStub) GetQueryResult(query string) (shim.StateQueryIteratorInterface, error) {
	return nil, errors.New(notImplemented)
}

// GetHistoryForKey not supported for Snap
func (sc *SnapStub) GetHistoryForKey(key string) (shim.StateQueryIteratorInterface, error) {
	return nil, errors.New(notImplemented)
}

// GetCreator not supported for Snap
func (sc *SnapStub) GetCreator() ([]byte, error) {
	return nil, errors.New(notImplemented)
}

// GetTransient not supported for Snap
func (sc *SnapStub) GetTransient() (map[string][]byte, error) {
	return nil, errors.New(notImplemented)
}

// GetBinding not supported for Snap
func (sc *SnapStub) GetBinding() ([]byte, error) {
	return nil, errors.New(notImplemented)
}

// GetArgsSlice not supported for Snap
func (sc *SnapStub) GetArgsSlice() ([]byte, error) {
	return nil, errors.New(notImplemented)
}

// GetTxTimestamp not supported for Snap
func (sc *SnapStub) GetTxTimestamp() (*timestamp.Timestamp, error) {
	return nil, errors.New(notImplemented)
}

// SetEvent saves the event to be sent when a transaction is made part of a block
func (sc *SnapStub) SetEvent(name string, payload []byte) error {
	return errors.New(notImplemented)
}

// InvokeChaincode not supported for Snap.
func (sc *SnapStub) InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response {
	response := pb.Response{Message: notImplemented}
	return response
}

//NewSnapStub ...
func NewSnapStub(args [][]byte) shim.ChaincodeStubInterface {
	ssc := SnapStub{}
	ssc.Args = args
	return &ssc
}