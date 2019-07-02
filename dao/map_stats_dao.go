package dao

import (
	"bytes"
	"core-api/entity"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

// 统计局行政区划DAO
var (
	mapStatsDaoOnce sync.Once
	mapStatsInstance *MapStatsDao
)

type MapStatsDao struct {
	engine *xorm.Engine
}

func NewMapSatsDao(engine *xorm.Engine) *MapStatsDao {
	mapStatsDaoOnce.Do(func() {
		mapStatsInstance = &MapStatsDao{
			engine: engine,
		}
		fmt.Println("MapStatsDao,instance...")
	})
	return mapStatsInstance
}

// 查询一条或多条省级记录（不传任何值，返回所有）
func (d *MapStatsDao) GetProvince(where entity.MapStatsEntity) []entity.MapStatsEntity {
	datalist := make([]entity.MapStatsEntity,0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("SELECT `place_code` , `place_name`  FROM stats_province2019 WHERE  1 = 1 ")
	if where.PlaceCode != "" {
		sql.WriteString(" and `place_code` = '"+where.PlaceCode+"'")
	}
	if where.PlaceName != "" {
		sql.WriteString(" and `place_name` LIKE CONCAT('"+where.PlaceName+"', '%')")
	}
	err = d.engine.SQL(sql.String()).Find(&datalist)
	if err != nil{
		log.Println(err)
		return datalist
	}else{

		return datalist
	}
}

// 查询一条或多条市级记录（不传任何值，返回所有）
func (d *MapStatsDao) GetCity(where entity.MapStatsEntity) []entity.MapStatsEntity {
	datalist := make([]entity.MapStatsEntity,0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("SELECT `place_code` , `place_name`, `full_name`,`province_code` as super_code  FROM stats_city2019 WHERE  1 = 1 ")
	if where.PlaceCode != "" {
		sql.WriteString(" and `place_code` = '"+where.PlaceCode+"'")
	}
	if where.PlaceName != "" {
		sql.WriteString(" and `place_name` LIKE CONCAT('"+where.PlaceName+"', '%')")
	}
	if where.SuperCode != "" {
		sql.WriteString(" and `province_code` = '"+where.SuperCode+"'")
	}
	err = d.engine.SQL(sql.String()).Find(&datalist)
	if err != nil{
		log.Println(err)
		return datalist
	}else{

		return datalist
	}
}

// 查询一条或多条县区级记录（不传任何值，返回所有）
func (d *MapStatsDao) GetDistinct(where entity.MapStatsEntity) []entity.MapStatsEntity {
	datalist := make([]entity.MapStatsEntity,0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("SELECT `place_code` , `place_name`, `full_name`,`city_code` as super_code  FROM stats_distinct2019 WHERE  1 = 1 ")
	if where.PlaceCode != "" {
		sql.WriteString(" and `place_code` = '"+where.PlaceCode+"'")
	}
	if where.PlaceName != "" {
		sql.WriteString(" and `place_name` LIKE CONCAT('"+where.PlaceName+"', '%')")
	}
	if where.SuperCode != "" {
		sql.WriteString(" and `city_code` = '"+where.SuperCode+"'")
	}
	err = d.engine.SQL(sql.String()).Find(&datalist)
	if err != nil{
		log.Println(err)
		return datalist
	}else{

		return datalist
	}
}

// 查询一条或多条镇级记录（不传任何值，返回所有）
func (d *MapStatsDao) GetTown(where entity.MapStatsEntity) []entity.MapStatsEntity {
	datalist := make([]entity.MapStatsEntity,0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("SELECT `place_code` , `place_name`, `full_name`,`distinct_code` as super_code  FROM stats_town2019 WHERE  1 = 1 ")
	if where.PlaceCode != "" {
		sql.WriteString(" and `place_code` = '"+where.PlaceCode+"'")
	}
	if where.PlaceName != "" {
		sql.WriteString(" and `place_name` LIKE CONCAT('"+where.PlaceName+"', '%')")
	}
	if where.SuperCode != "" {
		sql.WriteString(" and `distinct_code` = '"+where.SuperCode+"'")
	}
	err = d.engine.SQL(sql.String()).Find(&datalist)
	if err != nil{
		log.Println(err)
		return datalist
	}else{

		return datalist
	}
}

// 查询一条或多条村级记录（不传任何值，返回所有）
func (d *MapStatsDao) GetVillage(where entity.MapStatsEntity) []entity.MapStatsEntity {
	datalist := make([]entity.MapStatsEntity,0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("SELECT `place_code` , `place_name`, `full_name`,`town_code` as super_code  FROM stats_village2019 WHERE  1 = 1 ")
	if where.PlaceCode != "" {
		sql.WriteString(" and `place_code` = '"+where.PlaceCode+"'")
	}
	if where.PlaceName != "" {
		sql.WriteString(" and `place_name` LIKE CONCAT('"+where.PlaceName+"', '%')")
	}
	if where.SuperCode != "" {
		sql.WriteString(" and `town_code` = '"+where.SuperCode+"'")
	}
	err = d.engine.SQL(sql.String()).Find(&datalist)
	if err != nil{
		log.Println(err)
		return datalist
	}else{

		return datalist
	}
}