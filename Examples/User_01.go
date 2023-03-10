
package main

import (
    "fmt"
    "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
    "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
    "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
    "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)


func main() {
    
    // Create SDK setup for the integration tests
    sdk, err := fabsdk.New(config.FromFile("./connection.yaml"))
    if err != nil {
        fmt.Println("Failed to create new SDK: ", err)
        return
    }
    
    // Prepare client context
    orgName := "org1"
    orgAdmin := "Admin"
    orgMsp := "Org1MSP"

    clientChannelContext := sdk.ChannelContext("mychannel", channel.WithUser(orgAdmin), channel.WithOrg(orgName))
    if clientChannelContext == nil {
    fmt.Println("Failed to create client channel context")
    return
    }
    
    client, err := channel.New(clientChannelContext)
     if err != nil {
        fmt.Println("Failed to create new client: ", err)
        return
    }
    
      // Get user identity
    userOrg1, err := sdk.Context(fabsdk.WithUser("user1"), fabsdk.WithOrg(orgName))
    if err != nil {
        fmt.Println("Failed to get user context: ", err)
        return
    }
    
      identityOrg1, err := userOrg1.Identity()
    if err != nil {
        fmt.Println("Failed to get user identity: ", err)
        return
    }
    
    
     // Prepare request
    request := channel.Request{
        ChaincodeID: "mychaincode",
        Fcn:         "getBalance",
        Args: [][]byte{
            []byte("user1"),
        },
        TxID: "",
        IsInit: false,
    }
    
      // Set user context
    client.SetUserContext(identityOrg1.(msp.SigningIdentity), orgMsp) 
    
    // Send request and get response
    response, err := client.Query(request)
    if err != nil {
        fmt.Println("Failed to query chaincode: ", err)
        return
    }

    fmt.Printf("User1 balance is %s\n", response.Payload)
    
}

/*
This code is  used to query the account balance of a user named "user1" from the smart contract named "mychaincode".
The code first performs the steps to get the SDK configuration and user ID. 
It then connects to the network from a YAML file with connection information and creates a channel to make the request. 
Finally it prepares a transaction to make the request, sets the user ID, executes the transaction, and prints the results to the console.



*/
