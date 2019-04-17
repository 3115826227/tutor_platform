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
	data := QueryAppointmentMessageList("88034197", 1, 20)
	fmt.Println(data)
}
