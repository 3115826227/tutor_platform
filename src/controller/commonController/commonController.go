package commonController

import (
	"github.com/gin-gonic/gin"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	"tutor_platform/src/dao/commonDao"
	"tutor_platform/src/dao/recruitDao"
	"tutor_platform/src/dao/studentDao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/data/recruitData"
	"tutor_platform/src/data/studentData"
	"tutor_platform/src/util/redis"
	"net/http"
	"time"
)

func GenerateUserId() string {
	userId := common.GetUserId()
	if commonDao.ExistLogin(&commonData.Login{UserId: userId}) == false {
		return userId
	}
	return GenerateUserId()
}

func GetInfo(c *gin.Context) interface{} {
	userId := GetId(c)
	var user = &commonData.User{
		UserId: userId,
	}
	return commonDao.GetInfo(user)
}

func Logout(c *gin.Context) interface{} {
	userId := GetId(c)
	redis.DelValue(userId)
	return common.SuccessResponse(nil)
}

func GetId(c *gin.Context) string {
	user, _ := c.Get("token")
	return (user.(*commonData.User)).UserId
}

func Register(c *gin.Context, identify string) {
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

func Personal(c *gin.Context, SuccessHTML string, FailHTML string) {
	userId := c.Query("user_id")
	accountId := c.Query("account_id")
	if redis.GetValue(userId) == accountId && accountId != "" {
		c.HTML(http.StatusOK, SuccessHTML, nil)
	} else {
		c.HTML(http.StatusOK, FailHTML, nil)
	}
}

func UpdateLogin(c *gin.Context, Identify string) {
	userId := c.PostForm("user_id")
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
