package controller

import (
	"github.com/gin-gonic/gin"
	"tutor_platform/src/config"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, config.IndexHTML, nil)
}

func StudentLogin(c *gin.Context) {
	c.HTML(http.StatusOK, config.StudentLoginHTML, nil)
}

func RecruitLogin(c *gin.Context) {
	c.HTML(http.StatusOK, config.RecruitLoginHTML, nil)
}

func StudentRegister(c *gin.Context) {
	c.HTML(http.StatusOK, config.StudentRegisterHTML, nil)
}

func RecruitRegister(c *gin.Context) {
	c.HTML(http.StatusOK, config.RecruitRegisterHTML, nil)
}

func Tutor(c *gin.Context) {
	c.HTML(http.StatusOK, config.TutorHTML, nil)
}

func Waiter(c *gin.Context) {
	c.HTML(http.StatusOK, config.WaiterHTML, nil)
}

func Takeaway(c *gin.Context) {
	c.HTML(http.StatusOK, config.TakeawayHTML, nil)
}

func TemporaryWorker(c *gin.Context) {
	c.HTML(http.StatusOK, config.TemporaryWorkerHTML, nil)
}

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, config.ContactHTML, nil)
}

func TutorPage(c *gin.Context) {
	c.HTML(http.StatusOK, config.TutorHTML, nil)
}
