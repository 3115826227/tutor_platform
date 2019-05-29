package redis

import (
	"fmt"
	"testing"
)

func TestGetValue(t *testing.T) {
	val := GetValue("test")
	fmt.Println(val)
}

func TestAddValue(t *testing.T) {
	AddValue("hello", "world")
}

func TestDelValue(t *testing.T) {
	DelValue("0")
}

func TestZRevRange(t *testing.T) {

}

func TestSAdd(t *testing.T) {
	SAdd("test", "123")
}

func TestSJudgeExist(t *testing.T) {
	ok := SJudgeExist("test", "123")
	fmt.Println(ok)
}

func TestZRev(t *testing.T) {
}

func TestSDel(t *testing.T) {
	SDel("test", "123")
}
