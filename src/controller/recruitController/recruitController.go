package recruitController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	"tutor_platform/src/controller/commonController"
	"tutor_platform/src/dao/commonDao"
	"tutor_platform/src/dao/messageDao"
	"tutor_platform/src/dao/recruitDao"
	"tutor_platform/src/dao/studentDao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/data/messageData"
	"tutor_platform/src/data/recruitData"
	"tutor_platform/src/data/studentData"
	"tutor_platform/src/response"
	"tutor_platform/src/util/redis"
	"net/http"
	"time"
)

func RecruitPersonal(c *gin.Context) {
	commonController.Personal(c, config.RecruitPersonalHTML, config.IndexHTML)
}

func Info(c *gin.Context) {
	c.JSON(http.StatusOK, commonController.GetInfo(c))
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, commonController.Logout(c))
}

func Register(c *gin.Context) {
	commonController.Register(c, config.RecruitIdentify)
}

func UpdateLogin(c *gin.Context) {
	commonController.UpdateLogin(c, config.RecruitIdentify)
}

func RecruitLookOwnInfo(c *gin.Context) {
	userId := commonController.GetId(c)
	user := commonData.User{
		UserId: userId,
	}
	if commonDao.FindUser(&user) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBQueryErrorCode))
		return
	}
	recruit := recruitData.Recruit{
		UserId: userId,
	}
	if recruitDao.FindRecruit(&recruit) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBQueryErrorCode))
		return
	}
	recruitResponse := response.RecruitInfoResponse{
		UserId:   user.UserId,
		UserName: user.UserName,
		Name:     user.Name,
		Location: common.GetCity(user.City),
		Identify: user.Identify,
		Phone:    user.Phone,
		Gender:   user.Gender,
	}
	c.JSON(http.StatusOK, common.SuccessResponse(recruitResponse))
}

func RecruitUpdateOwnInfo(c *gin.Context) {
	userId := commonController.GetId(c)
	username := c.PostForm("username")
	name := c.PostForm("name")
	gender := c.PostForm("gender")
	city := c.PostForm("city")
	phone := c.PostForm("phone")

	user := commonData.User{
		UserId:     userId,
		UserName:   username,
		Name:       name,
		Gender:     gender,
		City:       city,
		Phone:      phone,
		Identify:   config.RecruitIdentify,
		UpdateTime: time.Now(),
	}
	recruit := recruitData.Recruit{
		UserId:     userId,
		Name:       name,
		Gender:     gender,
		City:       city,
		Phone:      phone,
		Identify:   config.RecruitIdentify,
		UpdateTime: time.Now(),
	}
	if commonDao.UpdateUserJudge(&user) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBUpdateErrorCode))
		return
	}
	c.JSON(http.StatusOK, recruitDao.UpdateRecruitInfo(&recruit))
}

func LookStudentContact(c *gin.Context) {
	userId := commonController.GetId(c)
	studentId := c.Query("student_id")
	mId := c.Query("m_id")
	if JudgeAppointmentExist(userId, studentId, mId) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.AppointmentQueryErrorCode))
		return
	}
	student := studentData.Student{
		UserId: studentId,
	}
	if studentDao.Query(&student) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBQueryErrorCode))
		return
	}
	phoneResponse := response.PhoneInfo{
		UserId: student.UserId,
		Name:   student.Name,
		Phone:  student.Phone,
	}
	c.JSON(http.StatusOK, common.SuccessResponse(phoneResponse))
}

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

func RecruitLookMessageList(c *gin.Context) {
	userId := commonController.GetId(c)
	page := common.StrToInt(c.Query("page"))
	pageSize := common.StrToInt(c.Query("page_size"))
	c.JSON(http.StatusOK, messageDao.QueryMessageList(userId, page, pageSize))
}

func RecruitLookMessageAppointmentList(c *gin.Context) {
	userId := commonController.GetId(c)
	mId := c.Query("m_id")
	message := messageData.Message{
		Mid: mId,
	}
	ok := messageDao.FindMessage(&message)
	if ok == false || userId != message.OwnAccountId {
		c.JSON(http.StatusOK, common.FailResponse(code.AppointmentQueryErrorCode))
		return
	}
	messageKey := fmt.Sprintf("%v:%v", config.MessageAppointmentKey, mId)
	studentList := redis.SMembers(messageKey)
	c.JSON(http.StatusOK, studentDao.BatchQueryStudentInfo(studentList))
}
