package messageController

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	"tutor_platform/src/controller/commonController"
	"tutor_platform/src/dao/commonDao"
	"tutor_platform/src/dao/messageDao"
	"tutor_platform/src/data/messageData"
	"tutor_platform/src/response"
	"tutor_platform/src/util/redis"
	"net/http"
	"strconv"
	"time"
)

func LookStudentInfo(c *gin.Context) {

}

func LookRecruitInfo(c *gin.Context) {

}

func AddTutor(c *gin.Context) {
	userId := commonController.GetId(c)
	studentLevel := c.PostForm("student_level")
	studentGender := c.PostForm("student_gender")
	city := c.PostForm("city")
	local := c.PostForm("local")
	salary := c.PostForm("salary")
	describe := c.PostForm("describe")
	mid := common.GetUserId()

	tutor := messageData.Tutor{
		Mid:           mid,
		OwnAccountId:  userId,
		StudentLevel:  studentLevel,
		StudentGender: studentGender,
		TypeCode:      1,
		City:          city,
		Local:         local,
		Salary:        salary,
		Describe:      describe,
	}
	message := messageData.Message{
		Mid:          mid,
		OwnAccountId: userId,
		TypeCode:     1,
		City:         city,
		Status:       "发布中",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	if messageDao.InsertMessage(&message) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBInsertErrorCode))
		return
	}
	c.JSON(http.StatusOK, messageDao.InsertTutor(&tutor))
}

func UpdateTutor(c *gin.Context) {
	userId := c.PostForm("user_id")
	studentLevel := c.PostForm("student_level")
	studentGender := c.PostForm("student_gender")
	city := c.PostForm("city")
	local := c.PostForm("local")
	salary := c.PostForm("salary")
	describe := c.PostForm("describe")
	mid := c.PostForm("mid")

	oldTutor := messageData.Tutor{
		Mid: mid,
	}
	if messageDao.FindTutor(&oldTutor) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBQueryErrorCode))
		return
	}
	message := messageData.Message{
		Mid:        mid,
		UpdateTime: time.Now(),
	}
	if messageDao.UpdateMessage(&message) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBUpdateErrorCode))
		return
	}
	tutor := messageData.Tutor{
		Mid:           mid,
		OwnAccountId:  userId,
		StudentLevel:  studentLevel,
		StudentGender: studentGender,
		City:          city,
		Local:         local,
		Salary:        salary,
		Describe:      describe,
	}
	c.JSON(http.StatusOK, messageDao.UpdateTutor(&tutor))
}

func DeleteMessage(mid string) bool {
	message := messageData.Message{
		Mid: mid,
	}
	return messageDao.DeleteMessage(&message)
}

func DeleteTutor(c *gin.Context) {
	mid := c.Query("mid")
	tutor := messageData.Tutor{
		Mid: mid,
	}
	if DeleteMessage(mid) && messageDao.DeleteTutor(&tutor) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBDeleteErrorCode))
	} else {
		c.JSON(http.StatusOK, common.SuccessResponse(nil))
	}
}

func QueryTutor(c *gin.Context) {
	mid := c.Query("mid")
	tutor := messageData.Tutor{
		Mid: mid,
	}
	if messageDao.FindTutor(&tutor) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBQueryErrorCode))
		return
	}
	key := fmt.Sprintf("%v:%v:%v", config.CityCodeKey, commonDao.QueryCity(tutor.City), config.ViewListKey)
	redis.ZDel(key, mid)
	views := common.StrToInt(redis.GetValue(fmt.Sprintf("%v:%v", config.ViewKey, mid))) + 1
	redis.UpdateValue(fmt.Sprintf("%v:%v", config.ViewKey, mid), strconv.Itoa(views))
	redis.ZAdd(key, mid, views)
	message := messageData.Message{
		Mid: mid,
	}
	messageDao.FindMessage(&message)
	tutorInfo := response.TutorInfo{
		Mid:               mid,
		Name:              tutor.Name,
		City:              tutor.City,
		Local:             tutor.Local,
		OwnAccountId:      tutor.OwnAccountId,
		StudentGender:     tutor.StudentGender,
		StudentLevel:      tutor.StudentLevel,
		Salary:            tutor.Salary,
		Describe:          tutor.Describe,
		Views:             views,
		AppointmentNumber: message.AppointmentNumber,
		CreateTime:        common.TimeToTime(message.CreateTime),
	}
	tutorResponse := response.TutorResponse{
		List: []response.TutorInfo{tutorInfo},
	}
	c.JSON(http.StatusOK, common.SuccessResponse(tutorResponse))
}

func QueryWeekTutor(c *gin.Context) {
	c.JSON(http.StatusOK, messageDao.QueryWeekTutor())
}

func QueryTop10Tutor(c *gin.Context) {
	city := c.Query("city")
	fmt.Println(city)
	c.JSON(http.StatusOK, messageDao.QueryTop10Tutor(city))
}

func QueryPageTutor(c *gin.Context) {
	city, _ := com.UrlDecode(c.Query("city"))
	page := common.StrToInt(c.Query("page"))
	pageSize := common.StrToInt(c.Query("page_size"))
	c.JSON(http.StatusOK, messageDao.QueryPageTutor(city, page, pageSize))
}

func JudgeAppointment(c *gin.Context) {
	userId := commonController.GetId(c)
	mId := c.Query("m_id")
	studentKey := fmt.Sprintf("%v:%v", config.StudentAppointmentKey, userId)
	if redis.SJudgeExist(studentKey, mId) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.RedisQueryErrorCode))
	} else {
		c.JSON(http.StatusOK, common.SuccessResponse(nil))
	}
}

func AppointmentTutor(c *gin.Context) {
	userId := commonController.GetId(c)
	mId := c.Query("m_id")
	studentKey := fmt.Sprintf("%v:%v", config.StudentAppointmentKey, userId)
	redis.SAdd(studentKey, mId)
	messageKey := fmt.Sprintf("%v:%v", config.MessageAppointmentKey, mId)
	redis.SAdd(messageKey, userId)
	message := messageData.Message{
		Mid: mId,
	}
	messageDao.FindMessage(&message)
	message.AppointmentNumber = message.AppointmentNumber + 1
	messageDao.UpdateMessage(&message)
	c.JSON(http.StatusOK, common.SuccessResponse("预约成功"))
}
