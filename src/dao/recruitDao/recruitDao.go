package recruitDao

import (
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	. "tutor_platform/src/dao"
	"tutor_platform/src/dao/commonDao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/data/recruitData"
)

func Login(login *commonData.Login) bool {
	return commonDao.QueryLogin(login, config.RecruitIdentify)
}

func Insert(recruit *recruitData.Recruit) interface{} {
	_, err := DB.Insert(recruit)
	if err != nil {
		return common.FailResponse(code.DBInsertErrorCode)
	}
	return common.SuccessResponse(recruit)
}

func FindRecruit(recruit *recruitData.Recruit) bool {
	has, _ := DB.Get(recruit)
	return has
}

func UpdateRecruitInfo(recruit *recruitData.Recruit) interface{} {
	_, err := DB.Update(recruit)
	if err != nil {
		return common.FailResponse(code.DBInsertErrorCode)
	}
	return common.SuccessResponse(recruit)
}
