package models

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (this *Application) GetOrderInfo(args []string) (string, error) {
	response,err :=this.FabricSetup.client.Query(channel.Request{ChaincodeID:this.FabricSetup.ChainCodeID,Fcn:args[0],Args:[][]byte{[]byte(args[1])}})
	if err!=nil{
		return "",fmt.Errorf("failed to query: %v",err)
	}

	return string(response.Payload),nil
}
