package entity

//市级实体类（民政部）
type CityMcaEntity struct {
	PlaceCode	string	`json:"PlaceCode" xorm:"place_code"`	//'行政区划代码',
	PlaceName	string	`json:"PlaceName" xorm:"place_name"`	//'省、自治区、直辖市',
	Resident	string	`json:"Resident" xorm:"resident"`	//'驻地',
	population	string	`json:"Population" xorm:"population"`	//'人口（万人）',
	Area	float32	`json:"Area" xorm:"area"`	//'面积（平方千米）',
	PhoneCode	string	`json:"PhoneCode" xorm:"phone_code"`	//'区号',
	PostCode	string	`json:"postCode" xorm:"post_code"`	//'邮编',
	ProvinceCode	string	`json:"ProvinceCode" xorm:"province_code"`	//'所属省、自治区、直辖市行政区划代码',

}

// 覆写表名
func (CityMcaEntity) TableName() string {
	return "mca_city2019"
}