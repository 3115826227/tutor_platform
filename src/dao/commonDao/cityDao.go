package commonDao

import (
	"tutor_platform/src/common"
	. "tutor_platform/src/dao"
	"tutor_platform/src/data/commonData"
)

func InsertCity(city *commonData.City) bool {
	has, _ := DB.Insert(city)
	return has != 0
}

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
