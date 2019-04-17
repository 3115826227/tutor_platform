package studentController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/controller/commonController"
	"tutor_platform/src/dao/studentDao"
	"tutor_platform/src/data/studentData"
	"tutor_platform/src/util/redis"
	"net/http"
)

func StudentPersonal(c *gin.Context) {
	commonController.Personal(c, config.StudentPersonalHTML, config.IndexHTML)
}

func Info(c *gin.Context) {
	c.JSON(http.StatusOK, commonController.GetInfo(c))
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, commonController.Logout(c))
}

func Register(c *gin.Context) {
	commonController.Register(c, config.StudentIdentify)
}

func UpdateLogin(c *gin.Context) {
	commonController.UpdateLogin(c, config.StudentIdentify)
}

func StudentEntireOwnInfo(c *gin.Context) {
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
		JobDesirableSalary: jobDesirableSalary,
		AdvantagesSubject:  advantagesSubject,
	}
	c.JSON(http.StatusOK, studentDao.UpdateStudentInfo(&student))
}

func AppointmentTutorList(c *gin.Context) {
	userId := commonController.GetId(c)
	studentKey := fmt.Sprintf("%v:%v", config.StudentAppointmentKey, userId)
	tutorList := redis.SMembers(studentKey)
	c.JSON(http.StatusOK, common.SuccessResponse(tutorList))
}
