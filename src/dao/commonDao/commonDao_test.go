package commonDao

import (
	"testing"
	"tutor_platform/src/config"
	"tutor_platform/src/data/commonData"
)

func TestQueryCity(t *testing.T) {
	QueryCity("")
}

func TestQueryLogin(t *testing.T) {
	login := commonData.Login{}
	QueryLogin(&login, config.StudentIdentify)
}

func TestExistLogin(t *testing.T) {
	login := commonData.Login{}
	ExistLogin(&login)
}
