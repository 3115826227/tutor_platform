package messageData

import "time"

type Message struct {
	Mid                 string    `xorm:"pk"`
	TypeCode            int       `json:"type_code",xorm:"type_code"`
	Name                string    `json:"name",xorm:"name"`
	City                string    `json:"city",xorm:"city"`
	OwnAccountId        string    `json:"own_account_id",xorm:"own_account_id"`
	Views               int       `json:"views",xorm:"views"`
	AppointmentNumber   int       `json:"appointment_number",xorm:"appointment_number"`
	AppointmentUserList string    `json:"appointmen_user_list",xorm:"appointment_user_list"`
	Status              string    `json:"status",xorm:"status"`
	CreateTime          time.Time `json:"create_time",xorm:"created_time"`
	UpdateTime          time.Time `json:"update_time",xorm:"updated_time"`
}

type Tutor struct {
	Mid           string `xorm:"pk"`
	TypeCode      int    `json:"type_code",xorm:"type_code"`
	Name          string `json:"name",xorm:"name"`
	City          string `json:"city",xorm:"city"`
	Local         string `json:"local",xorm:"local"`
	OwnAccountId  string `json:"own_account_id",xorm:"own_account_id"`
	StudentGender string `json:"student_gender",xorm:"student_gender"`
	StudentLevel  string `json:"student_level",xorm:"student_level"`
	Salary        string `json:"salary",xorm:"salary"`
	Describe      string `json:"describe",xorm:"describe"`
}
