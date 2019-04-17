package studentDao

import (
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	"tutor_platform/src/config/sql"
	. "tutor_platform/src/dao"
	"tutor_platform/src/dao/commonDao"
	"tutor_platform/src/data/commonData"
	"tutor_platform/src/data/studentData"
	"tutor_platform/src/response"
)

func Login(login *commonData.Login) bool {
	return commonDao.QueryLogin(login, config.StudentIdentify)
}

func Insert(student *studentData.Student) interface{} {
	_, err := DB.Insert(student)
	if err != nil {
		return common.FailResponse(code.DBInsertErrorCode)
	}
	return common.SuccessResponse(student)
}

func Query(student *studentData.Student) bool {
	has, _ := DB.Get(student)
	return has
}

func UpdateStudentInfo(student *studentData.Student) interface{} {
	_, err := DB.Update(student)
	if err != nil {
		return common.FailResponse(code.DBInsertErrorCode)
	}
	return common.SuccessResponse(student)
}

func GetStudentList(dataMapList []map[string][]byte) []response.StudentInfoResponse {
	var studentList = make([]response.StudentInfoResponse, 0)
	for _, dataMap := range dataMapList {
		studentInfo := response.StudentInfoResponse{
			UserId: string(dataMap["user_id"]),
			Name:   string(dataMap["name"]),
		}
		studentList = append(studentList, studentInfo)
	}
	return studentList
}

func BatchQueryStudentInfo(idList []string) interface{} {
	sql := sql.QueryBatchStudentSQL(idList)
	res, err := DB.Query(sql)
	if err != nil {
		return common.FailResponse(code.DBQueryErrorCode)
	}
	return common.SuccessResponse(GetStudentList(res))
}
