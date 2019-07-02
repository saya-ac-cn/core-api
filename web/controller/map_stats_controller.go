package controller

import (
	"core-api/entity"
	"core-api/service"
	"core-api/tools"
	"github.com/kataras/iris"
)

// 行政区域查询(统计局)

type MapStatsController struct {
	Ctx iris.Context
	MapService service.MapService
	//当使用单例控制器时，由开发人员负责访问安全,所有客户端共享相同的控制器实例。
	//注意任何控制器的方法,是每个客户端，但结构的字段可以在多个客户端共享（如果是结构）
	//没有任何依赖于的动态struct字段依赖项
	//并且所有字段的值都不为零，在这种情况下我们使用uint64，它不是零（即使我们没有设置它手动易于理解的原因）因为它的值为＆{0}
	//以上所有都声明了一个Singleton，请注意，您不必编写一行代码来执行此操作，Iris足够聪明。
	//见`Get`
	Visits uint64
}

// 查询统计局省级数据
func (c *MapStatsController) GetProvince()  {
	placeCode := c.Ctx.URLParamDefault("PlaceCode","")
	placeName := c.Ctx.URLParamDefault("PlaceName","")
	where := entity.MapStatsEntity{
		PlaceCode:placeCode,
		PlaceName:placeName,
	}
	c.Ctx.StatusCode(iris.StatusOK)
	c.Ctx.JSON(tools.NewResultSuccess(c.MapService.GetStatsProvince(where)))
}

// 查询统计局地级市级数据
func (c *MapStatsController) GetCity()  {
	placeCode := c.Ctx.URLParamDefault("PlaceCode","")
	placeName := c.Ctx.URLParamDefault("PlaceName","")
	superCode := c.Ctx.URLParamDefault("SuperCode","")
	if placeName == "" && placeCode == "" && superCode == "" {
		c.Ctx.JSON(tools.NewResultError(-1,"缺少参数"))
	}else {
		where := entity.MapStatsEntity{
			PlaceCode:placeCode,
			PlaceName:placeName,
			SuperCode:superCode,
		}
		c.Ctx.StatusCode(iris.StatusOK)
		c.Ctx.JSON(tools.NewResultSuccess(c.MapService.GetStatsCity(where)))
	}
}

// 查询统计局县区级数据
func (c *MapStatsController) GetDistinct()  {
	placeCode := c.Ctx.URLParamDefault("PlaceCode","")
	placeName := c.Ctx.URLParamDefault("PlaceName","")
	superCode := c.Ctx.URLParamDefault("SuperCode","")
	if placeName == "" && placeCode == "" && superCode == "" {
		c.Ctx.JSON(tools.NewResultError(-1,"缺少参数"))
	}else {
		where := entity.MapStatsEntity{
			PlaceCode:placeCode,
			PlaceName:placeName,
			SuperCode:superCode,
		}
		c.Ctx.StatusCode(iris.StatusOK)
		c.Ctx.JSON(tools.NewResultSuccess(c.MapService.GetStatsDistinct(where)))
	}
}

// 查询统计局镇级数据
func (c *MapStatsController) GetTown()  {
	placeCode := c.Ctx.URLParamDefault("PlaceCode","")
	placeName := c.Ctx.URLParamDefault("PlaceName","")
	superCode := c.Ctx.URLParamDefault("SuperCode","")
	if placeName == "" && placeCode == "" && superCode == "" {
		c.Ctx.JSON(tools.NewResultError(-1,"缺少参数"))
	}else {
		where := entity.MapStatsEntity{
			PlaceCode:placeCode,
			PlaceName:placeName,
			SuperCode:superCode,
		}
		c.Ctx.StatusCode(iris.StatusOK)
		c.Ctx.JSON(tools.NewResultSuccess(c.MapService.GetStatsTown(where)))
	}
}

// 查询统计局村级数据
func (c *MapStatsController) GetVillage()  {
	placeCode := c.Ctx.URLParamDefault("PlaceCode","")
	placeName := c.Ctx.URLParamDefault("PlaceName","")
	superCode := c.Ctx.URLParamDefault("SuperCode","")
	if placeName == "" && placeCode == "" && superCode == "" {
		c.Ctx.JSON(tools.NewResultError(-1,"缺少参数"))
	}else {
		where := entity.MapStatsEntity{
			PlaceCode:placeCode,
			PlaceName:placeName,
			SuperCode:superCode,
		}
		c.Ctx.StatusCode(iris.StatusOK)
		c.Ctx.JSON(tools.NewResultSuccess(c.MapService.GetStatsVillage(where)))
	}
}