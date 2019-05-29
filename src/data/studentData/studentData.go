package studentData

import "time"

//用户完善信息表
type Student struct {
	UserId               string    `json:"user_id",xorm:"user_id"`
	Name                 string    `json:"name",xorm:"name"`
	Gender               string    `json:"gender",xorm:"gender"`
	City                 string    `json:"city",xorm:"city"`
	Identify             string    `json:"identify",xorm:"identify"`
	Phone                string    `json:"phone",xorm:"phone"`
	StudyMajor           string    `json:"study_major",xorm:"study_major"`
	School               string    `json:"school",xorm:"school"`
	Certificate          string    `json:"certificate",xorm:"certificate"`
	Hobby                string    `json:"hobby",xorm:"hobby"`
	Describe             string    `json:"describe",xorm:"describe"`
	EducationLevel       string    `json:"education_level",xorm:"education_level"`
	JobStatus            string    `json:"job_status",xorm:"job_status"`
	JobCategory          string    `json:"job_category",xorm:"job_category"`
	JobDesirableSalary   int64     `json:"job_desirable_salary",xorm:"job_desirable_salary"`
	AdvantagesSubject    string    `json:"advantage_subject",xorm:"advantages_subject"`
	AppointmentTutorList string    `json:"appointment_tutor_list",xorm:"appointment_tutor_list"`
	CreateTime           time.Time `json:"create_time",xorm:"created_time"`
	UpdateTime           time.Time `json:"update_time",xorm:"updated_time"`
}

//用户操作记录表
type StudentOperationInfo struct {
	UserId                 int64     `json:"user_id",xorm:"user_id"`
	LookNumber             int       `json:"look_number",xorm:"look_number"`
	AppointmentList        string    `json:"appointment_list",xorm:"appointment_list"`
	AppointmentSuccessList string    `json:"appointmentSuccess_list",xorm:"appointment_success_list"`
	CreateTime             time.Time `json:"create_time",xorm:"created_time"`
	UpdateTime             time.Time `json:"update_time",xorm:"updated_time"`
}

//用户工作记录表
type StudentWorkInfo struct {
	UserId           int64     `json:"user_id",xorm:"user_id"`
	LastTransTime    time.Time `json:"last_trans_time",xorm:"last_trans_time"`
	LastTransMessage time.Time `json:"last_trans_message",xorm:"last_trans_message"`
	CreateTime       time.Time `json:"create_time",xorm:"created_time"`
	UpdateTime       time.Time `json:"update_time",xorm:"updated_time"`
}
