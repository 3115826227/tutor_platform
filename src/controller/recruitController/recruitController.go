package recruitController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
)

func RecruitPersonal(c *gin.Context) {
	commonController.CommonEnterHomePage(c, config.RecruitPersonalHTML, config.IndexHTML)
}

func BaseInfo(c *gin.Context) {
	c.JSON(http.StatusOK, commonController.CommonGetUserBaseInfo(c))
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, commonController.CommonLogout(c))
}

func Register(c *gin.Context) {
	commonController.CommonRegister(c, config.RecruitIdentify)
}

func UpdateLogin(c *gin.Context) {
	commonController.CommonUpdateLogin(c, config.RecruitIdentify)
}

/*
	招聘用户查看个人信息
*/
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

/*
	招聘用户更新个人信息
*/
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

/*
	招聘用户查看自己发布的招聘信息列表
*/
func RecruitLookMessageList(c *gin.Context) {
	userId := commonController.GetId(c)
	c.JSON(http.StatusOK, messageDao.QueryMessageList(userId))
}

/*
	招聘用户查看具体招聘信息被预约列表
*/
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

/*
	招聘用户查看被预约的学生用户联系方式
*/
func RecruitLookStudentContact(c *gin.Context) {
	userId := commonController.GetId(c)
	studentId := c.Query("student_id")
	mId := c.Query("m_id")
	if commonController.JudgeAppointmentExist(userId, studentId, mId) == false {
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
