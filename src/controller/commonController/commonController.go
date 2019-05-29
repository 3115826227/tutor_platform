package commonController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	"tutor_platform/src/dao/commonDao"
	"tutor_platform/src/dao/messageDao"
	"tutor_platform/src/dao/recruitDao"
	"tutor_platform/src/dao/studentDao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/data/messageData"
	"tutor_platform/src/data/recruitData"
	"tutor_platform/src/data/studentData"
	"tutor_platform/src/util/redis"
)

/*
	用于注册随机id功能
*/
func GenerateUserId() string {
	userId := common.GetUserId()
	//判断是否已经存在该id用户
	if commonDao.ExistLogin(&commonData.Login{UserId: userId}) == false {
		return userId
	}
	return GenerateUserId()
}

/*
	公共获取用户id逻辑
*/
func GetId(c *gin.Context) string {
	user, _ := c.Get("token")
	return (user.(*commonData.User)).UserId
}

/*
	公共获取基本信息逻辑
*/
func CommonGetUserBaseInfo(c *gin.Context) interface{} {
	userId := GetId(c)
	var user = &commonData.User{
		UserId: userId,
	}
	return commonDao.QueryUserBaseInfo(user)
}

/*
	公共退出登录逻辑
*/
func CommonLogout(c *gin.Context) interface{} {
	userId := GetId(c)
	redis.DelValue(userId)
	return common.SuccessResponse(nil)
}

/*
	公共注册逻辑
*/
func CommonRegister(c *gin.Context, identify string) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")
	gender := c.PostForm("gender")
	city := c.PostForm("city")
	phone := c.PostForm("phone")
	userId := GenerateUserId()

	login := &commonData.Login{
		UserId:     userId,
		UserName:   username,
		Password:   password,
		Identify:   identify,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	ok := commonDao.InsertLogin(login)
	if ok == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBInsertErrorCode))
		return
	}
	user := &commonData.User{
		UserId:     userId,
		UserName:   username,
		Name:       name,
		Gender:     gender,
		City:       city,
		Phone:      phone,
		Identify:   identify,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	ok = commonDao.InsertUser(user)
	if ok == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBInsertErrorCode))
		return
	}
	if identify == config.StudentIdentify {
		var studentInfo = &studentData.Student{
			UserId:     user.UserId,
			Name:       name,
			Gender:     gender,
			City:       city,
			Phone:      phone,
			Identify:   identify,
			CreateTime: common.GetTime(),
			UpdateTime: common.GetTime(),
		}
		c.JSON(http.StatusOK, studentDao.Insert(studentInfo))
	} else if identify == config.RecruitIdentify {
		var recruitInfo = &recruitData.Recruit{
			UserId:     user.UserId,
			Name:       name,
			Gender:     gender,
			City:       city,
			Identify:   identify,
			Phone:      phone,
			CreateTime: common.GetTime(),
			UpdateTime: common.GetTime(),
		}
		c.JSON(http.StatusOK, recruitDao.Insert(recruitInfo))
	}
}

/*
	公共进入个人主页逻辑
*/
func CommonEnterHomePage(c *gin.Context, SuccessHTML string, FailHTML string) {
	userId := c.Query("user_id")
	accountId := c.Query("account_id")
	if redis.GetValue(userId) == accountId && accountId != "" {
		c.HTML(http.StatusOK, SuccessHTML, nil)
	} else {
		c.HTML(http.StatusOK, FailHTML, nil)
	}
}

/*
	公共更新登录信息逻辑
*/
func CommonUpdateLogin(c *gin.Context, Identify string) {
	userId := GetId(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	login := &commonData.Login{
		UserId:   userId,
		UserName: username,
		Password: password,
	}
	ok := commonDao.QueryLogin(login, Identify)
	if ok {
		newPassword := c.PostForm("new_password")
		login.Password = newPassword
		c.JSON(http.StatusOK, commonDao.UpdateLogin(login))
	} else {
		c.JSON(http.StatusOK, common.FailResponse(code.PasswordErrorCode))
	}
}

/*
	判断预约是否存在逻辑
*/
func JudgeAppointmentExist(recruitId, studentId, mId string) bool {
	message := messageData.Message{
		Mid: mId,
	}
	ok := messageDao.FindMessage(&message)
	if ok == false || recruitId != message.OwnAccountId {
		return false
	}
	studentKey := fmt.Sprintf("%v:%v", config.StudentAppointmentKey, studentId)
	messageKey := fmt.Sprintf("%v:%v", config.MessageAppointmentKey, mId)
	if redis.SJudgeExist(studentKey, mId) == false || redis.SJudgeExist(messageKey, mId) == false {
		return false
	}
	return true
}
