package router

import (
	"github.com/gin-gonic/gin"
	"tutor_platform/src/controller"
	"tutor_platform/src/controller/messageController"
	"tutor_platform/src/controller/recruitController"
	"tutor_platform/src/controller/studentController"
	"tutor_platform/src/router/publicRouterUrl"
	"net/http"
)

func CommonRouter(commonRouter *gin.Engine) {
	commonRouter.GET("/info", func(c *gin.Context) {
		c.String(http.StatusOK, "info")
	})

	commonRouter.GET(publicRouterUrl.Index, controller.Index)
	commonRouter.GET(publicRouterUrl.StudentLogin, controller.StudentLogin)
	commonRouter.GET(publicRouterUrl.RecruitLogin, controller.RecruitLogin)
	commonRouter.GET(publicRouterUrl.StudentRegister, controller.StudentRegister)
	commonRouter.GET(publicRouterUrl.RecruitRegister, controller.RecruitRegister)
	commonRouter.GET(publicRouterUrl.Tutor, controller.Tutor)
	commonRouter.GET(publicRouterUrl.Waiter, controller.Waiter)
	commonRouter.GET(publicRouterUrl.Takeaway, controller.Takeaway)
	commonRouter.GET(publicRouterUrl.TemporaryWorker, controller.TemporaryWorker)
	commonRouter.GET(publicRouterUrl.Contact, controller.Contact)
	commonRouter.GET(publicRouterUrl.StudentPersonal, studentController.StudentPersonal)
	commonRouter.GET(publicRouterUrl.RecruitPersonal, recruitController.RecruitPersonal)
	commonRouter.GET(publicRouterUrl.TutorPage, controller.TutorPage)

	commonRouter.POST(publicRouterUrl.StudentRegister, studentController.Register)
	commonRouter.POST(publicRouterUrl.RecruitRegister, recruitController.Register)
	commonRouter.POST(publicRouterUrl.LookStudentInfo, messageController.LookStudentInfo)
	commonRouter.POST(publicRouterUrl.LookRecruitInfo, messageController.LookRecruitInfo)
	commonRouter.GET(publicRouterUrl.QueryTutor, messageController.QueryTutor)
	commonRouter.GET(publicRouterUrl.QueryWeekTutor, messageController.QueryWeekTutor)
	commonRouter.GET(publicRouterUrl.QueryTop10Tutor, messageController.QueryTop10Tutor)
	commonRouter.GET(publicRouterUrl.QueryPageTutor, messageController.QueryPageTutor)
	//commonRouter.POST(publicRouterUrl.SelectCity, messageController.SelectCity)
}
