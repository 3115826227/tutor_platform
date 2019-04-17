package commonData

import "time"

type Login struct {
	UserId     string    `xorm:"user_id"`
	UserName   string    `xorm:"username"`
	Password   string    `xorm:"password"`
	Identify   string    `xorm:"identify"`
	CreateTime time.Time `xorm:"created_time"`
	UpdateTime time.Time `xorm:"updated_time"`
}

type User struct {
	UserId     string    `json:"user_id",xorm:"user_id"`
	UserName   string    `json:"username",xorm:"username"`
	Name       string    `json:"name",xorm:"name"`
	Gender     string    `json:"gender",xorm:"gender"`
	City       string    `json:"city",xorm:"city"`
	Identify   string    `json:"identify",xorm:"identify"`
	Phone      string    `json:"phone",xorm:"phone"`
	CreateTime time.Time `json:"create_time",xorm:"created_time"`
	UpdateTime time.Time `json:"update_time",xorm:"updated_time"`
}

type City struct {
	CityCode string `json:"city_code",xorm:"city_code"`
	City     string `json:"city",xorm:"city"`
}
