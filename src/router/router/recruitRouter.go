package router

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/controller/messageController"
	"tutor_platform/src/controller/recruitController"
	"tutor_platform/src/dao/recruitDao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/router/publicRouterUrl"
	"tutor_platform/src/util/redis"
)

func RecruitRouter(router *gin.Engine) {
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
			var login login
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := login.UserName
			password := login.Password

			loginVal = &commonData.Login{
				UserName: username,
				Password: password,
				Identify: config.RecruitIdentify,
			}
			ok := recruitDao.Login(loginVal)
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
	router.POST(publicRouterUrl.RecruitLogin, authMiddleware.LoginHandler)
	//测试权限
	recruitRouter := router.Group("recruit", authMiddleware.MiddlewareFunc())
	{
		recruitRouter.GET(publicRouterUrl.RecruitAuthorized, func(c *gin.Context) {
			c.String(http.StatusOK, "info")
		})
		recruitRouter.GET(publicRouterUrl.RecruitInfo, recruitController.BaseInfo)
		recruitRouter.GET(publicRouterUrl.RecruitLogout, recruitController.Logout)
		recruitRouter.POST(publicRouterUrl.RecruitUpdateLogin, recruitController.UpdateLogin)
		recruitRouter.GET(publicRouterUrl.RecruitLookOwnInfo, recruitController.RecruitLookOwnInfo)
		recruitRouter.POST(publicRouterUrl.RecruitUpdateOwnInfo, recruitController.RecruitUpdateOwnInfo)
		recruitRouter.GET(publicRouterUrl.RecruitLookStudentContact, recruitController.RecruitLookStudentContact)
		recruitRouter.POST(publicRouterUrl.RecruitEntireMessage, messageController.AddTutor)
		recruitRouter.POST(publicRouterUrl.RecruitUpdateMessage, messageController.UpdateTutor)
		recruitRouter.GET(publicRouterUrl.RecruitDeleteMessage, messageController.DeleteTutor)
		recruitRouter.GET(publicRouterUrl.RecruitLookMessageList, recruitController.RecruitLookMessageList)
		recruitRouter.GET(publicRouterUrl.RecruitLookMessageAppointmentList, recruitController.RecruitLookMessageAppointmentList)
		//recruitRouter.POST(publicRouterUrl.RecruitSelectMessageTransStudent, recruitCotroller.RecruitSelectMessageTransStudent)
	}
}
