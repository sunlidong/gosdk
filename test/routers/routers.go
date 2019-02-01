package routers

import (
	"sdktest/test/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/form_area",&controllers.MainController{},"get:FormArea")
	beego.Router("/form_house",&controllers.MainController{},"get:FormHouse")
	beego.Router("/form_orderer",&controllers.MainController{},"get:FormOrderer")
	beego.Router("/area_search",&controllers.MainController{},"get:AreaSearch")
	beego.Router("/house_search",&controllers.MainController{},"get:HouseSearch" )
	beego.Router("/orderer_search",&controllers.MainController{},"get:OrdererSearch")
}
