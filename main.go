package main

import (
	"github.com/gin-gonic/gin"
	_ "tutor_platform/src/config/code"
	"tutor_platform/src/router/router"
)

func main() {
	r := gin.Default()

	r.Static("static", "static")
	r.LoadHTMLGlob("views/**/*")

	router.CommonRouter(r)
	router.StudentRouter(r)
	router.RecruitRouter(r)

	r.Run(":8080")
}
