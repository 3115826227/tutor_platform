package commonDao

import (
	"fmt"
	"tutor_platform/src/common"
	"tutor_platform/src/config/code"
	. "tutor_platform/src/dao"
	"tutor_platform/src/data/commonData"
)

/*
	登录信息插入
*/
func InsertLogin(login *commonData.Login) bool {
	_, err := DB.Insert(login)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*
	登录信息查询
	  输入参数： 登录表、身份
	  输出参数： 查询结果
*/
func QueryLogin(login *commonData.Login, Identify string) bool {
	has, _ := DB.Get(login)
	fmt.Println(login)
	fmt.Println(has)
	return has && login.Identify == Identify
}

/*
	登录信息更新：
	  输入参数： 登录表
	  输出参数： JSON结果
*/
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
