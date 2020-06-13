package dao

import (
	"bytes"
	"core-api/entity"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

// 传递变长参数
// 民政部DAO层

var (
	mapMcaOnce     sync.Once
	mapMcaInstance *MapMcaDao
)

type MapMcaDao struct {
	engine *xorm.Engine
}

func NewMapMcaDao(engine *xorm.Engine) *MapMcaDao {
	mapMcaOnce.Do(func() {
		mapMcaInstance = &MapMcaDao{
			engine: engine,
		}
		fmt.Println("MapMcaDao,instance...")
	})
	return mapMcaInstance
}

// 查询一条或多条省级记录（不传任何值，返回所有）
// 存在注入危险 1 'or 'saya'='saya //已解决
func (d *MapMcaDao) GetProvince(where entity.ProvinceMcaEntity) []entity.ProvinceMcaEntity {
	args := make([]interface{}, 0)
	datalist := make([]entity.ProvinceMcaEntity, 0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("SELECT `place_code` , `place_name` , `alternative_name` , `provincial_capital` , `prefectural_level_city` , `autonomous_prefecture` , `municipal_district` , `county_level_city` , `county` , `autonomous_county` , `population` , `area` FROM mca_province2019 WHERE  1 = 1 ")
	if where.PlaceCode != "" {
		sql.WriteString(" and `place_code` = ?")
		args = append(args, where.PlaceCode)
	}
	if where.PlaceName != "" {
		sql.WriteString(" and `place_name` LIKE CONCAT(?, '%')")
		args = append(args, where.PlaceName)
	}
	err = d.engine.SQL(sql.String(), args...).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 查询一条或多条市级记录（不传任何值，返回所有）
func (d *MapMcaDao) GetCity(where entity.CityMcaEntity) []entity.CityMcaEntity {
	args := make([]interface{}, 0)
	datalist := make([]entity.CityMcaEntity, 0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("select `place_code` ,`place_name`,`resident`,`population`,`area`,`phone_code`,`post_code`,`province_code` from mca_city2019 where  1 = 1 ")
	if where.PlaceCode != "" {
		sql.WriteString(" and `place_code` = ?")
		args = append(args, where.PlaceCode)
	}
	if where.PlaceName != "" {
		sql.WriteString(" and `place_name` LIKE CONCAT(?, '%')")
		args = append(args, where.PlaceName)
	}
	if where.ProvinceCode != "" {
		sql.WriteString(" and `province_code` = ?")
		args = append(args, where.ProvinceCode)
	}
	err = d.engine.SQL(sql.String(),args...).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

// 查询一条或多条县区级记录（不传任何值，返回所有）
func (d *MapMcaDao) GetDistinct(where entity.DistinctMcaEntity) []entity.DistinctMcaEntity {
	args := make([]interface{}, 0)
	datalist := make([]entity.DistinctMcaEntity, 0)
	err := error(nil)
	var sql bytes.Buffer
	sql.WriteString("select `place_code` ,`place_name`,`resident`,`population`,`area`,`phone_code`,`post_code`,`city_code` from mca_distinct2019 where  1 = 1 ")
	if where.PlaceCode != "" {
		sql.WriteString(" and `place_code` = ?")
		args = append(args, where.PlaceCode)
	}
	if where.PlaceName != "" {
		sql.WriteString(" and `place_name` LIKE CONCAT(?, '%')")
		args = append(args, where.PlaceName)
	}
	if where.CityCode != "" {
		sql.WriteString(" and `city_code` = ?")
		args = append(args, where.CityCode)
	}
	err = d.engine.SQL(sql.String(),args...).Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

/**
[xorm] [info]  2019/07/05 00:12:35.544274 [SQL] SELECT `place_code` , `place_name` , `alternative_name` , `provincial_capital` , `prefectural_level_city` , `autonomous_prefecture` , `municipal_district` , `county_level_city` , `county` , `autonomous_county` , `population` , `area` FROM mca_province2019 WHERE  1 = 1  and `place_name` LIKE CONCAT(?, '%') []interface {}{[]interface {}{"四川"}}
2019/07/05 00:12:35 sql: converting argument $1 type: unsupported type []interface {}, a slice of interface
[xorm] [info]  2019/07/05 00:15:26.209667 [SQL] SELECT `place_code` , `place_name` , `alternative_name` , `provincial_capital` , `prefectural_level_city` , `autonomous_prefecture` , `municipal_district` , `county_level_city` , `county` , `autonomous_county` , `population` , `area` FROM mca_province2019 WHERE  1 = 1  and `place_name` LIKE CONCAT(?, '%') []interface {}{"四川"}
*/
