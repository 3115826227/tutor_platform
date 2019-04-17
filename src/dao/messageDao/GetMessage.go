package messageDao

import (
	"fmt"
	"github.com/go-xorm/xorm"
	. "tutor_platform/src/dao"
	"tutor_platform/src/data/messageData"
	"time"
)

func GetMessage() {
	db, err := xorm.NewEngine("mysql", "root:1234@tcp(localhost:3307)/info?charset=utf8")
	if err != nil {
		panic(err)
	}
	res, err := db.Query("select * from info")
	if err != nil {
		panic(err)
	}
	for _, dataMap := range res {
		id := string(dataMap["id"])
		level := string(dataMap["level"])
		gender := string(dataMap["gender"])
		local := string(dataMap["local"])
		salary := string(dataMap["salary"])
		describe := string(dataMap["describe"])
		message := &messageData.Message{
			Mid:          id,
			TypeCode:     1,
			OwnAccountId: "88034197",
			Status:       "发布中",
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		}
		DB.Insert(message)
		tutor := &messageData.Tutor{
			Mid:           id,
			TypeCode:      1,
			City:          "上海市",
			Local:         local,
			OwnAccountId:  message.OwnAccountId,
			StudentGender: gender,
			StudentLevel:  level,
			Salary:        salary,
			Describe:      describe,
		}
		DB.Insert(tutor)
	}
}

func UpdateMessageInfo() {
	res, err := DB.Query("select mid from message")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dataMap := range res {
		mid := string(dataMap["mid"])
		message := messageData.Message{
			Mid:  mid,
			City: "上海市",
		}
		DB.Id(message.Mid).Update(&message)
	}
}
