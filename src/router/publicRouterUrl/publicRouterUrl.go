package publicRouterUrl

const (
	Index           = "/"
	Tutor           = "/tutor"
	Waiter          = "/waiter"
	Takeaway        = "/takeaway"
	TemporaryWorker = "/temporary_worker"
	Contact         = "/contact"

	StudentAuthorized = "/student"
	//学生用户注册
	StudentRegister = "/student/register"
	//学生用户登录
	StudentLogin = "/student/login"
	//学生用户退出登录
	StudentLogout = "/logout"
	//学生个人中心
	StudentPersonal = "student/personal"
	//学生姓名查看
	StudentInfo = "/info"
	//学生用户修改登录昵称密码
	StudentUpdateLogin = "/update_login"
	//学生用户完善信息
	StudentEntireOwnInfo = "/entire_own_info"
	//学生用户修改信息
	StudentUpdateOwnInfo = "/update_own_info"
	//学生用户查看信息
	StudentLookOwnInfo = "/look_own_info"
	//学生用户查看通知
	StudentLookNotify = "/look_notify"
	//学生用户查看招聘用户联系方式
	StudentLookRecruitContact = "/look_recruit_contact"

	RecruitAuthorized = "/recruit"
	//招聘用户注册
	RecruitRegister = "/recruit/register"
	//招聘用户登录
	RecruitLogin = "/recruit/login"
	//学生用户退出登录
	RecruitLogout = "/logout"
	//招聘用户个人中心
	RecruitPersonal = "recruit/personal"
	//学生姓名查看
	RecruitInfo = "/info"
	//招聘用户修改登录密码
	RecruitUpdateLogin = "/update_login"
	//招聘用户查看自己信息
	RecruitLookOwnInfo = "/look_own_info"
	//招聘用户完善信息
	RecruitEntireOwnInfo = "/entire_own_info"
	//招聘用户修改信息
	RecruitUpdateOwnInfo = "/update_own_info"
	//招聘用户查看通知
	RecruitLookNotify = "/look_notify"
	//招聘用户查看招聘信息被预约人联系方式
	RecruitLookStudentContact = "/look_student_contact"

	//招聘用户查看学生用户信息
	LookStudentInfo = "/look_student_info"
	//学生用户查看招聘用户信息
	LookRecruitInfo = "/look_recruit_info"
	//查看发布信息
	LookMessage = "/look_message"
	//查看单个信息
	QueryTutor = "/query_tutor"
	//进入当个信息页面
	TutorPage = "/tutor_page"
	//查看近七天内信息
	QueryWeekTutor = "/query_week_tutor"
	//查看本市内浏览量最大的十条信息
	QueryTop10Tutor = "/query_top_tutor"
	//按时间翻页查看信息
	QueryPageTutor = "/query_page_tutor"

	//招聘用户填写发布信息
	RecruitEntireMessage = "/entire_message"
	//招聘用户修改发布信息
	RecruitUpdateMessage = "/update_message"
	//招聘用户删除发布信息
	RecruitDeleteMessage = "/delete_message"
	//招聘用户查看已发的招聘信息列表
	RecruitLookMessageList = "/look_message_list"
	//招聘用户查看发布信息
	RecruitLookMessage = "/look_message"

	//用户城市切换
	SelectCity = "/select_city"
	//学生用户预约招聘信息
	AppointmentTutor = "/appointment_tutor"
	//学生用户取消预约

	//判断学生是否已经预约该信息
	JudgeAppointment = "/judge_appointment"
	//学生用户查看招聘信息
	StudentLookMessageAppointmentNumber = "/look_message_appointment_number"
	//学生用户查看招聘信息被预约人数
	StudentAppointmentMessage = "/appointment_message"
	//学生用户撤销预约
	StudentDeleteAppointment = "/delete_appointment"
	//学生用户查看已预约信息列表
	StudentLookOwnAppointmentMessageList = "/look_own_appointment_message_list"

	//招聘用户查看招聘信息被预约人数列表
	RecruitLookMessageAppointmentList = "/look_message_appointment_list"
	//招聘用户选择招聘信息被预约人
	RecruitSelectMessageTransStudent = "/select_message_trans_student"
)
