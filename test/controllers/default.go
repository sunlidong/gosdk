package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"ht/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

// 发布房源信息
func (this *MainController) FormHouse() {
	this.TplName = "form_house.html"

	rentingID := this.GetString("rentingID")
	if rentingID == "" {
		return
	} else {
		fczbh := this.GetString("fczbh")
		fzxm := this.GetString("fzxm")
		djrq := this.GetString("djrq")
		zfmj := this.GetString("zfmj")
		fwsjyt := this.GetString("fwsjyt")
		sfdy := this.GetString("sfdy")

		var args []string
		args = append(args, "addHouseInfo")
		args = append(args, rentingID)
		args = append(args, fczbh)
		args = append(args, fzxm)
		args = append(args, djrq)
		args = append(args, zfmj)
		args = append(args, fwsjyt)
		args = append(args, sfdy)

		ret, err := models.App.AddHouseItem(args)
		if err != nil {
			fmt.Println("AddHouseItem err...")
		}

		fmt.Println("<--- 添加房源信息结果　--->：", ret)
	}
	this.TplName = "index.html"
}

// 查询房源信息
func (this *MainController) HouseSearch() {
	this.TplName = "house_search.html"

	//　１、获取用户输入的  rentingID
	key := this.GetString("rentingId")

	//　２、组织　chiancode 所需要的参数
	var args []string
	args = append(args, "getHouseInfo")
	args = append(args, key)

	// 3、调用　model 层函数，查询数据
	response, err := models.App.GetHouseInfo(args)
	if err != nil {
		fmt.Println("models.App.GetHouseInfo err....")
	}

	//解析结果集
	var jsonData models.HouseInfo
	err = json.Unmarshal([]byte(response), &jsonData)
	if err != nil {
		fmt.Println("json.Unmarshal err....")
	}

	fmt.Println("----------- jsonData", jsonData)

	// 5、将数据展示在前端界面
	this.Data["houseId"] = jsonData.HouseID
	this.Data["houseOwner"] = jsonData.HouseOwner
	this.Data["regDate"] = jsonData.RegDate
	this.Data["houseArea"] = jsonData.HouseArea
	this.Data["houseUsed"] = jsonData.HouseUsed
	this.Data["isMortgage"] = jsonData.IsMortgage
}

// 发布社区信息
func (this *MainController) FormArea() {
	this.TplName = "form_area.html"

	rentingID := this.GetString("rentingID")
	if rentingID == "" {
		return
	} else {

		sqbh := this.GetString("sqbh")
		sqdz := this.GetString("sqdz")
		sqwlbh := this.GetString("sqwlbh")
		sqmjxm := this.GetString("sqmjxm")
		sqmjgh := this.GetString("sqmjgh")

		var args []string
		args = append(args, "addAreaInfo")
		args = append(args, rentingID)
		args = append(args, sqbh)
		args = append(args, sqdz)
		args = append(args, sqwlbh)
		args = append(args, sqmjxm)
		args = append(args, sqmjgh)

		ret, err := models.App.AddAreaItem(args)
		if err != nil {
			fmt.Println("AddEstateItem err....")
		}

		fmt.Println("<--- 查询房源信息结果　--->：", ret)
	}
}

// 查询社区信息
func (this *MainController) AreaSearch() {
	this.TplName = "area_search.html"

	// 1、获取用户输入　rentingID
	key := this.GetString("rentingId")

	// 2、组织　chiancode 所需要的参数
	var args []string
	args = append(args, "getAreaInfo")
	args = append(args, key)

	// 3、调用　model 层函数，查询数据
	response, err := models.App.GetAreaInfo(args)
	if err != nil {
		fmt.Println("models.App.GetAreaINfo err....")
	}

	fmt.Println("=========== response:", response)

	// 解析结果集
	var jsonData models.AreaInfo
	err = json.Unmarshal([]byte(response), &jsonData)
	if err != nil {
		fmt.Println("json.Unmarshal err....")
	}

	// 5、将数据展示在前端界面
	this.Data["areaId"] = jsonData.AreaID
	this.Data["areaAddr"] = jsonData.AreaAddress
	this.Data["comm_net"] = jsonData.BasicNetWork
	this.Data["areaPoliceName"] = jsonData.CPoliceName
	this.Data["areaPoliceNum"] = jsonData.CPoliceNum
}

// 发布订单信息
func (this *MainController) FormOrderer() {
	this.TplName = "form_orderer.html"
	rentingID := this.GetString("rentingID")
	if rentingID == "" {
		return
	} else {
		orderHash := this.GetString("orderHash")
		orderId := this.GetString("orderId")
		renterId := this.GetString("renterId")
		rentMoney := this.GetString("rentMoney")
		beginDate := this.GetString("beginDate")
		endDate := this.GetString("endDate")
		note := this.GetString("note")

		var args []string
		args = append(args, "addOrderInfo")
		args = append(args, rentingID)
		args = append(args, orderHash)
		args = append(args, orderId)
		args = append(args, renterId)
		args = append(args, rentMoney)
		args = append(args, beginDate)
		args = append(args, endDate)
		args = append(args, note)

		ret, err := models.App.AddOrderItem(args)
		if err != nil {
			fmt.Println("AddOrderItem err...")
		}

		fmt.Println("============= ret =", ret)

		fmt.Println("<--- 添加订单信息结果　--->：", ret)
	}

	this.TplName = "index.html"
}

//　查询订单信息
func (this *MainController) OrdererSearch() {
	this.TplName = "order_search.html"

	// 1、获取用户输入　rentingID
	key := this.GetString("rentingId")

	// 2、组织　chiancode 所需要的参数
	var args []string
	args = append(args, "getOrderInfo")
	args = append(args, key)

	// 3、调用　model 层函数，查询数据
	response, err := models.App.GetOrderInfo(args)
	if err != nil {
		fmt.Println("models.App.GetOrderInfo err....")
	}

	// 解析结果集
	var jsonData []models.OrderInfo
	err = json.Unmarshal([]byte(response), &jsonData)
	if err != nil {
		fmt.Println("json.Unmarshal err....")
	}
	fmt.Println("++++ jsonData:", jsonData)

	for i := 0; i < len(jsonData); i++ {
		fmt.Println("DocHash = ", jsonData[i].DocHash)
		fmt.Println("OrderID = ", jsonData[i].OrderID)
		fmt.Println("RenterID = ", jsonData[i].RenterID)
		fmt.Println("RentMoney = ", jsonData[i].RentMoney)
		fmt.Println("BeginDate = ", jsonData[i].BeginDate)
		fmt.Println("EndDate = ", jsonData[i].EndDate)
		fmt.Println("Note = ", jsonData[i].Note)

		// 5、将数据展示在前端界面
		var docHash string = fmt.Sprintf("%s%d", "docHash", i)
		var orderId string = fmt.Sprintf("%s%d", "orderId", i)
		var renterID string = fmt.Sprintf("%s%d", "renterID", i)
		var rentMoney string = fmt.Sprintf("%s%d", "rentMoney", i)
		var beginDate string = fmt.Sprintf("%s%d", "beginDate", i)
		var endData string = fmt.Sprintf("%s%d", "endData", i)
		var note string = fmt.Sprintf("%s%d", "note", i)

		/*
			fmt.Println(docHash)
			fmt.Println(orderId)
			fmt.Println(renterID)
			fmt.Println(rentMoney)
			fmt.Println(beginDate)
			fmt.Println(endData)
			fmt.Println(note)
		*/

		this.Data[docHash] = jsonData[i].DocHash
		this.Data[orderId] = jsonData[i].OrderID
		this.Data[renterID] = jsonData[i].RenterID
		this.Data[rentMoney] = jsonData[i].RentMoney
		this.Data[beginDate] = jsonData[i].BeginDate
		this.Data[endData] = jsonData[i].EndDate
		this.Data[note] = jsonData[i].Note
	}
}
