package dao

import (
	_ "github.com/go-sql-driver/MYSQL"
	"github.com/go-xorm/xorm"
	"tutor_platform/src/config"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/data/messageData"
	"tutor_platform/src/data/recruitData"
	"tutor_platform/src/data/studentData"
)

var DB *xorm.Engine

func init() {
	var err error
	DB, err = xorm.NewEngine("mysql", config.MySQLUrl)
	if err != nil {
		panic(err)
	}
	DB.ShowSQL(true)
	err = DB.Sync2(
		new(commonData.Login),
		new(commonData.User),
		new(studentData.Student),
		new(recruitData.Recruit),
		new(messageData.Message),
		new(messageData.Tutor),
		new(commonData.City),
	)
}
