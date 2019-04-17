package commonDao

import (
	"tutor_platform/src/common"
	"tutor_platform/src/config/code"
	. "tutor_platform/src/dao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/response"
	"tutor_platform/src/util/redis"
)

func InsertUser(user *commonData.User) bool {
	_, err := DB.Insert(user)
	if err != nil {
		return false
	}
	return true
}

func FindUser(user *commonData.User) bool {
	has, _ := DB.Get(user)
	return has
}

func QueryUser(user *commonData.User) interface{} {
	has, _ := DB.Get(user)
	if has == false {
		return common.FailResponse(code.DBQueryErrorCode)
	}
	return common.SuccessResponse(user)
}

func GetInfo(user *commonData.User) interface{} {
	has, _ := DB.Get(user)
	if has {
		userId := user.UserId
		accountId := redis.GetValue(userId)
		return common.SuccessResponse(response.UserInfoResponse{
			UserId:    userId,
			Name:      user.Name,
			AccountId: accountId,
			Location:  common.GetCity(user.City),
			Identify:  user.Identify,
		})
	}
	return common.FailResponse(code.DBQueryErrorCode)
}

func UpdateUserJudge(user *commonData.User) bool {
	has, _ := DB.Update(user)
	return has != 0
}

func UpdateUser(user *commonData.User) interface{} {
	has, _ := DB.Update(user)
	if has == 0 {
		return common.FailResponse(code.DBUpdateErrorCode)
	}
	return common.SuccessResponse(user)
}
