// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Phongthi structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
type Phongthi struct {
	Ghichu string `json:"ghichu"`
	Hoten string `json:"hoten"`
	Vipham  string `json:"vipham"`
	Diemthi  string `json:"diemthi"`
}

/*
 * The Init method *
 called when the Smart Contract "phongthi-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "phongthi-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "timID" {
		return s.timID(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "taoID" {
		return s.taoID(APIstub, args)
	} else if function == "timTatcaID" {
		return s.timTatcaID(APIstub)
	} else if function == "suaDiemphongthi" {
		return s.suaDiemphongthi(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The timID method *
Used to view the records of one particular phongthi
It takes one argument -- the key for the phongthi in question
 */
func (s *SmartContract) timID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	phongthiAsBytes, _ := APIstub.GetState(args[0])
	if phongthiAsBytes == nil {
		return shim.Error("Could not locate phongthi")
	}
	return shim.Success(phongthiAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 phongthi catches)to our network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	phongthi := []Phongthi{
	    Phongthi{Ghichu: "Không", Vipham: "Không", Hoten: "Nguyễn Văn An", Diemthi: "5"},
		Phongthi{Ghichu: "Không", Vipham: "Không", Hoten: "Nguyễn Thanh Bình", Diemthi: "8"},
		Phongthi{Ghichu: "Không", Vipham: "Không", Hoten: "Lê Văn Cường", Diemthi: "6"},
		Phongthi{Ghichu: "Không", Vipham: "Không", Hoten: "Nguyễn Bích Duyên", Diemthi: "9"},
		Phongthi{Ghichu: "MTL", Vipham: "Mở tài liệu", Hoten: "Trần Thanh Huyền", Diemthi: "4"},
		Phongthi{Ghichu: "Không", Vipham: "Không", Hoten: "Nguyễn Văn Linh", Diemthi: "7"},
		Phongthi{Ghichu: "Không", Vipham: "Không", Hoten: "Lê Văn Nam", Diemthi: "8"},
		Phongthi{Ghichu: "Không", Vipham: "Không", Hoten: "Lê Duy Long", Diemthi: "9"},
		Phongthi{Ghichu: "BT", Vipham: "Bỏ thi", Hoten: "Hoàng Tuấn Kiệt", Diemthi: "0"},
		Phongthi{Ghichu: "Không", Vipham: "Mở tài liệu", Hoten: "Lý Nam Ứng", Diemthi: "5"},
	}


	i := 0
	for i < len(phongthi) {
		fmt.Println("i is ", i)
		phongthiAsBytes, _ := json.Marshal(phongthi[i])
		APIstub.PutState(strconv.Itoa(i+1), phongthiAsBytes)
		fmt.Println("Added", phongthi[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The taoID method *
Fisherman like Sarah would use to record each of her phongthi catches. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) taoID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var phongthi = Phongthi{ Ghichu: args[1], Vipham: args[2], Hoten: args[3], Diemthi: args[4] }

	phongthiAsBytes, _ := json.Marshal(phongthi)
	err := APIstub.PutState(args[0], phongthiAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record phongthi catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The timTatcaID method *
allows for assessing all the records added to the ledger(all phongthi catches)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) timTatcaID(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- timTatcaID:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The suaDiemphongthi method *
The data in the world state can be updated with who has possession. 
This function takes in 2 arguments, phongthi id and new diemthi name. 
 */
func (s *SmartContract) suaDiemphongthi(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	phongthiAsBytes, _ := APIstub.GetState(args[0])
	if phongthiAsBytes == nil {
		return shim.Error("Could not locate phongthi")
	}
	phongthi := Phongthi{}

	json.Unmarshal(phongthiAsBytes, &phongthi)
	// Normally check that the specified argument is a valid diemthi of phongthi
	// we are skipping this check for this example
	phongthi.Diemthi = args[1]

	phongthiAsBytes, _ = json.Marshal(phongthi)
	err := APIstub.PutState(args[0], phongthiAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change phongthi diemthi: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}