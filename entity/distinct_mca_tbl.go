package entity

//县区级实体类（民政部）
type DistinctMcaEntity struct {
	PlaceCode	string	`json:"PlaceCode" xorm:"place_code"`	//'行政区划代码',
	PlaceName	string	`json:"PlaceName" xorm:"place_name"`	//'省、自治区、直辖市',
	Resident	string	`json:"Resident" xorm:"resident"`	//'驻地',
	population	string	`json:"Population" xorm:"population"`	//'人口（万人）',
	Area	float32	`json:"Area" xorm:"area"`	//'面积（平方千米）',
	PhoneCode	string	`json:"PhoneCode" xorm:"phone_code"`	//'区号',
	PostCode	string	`json:"postCode" xorm:"post_code"`	//'邮编',
	CityCode	string	`json:"CityCode" xorm:"city_code"`	//'所属地级市行政区划代码',

}

// 覆写表名
func (DistinctMcaEntity) TableName() string {
	return "mca_distinct2019"
}