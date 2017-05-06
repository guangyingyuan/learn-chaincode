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

package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

// RoamingSolutionChaincode is simple chaincode implementing logging to Blockchain


type RoamingSolutionChaincode struct {
}


type CDR struct {

		Number string `json:"Number"`
		TimeStamp string `json:"TimeStamp"`
		CallDuration string `json:"CallDuration"`
		CallCost string `json:"CallCost"`
	    DataDuration string `json:"DataDuration"`
	    DataCost string `json:"DataCost"`
		Status string `json:"Status"`
}



type Subscriber struct {

		Number string `json:"Number"`
		TimeStamp string `json:"TimeStamp"`
		CallDuration string `json:"CallDuration"`
	    DataDuration string `json:"DataDuration"`
		Status string `json:"Status"`
}







// Init method will be called during deployment.

func (t *RoamingSolutionChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Println("Init Chaincode...")
	
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}
	
	// Create Subscriber Details table
	err := stub.CreateTable("RoamingDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "Number", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "TimeStamp", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "CallDuration", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "CallCost", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "DataDuration", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "DataCost", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Status", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	
	if err != nil {
		return nil, errors.New("Failed creating RoamingDetails table.")
	}

	fmt.Println("Init Chaincode...done")

	return nil, nil
}

// EntitlementFromHPMN Query function

func (t *RoamingSolutionChaincode) EntitlementFromHPMNQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


    fmt.Println("EntitlementFromHPMN Query Begins...")

    //update the row with new ServiceProvider
	 
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}
	columns = append(columns, col1)
	columns = append(columns, col2)

	row, err := stub.GetRow("RoamingDetails", columns)
	if err != nil {
		fmt.Println("Failed retriving details of %s: %s", string(args[0]), err)
		return nil, fmt.Errorf("Failed retriving details of %s: %s", string(args[0]), err)
	}
	
    if len(row.Columns) != 0{
		
			CallDuration := row.Columns[3].GetString_()
			CallCost := row.Columns[4].GetString_()
			DataDuration := row.Columns[5].GetString_()
			DataCost := row.Columns[6].GetString_()
			Status1 := row.Columns[7].GetString_()
		
		
            CDRobj := CDR{Number: args[0], TimeStamp: args[1], CallDuration: CallDuration, CallCost: CallCost, DataDuration: DataDuration, DataCost: DataCost, Status: Status1}
			res2F, _ := json.Marshal(CDRobj)
		    fmt.Println(string(res2F))
	
			fmt.Println("EntitlementFromHPMN Query ends...")
			return res2F, nil
		
   
     }
	 
	 return nil, fmt.Errorf("Failed retriving details of %s: %s", string(args[0]), err)

}


// EntitlementFromVPMN Query function

func (t *RoamingSolutionChaincode) EntitlementFromVPMNQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


    fmt.Println("EntitlementFromVPMN Query Begins...")

    //update the row with new ServiceProvider
	 
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}
	columns = append(columns, col1)
	columns = append(columns, col2)
	
	 fmt.Println("columns of query",columns);

	row, err := stub.GetRow("RoamingDetails", columns)
	if err != nil {
		fmt.Println("Failed retriving details of %s: %s", string(args[0]), err)
		return nil, fmt.Errorf("Failed retriving details of %s: %s", string(args[0]), err)
	}
	
    if len(row.Columns) != 0{
		
			CallDuration := row.Columns[3].GetString_()
			CallCost := row.Columns[4].GetString_()
			DataDuration := row.Columns[5].GetString_()
			DataCost := row.Columns[6].GetString_()
			Status1 := row.Columns[7].GetString_()
		
		
            CDRobj := CDR{Number: args[0], TimeStamp: args[1], CallDuration: CallDuration, CallCost: CallCost, DataDuration: DataDuration, DataCost: DataCost, Status: Status1}
			res2F, _ := json.Marshal(CDRobj)
		    fmt.Println(string(res2F))
	
			fmt.Println("EntitlementFromVPMN Query ends...")
			return res2F, nil
		
   
     }
	 return nil, fmt.Errorf("Failed retriving details of %s: %s", string(args[0]), err)

}

// EntitlementFromHPMN Invoke function

func (t *RoamingSolutionChaincode) EntitlementFromHPMN(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("EntitlementFromHPMN invoke Begins...")


	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	
	    var err error 
		
		Status1 := "CDRApprovedByHPMN"
		key := args[0]
		
		valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		} else if len(valAsbytes) == 0{
			jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		res := CDR{}
        json.Unmarshal(valAsbytes, &res)
		
		CallDuration := res.CallDuration
		CallCost := res.CallCost
		DataDuration := res.DataDuration
		DataCost := res.DataCost
		
		
		// Delete the previous Instance 
		
		
		
		
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}
		columns = append(columns, col1)
		columns = append(columns, col2)
		
		
		
		
		
		err = stub.DeleteRow("RoamingDetails",columns)
		
		// Insert Data to Internal RocksDB
		
		ok, errNew := stub.InsertRow("RoamingDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
			&shim.Column{Value: &shim.Column_String_{String_: CallDuration}},
			&shim.Column{Value: &shim.Column_String_{String_: CallCost}},
			&shim.Column{Value: &shim.Column_String_{String_: DataDuration}},
			&shim.Column{Value: &shim.Column_String_{String_: DataCost}},
			&shim.Column{Value: &shim.Column_String_{String_: Status1}},
			},
	    })

		if !ok && errNew == nil {
		return nil, errors.New("Insertion Failed")
		}

		
            CDRobj := CDR{Number: args[0], CallDuration: CallDuration, CallCost: CallCost, DataDuration: DataDuration, DataCost: DataCost, Status: Status1}
			res2F, _ := json.Marshal(CDRobj)
		    fmt.Println(string(res2F))
		    err = stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
			
	
		fmt.Println("EntitlementFromHPMN invoke ends...")
		return nil, nil
		
   
}

// EntitlementFromVPMN Invoke function

func (t *RoamingSolutionChaincode) EntitlementFromVPMN(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
      
        if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	    }
		
		 var CallDurationint,CallCostint,DataDurationint,DataCostint int
		
	
		Status1 := "SubscriberDetailsReceived"
		key := args[0]
		
		
		Subscriberobj := Subscriber{Number: args[0], TimeStamp: args[1], CallDuration: args[2], DataDuration: args[3], Status: Status1}
	    res2F, _ := json.Marshal(Subscriberobj)
        fmt.Println(string(res2F))
	    err := stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
	
		
		valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		} else if len(valAsbytes) == 0{
			jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		res := Subscriber{}
        json.Unmarshal(valAsbytes, &res)
        
		
		
	    fmt.Println("Subscriber Details Structure",res)
	    
	   
		
		
	    CallDurationint, err = strconv.Atoi(res.CallDuration)
		if err != nil {
		return nil, err
	    }
		
		DataDurationint, err = strconv.Atoi(res.DataDuration)
		if err != nil {
		return nil, err
	    }
		
			
		// Calculate cost details for VPMN Service
		
		if DataDurationint <= 0 {
		   DataCostint = 0
		}
		
		if CallDurationint <= 0 {
		   CallCostint = 0
		}
		
		if CallDurationint >= 1 && CallDurationint <= 300{
		   CallCostint = CallDurationint * 2;
		}
		
		if CallDurationint > 300 && CallDurationint <= 1500{
		   CallCostint = CallDurationint * 3;
		}

        if CallDurationint > 1500 {
		   CallCostint = CallDurationint * 4;
		}
		
		
		if DataDurationint >= 1 && DataDurationint <= 300{
		   DataCostint = DataDurationint * 3;
		}
		
		if DataDurationint > 300 && DataDurationint <= 1500{
		   DataCostint = DataDurationint * 4;
		}

        if DataDurationint > 1500 {
		   DataCostint = DataDurationint * 5;
		}
		
 
         CallCost := strconv.Itoa(CallCostint)
         DataCost := strconv.Itoa(DataCostint)
         CallDuration := strconv.Itoa(CallDurationint)
         DataDuration := strconv.Itoa(DataDurationint)
		 
		 
		// Put the state of CDR
		 
          Status1 = "CDRApprovalPending"
		  
		  
		
		// Insert Data to Internal RocksDB
		
		ok, errNew := stub.InsertRow("RoamingDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
			&shim.Column{Value: &shim.Column_String_{String_: CallDuration}},
			&shim.Column{Value: &shim.Column_String_{String_: CallCost}},
			&shim.Column{Value: &shim.Column_String_{String_: DataDuration}},
			&shim.Column{Value: &shim.Column_String_{String_: DataCost}},
			&shim.Column{Value: &shim.Column_String_{String_: Status1}},
			},
	    })

		 if !ok && errNew == nil {
		return nil, errors.New("Insertion Failed")
		}
		
		 if !ok {
		return nil, errors.New("Insertion Failed (not OK )")
		}
		
		
		fmt.Println("Error Structure after insertion",errNew)
		 
		 // Update World State
		
            CDRobj := CDR{Number: args[0], TimeStamp: args[1], CallDuration: CallDuration, CallCost: CallCost, DataDuration: DataDuration, DataCost: DataCost, Status: Status1}
			res2F, _ = json.Marshal(CDRobj)
		    fmt.Println(string(res2F))
		    err = stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
		
		
		
		
		fmt.Println("CDR Details Structure",CDRobj)
			
			
		fmt.Println("Invoke EntitlementFromVPMN Chaincode... end") 
		return nil,nil
	
	

}



// Invoke Function

func (t *RoamingSolutionChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke RoamingSolution Chaincode... start") 

	
	// Handle different functions UserAcceptance
	if function == "EntitlementFromVPMN" {
		return t.EntitlementFromVPMN (stub, args)
	} else if function == "EntitlementFromHPMN" {
		return t.EntitlementFromHPMN(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'EntitlementFromHPMN' or 'EntitlementFromVPMN' but found '" + function + "'")
	}
	
	fmt.Println("Invoke RoamingSolution Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *RoamingSolutionChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query RoamingSolution Chaincode... start") 
	
	
	if function == "EntitlementFromVPMNQuery" {
		return t.EntitlementFromVPMNQuery(stub, args)
	} 
	if function == "EntitlementFromHPMNQuery" {
		return t.EntitlementFromHPMNQuery(stub, args)
	} 

	key := args[0]

    
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query RoamingSolution Chaincode... end") 
    return valAsbytes, nil 
  
	
}



func main() {
	err := shim.Start(new(RoamingSolutionChaincode))
	if err != nil {
		fmt.Println("Error starting RoamingSolutionChaincode: %s", err)
	}
}
