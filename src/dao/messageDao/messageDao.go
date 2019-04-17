package messageDao

import (
	"fmt"
	_ "github.com/go-sql-driver/MYSQL"
	"tutor_platform/src/common"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	"tutor_platform/src/config/sql"
	. "tutor_platform/src/dao"
	"tutor_platform/src/dao/commonDao"
	"tutor_platform/src/data/messageData"
	"tutor_platform/src/response"
	"tutor_platform/src/util/redis"
)

func GetTutorList(dataMapList []map[string][]byte) []response.TutorInfo {
	var tutorList = make([]response.TutorInfo, 0)
	for _, dataMap := range dataMapList {
		tutorInfo := response.TutorInfo{
			Mid:           string(dataMap["mid"]),
			StudentLevel:  string(dataMap["student_level"]),
			StudentGender: string(dataMap["student_gender"]),
			OwnAccountId:  string(dataMap["own_account_id"]),
			City:          string(dataMap["city"]),
			Local:         string(dataMap["local"]),
			Salary:        string(dataMap["salary"]),
			Describe:      string(dataMap["describe"]),
		}
		tutorList = append(tutorList, tutorInfo)
	}
	return tutorList
}

func GetMessageList(dataMapList []map[string][]byte) []messageData.Message {
	messageList := make([]messageData.Message, 0)
	for _, dataMap := range dataMapList {
		mid := string(dataMap["mid"])
		message := messageData.Message{
			Mid:               mid,
			Views:             common.StrToInt(redis.GetValue(fmt.Sprintf("%v:%v", config.ViewKey, mid))),
			AppointmentNumber: common.StrToInt(string(dataMap["appointment_number"])),
			CreateTime:        common.TimeFormat(string(dataMap["create_time"])),
		}
		messageList = append(messageList, message)
	}
	return messageList
}

func Query() interface{} {
	list, err := DB.Query("select * from tutor")
	if err != nil {
		return common.FailResponse(code.DBQueryErrorCode)
	}
	return common.SuccessResponse(GetTutorList(list))
}

func InsertMessage(message *messageData.Message) bool {
	has, _ := DB.Insert(message)
	return has == 1
}

func InsertTutor(tutor *messageData.Tutor) interface{} {
	has, _ := DB.Insert(tutor)
	if has == 0 {
		return common.FailResponse(code.DBInsertErrorCode)
	}
	return common.SuccessResponse(tutor)
}

func FindTutor(tutor *messageData.Tutor) bool {
	has, _ := DB.Get(tutor)
	return has
}

func FindMessage(message *messageData.Message) bool {
	has, _ := DB.Get(message)
	return has
}

func UpdateMessage(message *messageData.Message) bool {
	_, err := DB.Id(message.Mid).Update(message)
	if err != nil {
		return false
	}
	return true
}

func UpdateTutor(tutor *messageData.Tutor) interface{} {
	has, _ := DB.Update(tutor)
	if has == 0 {
		return common.FailResponse(code.DBUpdateErrorCode)
	}
	return common.SuccessResponse(tutor)
}

func DeleteMessage(message *messageData.Message) bool {
	has, _ := DB.Delete(message)
	return has != 0
}

func DeleteTutor(tutor *messageData.Tutor) bool {
	has, _ := DB.Delete(tutor)
	return has != 0
}

func BatchQueryMessage(idList []string) []messageData.Message {
	res, err := DB.Query(sql.QueryMessageSQL + sql.QueryBatchIDSQL(idList, "mid"))
	messageList := make([]messageData.Message, 0)
	if err != nil {
		return messageList
	}
	for _, dataMap := range res {
		mid := string(dataMap["mid"])
		views := common.StrToInt(redis.GetValue(fmt.Sprintf("%v:%v", config.ViewKey, mid)))
		appointmentNumber := common.StrToInt(string(dataMap["appointment_number"]))
		createTime := string(dataMap["create_time"])
		messageList = append(messageList, messageData.Message{
			Mid:               mid,
			Views:             views,
			AppointmentNumber: appointmentNumber,
			CreateTime:        common.TimeFormat(createTime),
		})
	}
	return messageList
}

func BatchQueryTutor(idList []string) []response.TutorInfo {
	res, err := DB.Query(sql.QueryTutorSQL + sql.QueryBatchIDSQL(idList, "mid"))
	tutorList := make([]response.TutorInfo, 0)
	if err != nil {
		return tutorList
	}
	return GetTutorList(res)
}

func QueryWeekTutor() interface{} {
	res, err := DB.Query(sql.QueryWeekSQL())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	idList := make([]string, 0)
	for _, dataMap := range res {
		mid := string(dataMap["mid"])
		idList = append(idList, mid)
	}
	sql := sql.QueryTutorSQL + sql.QueryBatchIDSQL(idList, "mid")
	list, err := DB.Query(sql)
	tutorList := GetTutorList(list)
	infoList := make([]response.TutorInfo, 0)
	for _, dataMap := range res {
		mid := string(dataMap["mid"])
		views := common.StrToInt(redis.GetValue(fmt.Sprintf("%v:%v", config.ViewKey, mid)))
		appointmentNumber := common.StrToInt(string(dataMap["appointment_number"]))
		createTime := string(dataMap["create_time"])
		tutorInfo := GetTutorInfo(mid, tutorList)
		tutorInfo.CreateTime = createTime
		tutorInfo.Views = views
		tutorInfo.AppointmentNumber = appointmentNumber
		infoList = append(infoList, tutorInfo)
	}
	return common.SuccessResponse(response.TutorResponse{
		List: infoList,
	})
}

func GetTutorInfo(mid string, tutorList []response.TutorInfo) response.TutorInfo {
	for _, tutor := range tutorList {
		if mid == tutor.Mid {
			return tutor
		}
	}
	return response.TutorInfo{}
}

func QueryTop10Tutor(city string) interface{} {
	key := fmt.Sprintf("%v:%v:%v", config.CityCodeKey, commonDao.QueryCity(city), config.ViewListKey)
	idViewMap := redis.ZRevRange(key, 10)
	idList := make([]string, 0)
	for id := range idViewMap {
		idList = append(idList, id)
	}
	messageList := BatchQueryMessage(idList)
	tutorInfoList := BatchQueryTutor(idList)
	tutorResponseList := make([]response.TutorInfo, 0)
	for _, message := range messageList {
		tutorInfo := GetTutorInfo(message.Mid, tutorInfoList)
		tutorInfo.CreateTime = common.TimeToStr(message.CreateTime)
		tutorInfo.Views = idViewMap[message.Mid]
		tutorInfo.AppointmentNumber = message.AppointmentNumber
		tutorResponseList = append(tutorResponseList, tutorInfo)
	}
	return common.SuccessResponse(response.TutorResponse{
		List: tutorResponseList,
	})
}

func QueryPageTutor(city string, page, pageSize int) interface{} {
	res, err := DB.Query(sql.QueryPageWhereCitySQL(city, page, pageSize))
	if err != nil {
		return common.FailResponse(code.DBQueryErrorCode)
	}
	messageList := GetMessageList(res)
	idList := make([]string, 0)
	for _, message := range messageList {
		idList = append(idList, message.Mid)
	}
	tutorInfoList := BatchQueryTutor(idList)
	tutorResponseList := make([]response.TutorInfo, 0)
	for _, message := range messageList {
		views := common.StrToInt(redis.GetValue(fmt.Sprintf("%v:%v", config.ViewKey, message.Mid)))
		tutorInfo := GetTutorInfo(message.Mid, tutorInfoList)
		tutorInfo.CreateTime = common.TimeToStr(message.CreateTime)
		tutorInfo.Views = views
		tutorInfo.AppointmentNumber = message.AppointmentNumber
		tutorResponseList = append(tutorResponseList, tutorInfo)
	}
	return common.SuccessResponse(response.TutorResponse{
		List: tutorResponseList,
	})
}

func QueryMessageList(recruitId string, page, pageSize int) interface{} {
	res, err := DB.Query(fmt.Sprintf("select * from message where own_account_id=%v order by create_time desc limit %v,%v", recruitId, page, pageSize))
	if err != nil {
		return common.FailResponse(code.DBQueryErrorCode)
	}
	messageList := GetMessageList(res)
	idList := make([]string, 0)
	for _, message := range messageList {
		idList = append(idList, message.Mid)
	}
	tutorInfoList := BatchQueryTutor(idList)
	tutorResponseList := make([]response.TutorInfo, 0)
	for _, message := range messageList {
		views := common.StrToInt(redis.GetValue(fmt.Sprintf("%v:%v", config.ViewKey, message.Mid)))
		tutorInfo := GetTutorInfo(message.Mid, tutorInfoList)
		tutorInfo.CreateTime = common.TimeToStr(message.CreateTime)
		tutorInfo.Views = views
		tutorInfo.AppointmentNumber = message.AppointmentNumber
		tutorResponseList = append(tutorResponseList, tutorInfo)
	}
	return common.SuccessResponse(response.TutorResponse{
		List: tutorResponseList,
	})
}
