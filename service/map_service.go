package service

import (
	"core-api/dao"
	"core-api/datasource"
	"core-api/entity"
	"fmt"
	"sync"
)

// 民政部行政区域数据业务层查询

var(
	mapServiceOnce sync.Once
	mapServiceInstance MapService
)

type MapService interface {
	//查询一条或多条民政部省级记录（不传任何值，返回所有）
	GetMcaProvince(where entity.ProvinceMcaEntity) []entity.ProvinceMcaEntity
	//查询一条或多条民政部地级市记录（不传任何值，返回所有）
	GetMcaCity(where entity.CityMcaEntity) []entity.CityMcaEntity
	//查询一条或多条民政部县区记录（不传任何值，返回所有）
	GetMcaDistinct(where entity.DistinctMcaEntity) []entity.DistinctMcaEntity
	//查询一条或多条统计局省级记录（不传任何值，返回所有）
	GetStatsProvince(where entity.MapStatsEntity) []entity.MapStatsEntity
	//查询一条或多条统计局市级记录（不传任何值，返回所有）
	GetStatsCity(where entity.MapStatsEntity) []entity.MapStatsEntity
	//查询一条或多条统计局县区级记录（不传任何值，返回所有）
	GetStatsDistinct(where entity.MapStatsEntity) []entity.MapStatsEntity
	//查询一条或多条统计局镇级记录（不传任何值，返回所有
	GetStatsTown(where entity.MapStatsEntity) []entity.MapStatsEntity
	//查询一条或多条统计局村级记录（不传任何值，返回所有）
	GetStatsVillage(where entity.MapStatsEntity) []entity.MapStatsEntity
}

type mapService struct {
	mcaDao *dao.MapMcaDao
	statsDao *dao.MapStatsDao
}

func NewMapService() MapService {
	mapServiceOnce.Do(func() {
		mapServiceInstance = &mapService{
			mcaDao:dao.NewMapMcaDao (datasource.InstanceMaster()),
			statsDao:dao.NewMapSatsDao(datasource.InstanceMaster()),
		}
		fmt.Println("ProvinceService,instance...")
	})
	return mapServiceInstance
}

//查询一条或多条民政部省级记录（不传任何值，返回所有）
func (s *mapService) GetMcaProvince(where entity.ProvinceMcaEntity) []entity.ProvinceMcaEntity {
	return s.mcaDao.GetProvince(where)
}

//查询一条或多条民政部地级市记录（不传任何值，返回所有）
func (s *mapService) GetMcaCity(where entity.CityMcaEntity) []entity.CityMcaEntity {
	return s.mcaDao.GetCity(where)
}

//查询一条或多条民政部县区记录（不传任何值，返回所有）
func (s *mapService) GetMcaDistinct(where entity.DistinctMcaEntity) []entity.DistinctMcaEntity {
	return s.mcaDao.GetDistinct(where)
}

//查询一条或多条统计局省级记录（不传任何值，返回所有）
func (s * mapService) GetStatsProvince(where entity.MapStatsEntity) []entity.MapStatsEntity{
	return s.statsDao.GetProvince(where)
}

//查询一条或多条统计局市级记录（不传任何值，返回所有）
func (s * mapService) GetStatsCity(where entity.MapStatsEntity) []entity.MapStatsEntity{
	return s.statsDao.GetCity(where)
}

//查询一条或多条统计局县区级记录（不传任何值，返回所有）
func (s * mapService) GetStatsDistinct(where entity.MapStatsEntity) []entity.MapStatsEntity{
	return s.statsDao.GetDistinct(where)
}

//查询一条或多条统计局镇级记录（不传任何值，返回所有）
func (s * mapService) GetStatsTown(where entity.MapStatsEntity) []entity.MapStatsEntity{
	return s.statsDao.GetTown(where)
}

//查询一条或多条统计局村级记录（不传任何值，返回所有）
func (s * mapService) GetStatsVillage(where entity.MapStatsEntity) []entity.MapStatsEntity{
	return s.statsDao.GetVillage(where)
}