package main

import (
	"fmt"
	"proj1/blockchain"
	"os" // Output Streaming to a third party application which is not done by fmt
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters 
		OrdererID: "orderer.hf.Rishi.com",

		// Channel parameters
		ChannelID:     "channelone",
		ChannelConfig: os.Getenv("GOPATH") + "/src/proj1/fixtures/artifacts/channelone.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "hf-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "proj1/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()	
}

