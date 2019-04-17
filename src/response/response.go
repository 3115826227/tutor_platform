package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type City struct {
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
}

type UserInfoResponse struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id"`
	Location  City   `json:"location"`
	Identify  string `json:"identify"`
}

type StudentInfoResponse struct {
	UserId             string `json:"user_id"`
	Name               string `json:"name"`
	Location           City   `json:"location"`
	Identify           string `json:"identify"`
	Gender             string `json:"gender"`
	StudyMajor         string `json:"study_major"`
	School             string `json:"school"`
	Certificate        string `json:"certificate"`
	Hobby              string `json:"hobby"`
	Phone              string `json:"phone"`
	Describe           string `json:"describe"`
	EducationLevel     string `json:"education_level"`
	JobStatus          string `json:"job_status"`
	JobCategory        string `json:"job_category"`
	JobDesirableSalary int    `json:"job_desirable_salary"`
	AdvantagesSubject  string `json:"advantage_subject"`
}

type RecruitInfoResponse struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	Location City   `json:"location"`
	Identify string `json:"identify"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
}

type TutorResponse struct {
	List []TutorInfo `json:"list"`
}

type TutorInfo struct {
	Mid               string `json:"m_id"`
	Name              string `json:"name"`
	City              string `json:"city"`
	Local             string `json:"local"`
	OwnAccountId      string `json:"own_account_id"`
	StudentGender     string `json:"student_gender"`
	StudentLevel      string `json:"student_level"`
	Salary            string `json:"salary"`
	Describe          string `json:"describe"`
	Views             int    `json:"views"`
	AppointmentNumber int    `json:"appointment_number"`
	CreateTime        string `json:"create_time"`
}

type PhoneInfo struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
}
