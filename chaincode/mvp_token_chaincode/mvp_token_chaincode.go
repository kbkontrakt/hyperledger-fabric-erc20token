/*
Copyright KB KONTRAKT LLC Corp. 2017 All Rights Reserved.
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

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// MVPTokenChaincode  - represents a MINIMUM VIABLE TOKEN chaincode implementation
type MVPTokenChaincode struct {
}

// Init - конструктор
func (t *MVPTokenChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	result, err := stub.GetState("isInited")
	if result != nil {
		return shim.Error("Try to call Init more than once.")
	}

	_, args := stub.GetFunctionAndParameters()
	/*
			    uint64 initialSupply
		        string tokenName
		        uint8 decimalUnits
				string tokenSymbol

	*/
	/*
		initialSupply, _ := strconv.ParseUint(args[0], 10, 64)
		tokenName := args[1]
		decimalUnits, _ := strconv.ParseUint(args[2], 10, 8)
		tokenSymbol := args[3]
	*/
	result, err = stub.GetCreator()

	stub.PutState("creator", result)
	stub.PutState("initialSupply", []byte(args[0]))
	stub.PutState("tokenName", []byte(args[1]))
	stub.PutState("decimalUnits", []byte(args[2]))
	stub.PutState("tokenSymbol", []byte(args[3]))

	err = stub.PutState("isInited", []byte{1})
	if err != nil {
		return shim.Error("PutState doesn't work")
	}

	return shim.Success(nil)
}

// Invoke - вызовы методов
func (t *MVPTokenChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	return shim.Error("Invalid invoke function name.")
}

func main() {
	err := shim.Start(new(MVPTokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
