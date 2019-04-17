package router

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/controller/messageController"
	"tutor_platform/src/controller/studentController"
	"tutor_platform/src/dao/studentDao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/router/publicRouterUrl"
	"tutor_platform/src/util/redis"
	"net/http"
	"time"
)

type login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func StudentRouter(router *gin.Engine) {
	var loginVal *commonData.Login
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "token",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*commonData.User); ok {
				return jwt.MapClaims{
					"token": v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &commonData.User{
				UserName: claims["token"].(string),
				UserId:   loginVal.UserId,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginInfo login
			if err := c.ShouldBind(&loginInfo); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginInfo.UserName
			password := loginInfo.Password

			loginVal = &commonData.Login{
				UserName: username,
				Password: password,
				Identify: config.StudentIdentify,
			}
			ok := studentDao.Login(loginVal)
			userId := loginVal.UserId
			if ok == true {
				accountId := common.GetMD5HashCode()
				redis.DelValue(userId)
				redis.AddValue(userId, accountId)
				return &commonData.User{
					UserName: loginVal.UserName,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*commonData.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:token",
		TimeFunc:    time.Now,
	})

	if err != nil {
		panic(err)
	}

	router.POST(publicRouterUrl.StudentLogin, authMiddleware.LoginHandler)
	//测试权限
	studentRouter := router.Group("student", authMiddleware.MiddlewareFunc())
	{
		studentRouter.GET(publicRouterUrl.StudentAuthorized, func(c *gin.Context) {
			c.String(http.StatusOK, "info")
		})
		studentRouter.GET(publicRouterUrl.StudentInfo, studentController.Info)
		studentRouter.GET(publicRouterUrl.StudentLogout, studentController.Logout)
		studentRouter.POST(publicRouterUrl.StudentUpdateLogin, studentController.UpdateLogin)
		studentRouter.POST(publicRouterUrl.StudentEntireOwnInfo, studentController.StudentEntireOwnInfo)
		studentRouter.GET(publicRouterUrl.AppointmentTutor, messageController.AppointmentTutor)
		studentRouter.GET(publicRouterUrl.JudgeAppointment, messageController.JudgeAppointment)
		studentRouter.GET(publicRouterUrl.StudentLookOwnAppointmentMessageList, studentController.AppointmentTutorList)
		//studentRouter.POST(publicRouterUrl.StudentUpdateOwnInfo, studentController.StudentUpdateOwnInfo)
		//studentRouter.GET(publicRouterUrl.StudentLookOwnInfo, studentController.StudentLookOwnInfo)
		//studentRouter.POST(publicRouterUrl.StudentLookNotify, studentController.StudentLookNotify)
		//studentRouter.POST(publicRouterUrl.StudentLookRecruitContact, studentController.StudentLookRecruitContact)
		//studentRouter.POST(publicRouterUrl.StudentLookMessageAppointmentNumber, studentController.StudentLookMessageAppointmentNumber)
		//studentRouter.POST(publicRouterUrl.StudentAppointmentMessage, studentController.StudentAppointmentMessage)
		//studentRouter.POST(publicRouterUrl.StudentDeleteAppointment, studentController.StudentDeleteAppointment)
		//studentRouter.POST(publicRouterUrl.StudentLookOwnAppointmentMessageList, studentController.StudentLookOwnAppointmentMessageList)
	}
}
