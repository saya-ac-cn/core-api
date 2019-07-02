package router

import (
	"core-api/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"core-api/bootstrap"
	"core-api/web/controller"
	"log"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	mapService := service.NewMapService()

	//使用路由组
	public := b.PartyFunc("/api/public", func(r iris.Party) {
		// 身份验证
		// r.Use(before)
		// 该控制器用的前缀
		mapMcaController := mvc.New(r.Party("/map/mca"))
		mapMcaController.Register(mapService)
		// 使用单列控制器
		mapMcaController.Handle(&controller.MapMcaController{Visits: 0})

		mapStatsController := mvc.New(r.Party("/map/stats"))
		mapStatsController.Register(mapService)
		// 使用单列控制器
		mapStatsController.Handle(&controller.MapStatsController{Visits: 0})

		// 微信相关控制器
		//wxController := mvc.New(r.Party("/wx"))
		//wxController.Handle(&controller.WxController{Visits:0})
	})
	public.Use(before)


}

// 身份验证
func before(ctx iris.Context) {
	key := ctx.URLParam("key")
	if key == ""{
		log.Println("请登录")
		err := iris.Map{
			"code":     403,
			"result":  nil,
			"message": "请登录",
		}
		ctx.JSON(err)
	}else{
		ctx.Next() //继续执行下一个handler
	}
}