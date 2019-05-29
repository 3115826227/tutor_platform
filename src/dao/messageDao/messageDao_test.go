package messageDao

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	resp := Query()
	fmt.Println(resp)
}

func TestGetMessage(t *testing.T) {
	GetMessage()
}

func TestUpdateMessageInfo(t *testing.T) {
	UpdateMessageInfo()
}

func TestQueryAppointmentMessageList(t *testing.T) {
}

func TestQueryTop10Tutor(t *testing.T) {
	QueryTop10Tutor("上海市")
}

func TestQuerySalary(t *testing.T) {
	QuerySalary("上海市")
}
