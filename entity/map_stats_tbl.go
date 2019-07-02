package entity

// 国家统计局5级区域实体模型

type MapStatsEntity struct {
		PlaceCode string `json:"PlaceCode" xorm:"place_code"` //'行政区划代码',
		PlaceName string `json:"PlaceName" xorm:"place_name"` //'省、自治区、直辖市',
		FullName string `json:"FullName" xorm:"full_name"` //'全名',
		SuperCode string `json:"SuperCode" xorm:"super_code"` //'上级编码',
}
