package router

import (
	"github.com/astaxie/beego"
	"github.com/duwu/SpikeSystem/SecKill/SecAdmin/controller/activity"
	"github.com/duwu/SpikeSystem/SecKill/SecAdmin/controller/product"
)

func init() {
	beego.Router("/product/list", &product.ProductController{}, "*:ListProduct")
	beego.Router("/", &product.ProductController{}, "*:ListProduct")
	beego.Router("/product/create", &product.ProductController{}, "*:CreateProduct")
	beego.Router("/product/submit", &product.ProductController{}, "*:SubmitProduct")

	beego.Router("/activity/create", &activity.ActivityController{}, "*:CreateActivity")
	beego.Router("/activity/list", &activity.ActivityController{}, "*:ListActivity")
	beego.Router("/activity/submit", &activity.ActivityController{}, "*:SubmitActivity")
}
