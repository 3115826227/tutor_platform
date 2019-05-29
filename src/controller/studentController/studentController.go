package studentController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	"tutor_platform/src/controller/commonController"
	"tutor_platform/src/dao/commonDao"
	"tutor_platform/src/dao/messageDao"
	"tutor_platform/src/dao/studentDao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/data/messageData"
	"tutor_platform/src/data/studentData"
	"tutor_platform/src/response"
	"tutor_platform/src/util/redis"
)

func StudentPersonal(c *gin.Context) {
	commonController.CommonEnterHomePage(c, config.StudentPersonalHTML, config.IndexHTML)
}

func BaseInfo(c *gin.Context) {
	c.JSON(http.StatusOK, commonController.CommonGetUserBaseInfo(c))
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, commonController.CommonLogout(c))
}

func Register(c *gin.Context) {
	commonController.CommonRegister(c, config.StudentIdentify)
}

func UpdateLogin(c *gin.Context) {
	commonController.CommonUpdateLogin(c, config.StudentIdentify)
}

/*
	学生用户查看个人信息
*/
func StudentLookOwnInfo(c *gin.Context) {
	userId := commonController.GetId(c)
	user := commonData.User{
		UserId: userId,
	}
	if commonDao.FindUser(&user) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBQueryErrorCode))
		return
	}
	student := studentData.Student{
		UserId: userId,
	}
	if studentDao.Query(&student) == false {
		c.JSON(http.StatusOK, common.FailResponse(code.DBQueryErrorCode))
		return
	}
	studentResponse := response.StudentInfoResponse{
		UserId:   user.UserId,
		UserName: user.UserName,
		Name:     user.Name,
		Location: common.GetCity(user.City),
		Identify: user.Identify,
		Phone:    user.Phone,
		Gender:   user.Gender,
	}
	c.JSON(http.StatusOK, common.SuccessResponse(studentResponse))
}

/*
	学生用户完善/更新个人信息
*/
func StudentUpdateOwnInfo(c *gin.Context) {
	userId := c.PostForm("user_id")
	school := c.PostForm("school")
	certificate := c.PostForm("certificate")
	describe := c.PostForm("describe")
	educationLevel := c.PostForm("education_level")
	jobStatus := c.PostForm("job_status")
	jobCategory := c.PostForm("job_category")
	jobDesirableSalary := c.PostForm("job_desirable_salary")
	advantagesSubject := c.PostForm("advantages_subject")

	student := studentData.Student{
		UserId:             userId,
		School:             school,
		Certificate:        certificate,
		Describe:           describe,
		EducationLevel:     educationLevel,
		JobStatus:          jobStatus,
		JobCategory:        jobCategory,
		JobDesirableSalary: common.StrToInt64(jobDesirableSalary),
		AdvantagesSubject:  advantagesSubject,
	}
	c.JSON(http.StatusOK, studentDao.UpdateStudentInfo(&student))
}

/*
	学生用户查看预约过的招聘信息列表
*/
func StudentLookAppointmentTutorList(c *gin.Context) {
	userId := commonController.GetId(c)
	studentKey := fmt.Sprintf("%v:%v", config.StudentAppointmentKey, userId)
	tutorList := redis.SMembers(studentKey)
	tutorData := make([]response.TutorInfo, 0)
	for _, tutorId := range tutorList {
		tutor := messageData.Tutor{
			Mid: tutorId,
		}
		messageDao.FindTutor(&tutor)
		message := messageData.Message{
			Mid: tutorId,
		}
		messageDao.FindMessage(&message)
		views := common.StrToInt(redis.GetValue(fmt.Sprintf("%v:%v", config.ViewKey, tutor.Mid)))
		tutorData = append(tutorData, response.TutorInfo{
			Mid:               tutor.Mid,
			Name:              tutor.Name,
			City:              tutor.City,
			Local:             tutor.Local,
			StudentGender:     tutor.StudentGender,
			StudentLevel:      tutor.StudentLevel,
			Salary:            tutor.Salary,
			Describe:          tutor.Describe,
			Views:             views,
			AppointmentNumber: message.AppointmentNumber,
			CreateTime:        common.TimeToStr(message.CreateTime),
		})
	}
	c.JSON(http.StatusOK, common.SuccessResponse(tutorData))
}
