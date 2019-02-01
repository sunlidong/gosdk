package models

import (
	"fmt"
	"os"
)

var App Application

func init() {
	fSetup := FabricSetup{
		// Network parameters
		OrdererID: "orderer.itcast.cn",

		// Channel parameters
		ChannelID:     "fgjchannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/origins/conf/channel-artifacts/fgjchannel.tx",

		// Chaincode parameters
		ChainCodeID:     "mycc",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "origins/chaincode",
		OrgAdmin:        "Admin",
		OrgName:         "ofgj",
		ConfigFile:      "conf/config.yaml",

		// User parameters
		UserName: "User1",
		eventID:"eventInvoke",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}

	// Close SDK
	//defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	// Launch the web application listening
	App = Application{
		FabricSetup: &fSetup,
	}
}
