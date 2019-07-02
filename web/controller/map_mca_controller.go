package controller

import (
	"core-api/entity"
	"core-api/service"
	"core-api/tools"
	"github.com/kataras/iris"
)

// 行政区域查询

type MapMcaController struct {
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

// 查询民政部省级数据
func (c *MapMcaController) GetProvince()  {
	placeCode := c.Ctx.URLParamDefault("PlaceCode","")
	placeName := c.Ctx.URLParamDefault("PlaceName","")
	where := entity.ProvinceMcaEntity{
		PlaceCode:placeCode,
		PlaceName:placeName,
	}
	c.Ctx.StatusCode(iris.StatusOK)
	c.Ctx.JSON(tools.NewResultSuccess(c.MapService.GetMcaProvince(where)))
}

// 查询民政部地级市级数据
func (c *MapMcaController) GetCity()  {
	placeCode := c.Ctx.URLParamDefault("PlaceCode","")
	placeName := c.Ctx.URLParamDefault("PlaceName","")
	provinceCode := c.Ctx.URLParamDefault("ProvinceCode","")
	if placeName == "" && placeCode == "" && provinceCode == "" {
		c.Ctx.JSON(tools.NewResultError(-1,"缺少参数"))
	}else {
		where := entity.CityMcaEntity{
			PlaceCode:placeCode,
			PlaceName:placeName,
			ProvinceCode:provinceCode,
		}
		c.Ctx.StatusCode(iris.StatusOK)
		c.Ctx.JSON(tools.NewResultSuccess(c.MapService.GetMcaCity(where)))
	}
}

// 查询民政部县区级数据
func (c *MapMcaController) GetDistinct()  {
	placeCode := c.Ctx.URLParamDefault("PlaceCode","")
	placeName := c.Ctx.URLParamDefault("PlaceName","")
	cityCode := c.Ctx.URLParamDefault("CityCode","")
	if placeName == "" && placeCode == "" && cityCode == "" {
		c.Ctx.JSON(tools.NewResultError(-1,"缺少参数"))
	}else {
		where := entity.DistinctMcaEntity{
			PlaceCode:placeCode,
			PlaceName:placeName,
			CityCode:cityCode,
		}
		c.Ctx.StatusCode(iris.StatusOK)
		c.Ctx.JSON(tools.NewResultSuccess(c.MapService.GetMcaDistinct(where)))
	}
}
