package config

import (
	"os"
	"strconv"
)

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}

var (
	MySQLUrl         string
	TimeLayout       string
	TokenRedisConf   RedisConf
	MessageRedisConf RedisConf
)

var (
	StudentIdentify       string
	RecruitIdentify       string
	StudentTokenName      string
	RecruitTokenName      string
	CityCodeKey           string
	ViewListKey           string
	ViewKey               string
	StudentAppointmentKey string
	MessageAppointmentKey string
)

func init() {
	TimeLayout = "2006-01-02 15:04:05"

	StudentTokenName = "studentToken"
	RecruitTokenName = "recruitToken"

	StudentAppointmentKey = "student:appointment"
	MessageAppointmentKey = "message:appointment"

	CityCodeKey = "citycode"
	ViewListKey = "view:list"
	ViewKey = "view"

	StudentIdentify = os.Getenv("STUDENT_IDENTIFY")
	if StudentIdentify == "" {
		StudentIdentify = "student"
	}
	RecruitIdentify = os.Getenv("RECRUIT_IDENTIFY")
	if RecruitIdentify == "" {
		RecruitIdentify = "recruit"
	}
	MySQLUrl = os.Getenv("MYSQL_URL")
	if MySQLUrl == "" {
		MySQLUrl = "root:1234@tcp(119.23.49.85:3306)/Graduation?charset=utf8"
	}

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		//redisUrl = "47.97.86.11:6378"
		redisUrl = "localhost:6379"
	}
	TokenRedisConf.Addr = redisUrl
	TokenRedisConf.Password = os.Getenv("REDIS_PASSWORD")
	tokenRedisDB := os.Getenv("TOKEN_REDIS_DB")
	if tokenRedisDB == "" {
		TokenRedisConf.DB = 8
	} else {
		TokenRedisConf.DB, _ = strconv.Atoi(tokenRedisDB)
	}

	MessageRedisConf.Addr = redisUrl
	MessageRedisConf.Password = os.Getenv("REDIS_PASSWORD")
	MessageRedisDB := os.Getenv("MESSAGE_REDIS_DB")
	if MessageRedisDB == "" {
		MessageRedisConf.DB = 5
	} else {
		MessageRedisConf.DB, _ = strconv.Atoi(MessageRedisDB)
	}
}
