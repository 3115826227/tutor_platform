package commonDao

import (
	"tutor_platform/src/common"
	. "tutor_platform/src/dao"
	"tutor_platform/src/data/commonData"
)

/*
	城市信息插入
	  输入参数： 城市表
	  输出参数： 布尔状态
*/
func InsertCity(city *commonData.City) bool {
	has, _ := DB.Insert(city)
	return has != 0
}

/*
	城市编号的查询
	  输入参数： 城市名
	  输出参数： 城市code
*/
func QueryCity(cityName string) string {
	city := &commonData.City{
		City: cityName,
	}
	ok, _ := DB.Get(city)
	if ok == false {
		code := common.GetUserId()
		city := &commonData.City{
			CityCode: code,
			City:     cityName,
		}
		if InsertCity(city) == false {
			return ""
		}
		return city.CityCode
	}
	return city.CityCode
}
