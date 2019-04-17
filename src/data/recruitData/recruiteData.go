package recruitData

import "time"

type Recruit struct {
	UserId     string    `json:"user_id",xorm:"user_id"`
	Name       string    `json:"name",xorm:"name"`
	Gender     string    `json:"gender",xorm:"gender"`
	City       string    `json:"city",xorm:"city"`
	Identify   string    `json:"identify",xorm:"identify"`
	Phone      string    `json:"phone",xorm:"phone"`
	CreateTime time.Time `json:"create_time",xorm:"created_time"`
	UpdateTime time.Time `json:"update_time",xorm:"updated_time"`
}
