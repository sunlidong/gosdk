package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
)

type HouseChainCode struct {
	// 房屋溯源链代码
}

// 房源信息
type RentingHouseInfo struct {
	RentingID        string    `json:"renting_id"`         // 统一编码
	RentingHouseInfo HouseInfo `json:"renting_house_info"` // 房屋信息
	RentingAreaInfo  AreaInfo  `json:"renting_area_info"`  //区域信息
	RentingOrderInfo OrderInfo `json:"renting_order_info"` //订单信息
}

type OrderAllInfo struct{
	RentingID        string    `json:"renting_id"`         // 统一编码
	RentingHouseInfo HouseInfo `json:"renting_house_info"` // 房屋信息
	RentingAreaInfo  AreaInfo  `json:"renting_area_info"`  //区域信息
	RentingOrderInfo []OrderInfo `json:"renting_order_info"` //订单信息
}

// 房屋信息  ofgj
type HouseInfo struct {
	HouseID    string `json:"house_id"`    // 房产证编号
	HouseOwner string `json:"house_owner"` // 房主
	RegDate    string `json:"reg_date"`    // 登记日期
	HouseArea  string `json:"house_area"`  // 住房面积
	HouseUsed  string `json:"house_used"`  // 房屋设计用途
	IsMortgage string `json:"is_mortgage"` // 是否抵押
}

// 社区信息 otgj
type AreaInfo struct {
	AreaID       string `json:"area_id"`       // 社区编号
	AreaAddress  string `json:"area_address"`  // 房源所在区域
	BasicNetWork string `json:"basic_net_work"` // 区域基础网络编号
	CPoliceName  string `json:"c_police_name"`  // 社区民警姓名
	CPoliceNum   string `json:"c_police_num"`   // 社区民警工号
}

// 订单信息	oagency
type OrderInfo struct {
	DocHash   string `json:"doc_hash"`    // 电子合同Hash
	OrderID   string `json:"order_id"`    // 订单编号
	RenterID  string `json:"renter_id"` // 承租人信息
	RentMoney string `json:"rent_money"`  // 租金
	BeginDate string `json:"begin_date"`  // 开始日期
	EndDate   string `json:"end_date"`    // 结束日期
	Note      string `json:"note"`       // 备注
}

func (this *HouseChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println(" ==== zufang_suyuan Init ====")
	fmt.Println("000000000000000000")
	return shim.Success(nil)
}

func (this *HouseChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 获取 函数名 和 参数列表
	fn, parameters := stub.GetFunctionAndParameters()

	// 通过函数名匹配对应的链代码函数调用
	if fn == "addHouseInfo" {
		return this.addHouseInfo(stub,parameters)
	}else if fn =="getHouseInfo"{
		return this.getHouseInfo(stub,parameters)
	}else if fn =="addOrderInfo"{
		return this.addOrderInfo(stub,parameters)
	}else if fn == "getOrderInfo"{
		return this.getOrderInfo(stub,parameters)
	}else if fn == "addAreaInfo"{
		return this.addAreaInfo(stub,parameters)
	}else if fn=="getAreaInfo"{
		return this.getAreaInfo(stub,parameters)
	}

	// 没有任何函数被匹配到，返回错误消息
	fmt.Println("==== fn = ",fn)
	return shim.Error("Received unknow function invocation")
}

// 添加房屋信息
func (this *HouseChainCode) addHouseInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// 定义一个房屋信息结构体
	var HouseInfos RentingHouseInfo

	if len(args) !=7{
		return shim.Error("Incorrect number oof arguments.")
	}

	HouseInfos.RentingID = args[0]
	if HouseInfos.RentingID == ""{
		return shim.Error("RentingId can't be empty.")
	}

	HouseInfos.RentingHouseInfo.HouseID = args[1]
	HouseInfos.RentingHouseInfo.HouseOwner = args[2]
	HouseInfos.RentingHouseInfo.RegDate = args[3]
	HouseInfos.RentingHouseInfo.HouseArea = args[4]
	HouseInfos.RentingHouseInfo.HouseUsed = args[5]
	HouseInfos.RentingHouseInfo.IsMortgage = args[6]

	HouseInfosJsonBytes,err := json.Marshal(HouseInfos)
	if err!=nil{
		return shim.Error(err.Error())
	}

	err = stub.PutState(HouseInfos.RentingID,HouseInfosJsonBytes)
	if err !=nil{
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("ok"))
}

// 查看房屋信息
func (this *HouseChainCode) getHouseInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	// 判断参数个数
	if len(args)!=1{
		return shim.Error("Incorrect number of arguments.")
	}
	// 获取唯一标识码
	rentingID := args[0]

	// 获得查询结果迭代器
	resultItreator ,err := stub.GetHistoryForKey(rentingID)
	if err!=nil{
		return shim.Error(err.Error())
	}

	// 释放迭代器资源
	defer resultItreator.Close()

	var rentHouseInfo HouseInfo

	// 使用迭代器遍历查询结果集
	for resultItreator.HasNext(){
		var HouseInfos RentingHouseInfo
		// 取一条结果
		response,err := resultItreator.Next()
		if err !=nil{
			return shim.Error(err.Error())
		}

		err = json.Unmarshal(response.Value,&HouseInfos)
		if err!=nil{
			return shim.Error(err.Error())
		}

		if HouseInfos.RentingHouseInfo.HouseOwner!=""{
			rentHouseInfo = HouseInfos.RentingHouseInfo
			continue
		}
	}

	jsonAsBytes ,err := json.Marshal(rentHouseInfo)
	if err!=nil{
		return shim.Error(err.Error())
	}

	// 返回数据查询结果
	return shim.Success(jsonAsBytes)
}

// 添加社区信息
func (this *HouseChainCode) addAreaInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	var HouseInfos RentingHouseInfo

	if len(args) != 6 {
		return shim.Error("Add estate incorrect number of arguments.")
	}

	HouseInfos.RentingID = args[0]
	if HouseInfos.RentingID == "" {
		return shim.Error("RentingID can not be empty.")
	}

	HouseInfos.RentingAreaInfo.AreaID = args[1]
	HouseInfos.RentingAreaInfo.AreaAddress = args[2]
	HouseInfos.RentingAreaInfo.BasicNetWork = args[3]
	HouseInfos.RentingAreaInfo.CPoliceName = args[4]
	HouseInfos.RentingAreaInfo.CPoliceNum = args[5]

	//if HouseInfos.RentingAreaInfo.AreaID == "" {
	//	return shim.Error("Area's num not be empty")
	//}

	AreaInfosJsonBytes, err := json.Marshal(HouseInfos)

	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(HouseInfos.RentingID, AreaInfosJsonBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// 获取社区信息
func (this *HouseChainCode) getAreaInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	if len(args) !=1{
		return shim.Error("Incorrent number of arguments.")
	}
	rentingID := args[0]

	resultsIterator,err := stub.GetHistoryForKey(rentingID)
	if err!=nil{
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	var rentAreaInfo AreaInfo

	for resultsIterator.HasNext(){
		var HouseInfos RentingHouseInfo
		response,err := resultsIterator.Next()
		if err!=nil{
			return shim.Error(err.Error())
		}

		json.Unmarshal(response.Value,&HouseInfos)

		if HouseInfos.RentingAreaInfo.AreaID!=""{

			rentAreaInfo =HouseInfos.RentingAreaInfo
			fmt.Println("************* rentAreaInfo = ",rentAreaInfo)
			continue
		}
	}
	jsonAsBytes,err := json.Marshal(rentAreaInfo)
	if err !=nil{
		return shim.Error(err.Error())
	}

	return shim.Success(jsonAsBytes)
}

// 添加订单信息
func (this *HouseChainCode) addOrderInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	var HouseInfos RentingHouseInfo

	if len(args)!=8{
		return shim.Error("Incorrect number of arguments.")
	}

	HouseInfos.RentingID = args[0]
	if HouseInfos.RentingID ==""{
		return shim.Error("RentingID can not be empty.")
	}

	HouseInfos.RentingOrderInfo.DocHash = args[1]
	HouseInfos.RentingOrderInfo.OrderID = args[2]
	HouseInfos.RentingOrderInfo.RenterID= args[3]
	HouseInfos.RentingOrderInfo.RentMoney = args[4]
	HouseInfos.RentingOrderInfo.BeginDate = args[5]
	HouseInfos.RentingOrderInfo.EndDate = args[6]
	HouseInfos.RentingOrderInfo.Note = args[7]

	fmt.Println("^^^^^^^^ = ",HouseInfos.RentingOrderInfo)
	HouseInfosJsonBytes,err := json.Marshal(HouseInfos)
	if err!=nil{
		return shim.Error(err.Error())
	}


	err = stub.PutState(HouseInfos.RentingID,HouseInfosJsonBytes)
	if err!=nil{
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("ok"))
}

// 查看订单信息
func (this *HouseChainCode) getOrderInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	var orderInfos []OrderInfo
	if len(args)!=1{
		return shim.Error("Incorrect number of arguments.")
	}

	rentingID := args[0]

	resultsIterator,err:=stub.GetHistoryForKey(rentingID)
	if err!=nil{
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var rentOrderInfo OrderInfo

	for resultsIterator.HasNext(){
		var HouseInfos RentingHouseInfo
		response,err := resultsIterator.Next()
		if err!=nil{
			return shim.Error(err.Error())
		}

		json.Unmarshal(response.Value,&HouseInfos)
		if HouseInfos.RentingOrderInfo.DocHash!=""{
			orderInfos = append(orderInfos,HouseInfos.RentingOrderInfo)
			continue
		}
	}

	fmt.Println("=============",rentOrderInfo)
	fmt.Println("2222222222222:",orderInfos)
	jsonAsBytes,err :=json.Marshal(orderInfos)


	if err!=nil{
		return shim.Error(err.Error())
	}
	return shim.Success(jsonAsBytes)
}

func main() {
	err := shim.Start(new(HouseChainCode))
	if err != nil {
		fmt.Println("Error starting food chaincode :%s", err)
	}
}
