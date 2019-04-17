package commonDao

import (
	"tutor_platform/src/common"
	"tutor_platform/src/config/code"
	. "tutor_platform/src/dao"
	"tutor_platform/src/data/commonData"
)

func InsertLogin(login *commonData.Login) bool {
	_, err := DB.Insert(login)
	if err != nil {
		return false
	}
	return true
}

func QueryLogin(login *commonData.Login, Identify string) bool {
	has, _ := DB.Get(login)
	return has && login.Identify == Identify
}

func UpdateLogin(login *commonData.Login) interface{} {
	has, _ := DB.Update(login)
	if has == 0 {
		return common.FailResponse(code.DBUpdateErrorCode)
	}
	return common.SuccessResponse(login)
}

func ExistLogin(login *commonData.Login) bool {
	has, _ := DB.Get(login)
	return has
}

//func DeleteLogin(login *commonData.Login) interface{} {
//
//}
