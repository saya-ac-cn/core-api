package entity

//省级实体类（民政部）
type ProvinceMcaEntity struct {
	PlaceCode	string	`json:"PlaceCode" xorm:"place_code"`	//'行政区划代码',
	PlaceName	string	`json:"PlaceName" xorm:"place_name"`	//'省、自治区、直辖市',
	AlternativeName	string	`json:"AlternativeName" xorm:"alternative_name"`	//'别称',
	ProvincialCapital	string	`json:"ProvincialCapital" xorm:"provincial_capital"`	//'省会',
	PrefecturalLevelCity	int	`json:"PrefecturalLevelCity" xorm:"prefectural_level_city"`	//'地级市',
	AutonomousPrefecture	int	`json:"AutonomousPrefecture" xorm:"autonomous_prefecture"`	//'自治州',
	MunicipalDistrict	int	`json:"MunicipalDistrict" xorm:"municipal_district"`	//'市辖区',
	CountyLevelCity	int	`json:"CountyLevelCity" xorm:"county_level_city"`	//'县级市',
	County	int	`json:"County" xorm:"county"`	//'县',
	AutonomousCounty	int	`json:"AutonomousCounty" xorm:"autonomous_county"`	//'自治县',
	Population	float32	`json:"Population" xorm:"population"`	//'人口（万人）',
	Area	float32	`json:"Area" xorm:"area"`	//'面积（平方千米）',
}

// 覆写表名
func (ProvinceMcaEntity) TableName() string {
	return "mca_province2019"
}
