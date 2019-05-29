package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"tutor_platform/src/config"
	"tutor_platform/src/config/code"
	"tutor_platform/src/response"
)

func SetToken(TokenString string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(TokenString))
	if err != nil {
		return ""
	}
	return tokenString
}

func TimeToTime(t time.Time) string {
	return time.Unix(t.Unix(), 64).Format(config.TimeLayout)
}

func TimeFormat(str string) time.Time {
	t, _ := time.Parse(config.TimeLayout, str)
	return t
}

func TimeToStr(t time.Time) string {
	return t.Format(config.TimeLayout)
}

func UnixToStr(timeInt int64) string {
	return time.Unix(timeInt, 64).Format(config.TimeLayout)
}

func SuccessResponse(data interface{}) interface{} {
	return response.Response{
		Code: http.StatusOK,
		Data: data,
	}
}

func FailResponse(Code int) interface{} {
	return response.Response{
		Code:    Code,
		Message: code.ErrorMap[Code],
	}
}

func StrToInt(s string) int {
	data, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return data
}

func IntToString(data int) string {
	return strconv.Itoa(data)
}

func GetTime() time.Time {
	t, _ := time.Parse(config.TimeLayout, time.Now().Format(config.TimeLayout))
	return t
}

func GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := r.Intn(10)
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetUserId() string {
	str := fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))
	if str[0] == '0' {
		str = GetUserId()
	}
	fmt.Println(str)
	return str
}

func GetMD5HashCode() string {
	message, _ := json.Marshal(GetRandomString())
	hash := md5.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
}

func GetCity(data string) response.City {
	var city response.City
	cityByte := bytes.NewBufferString(data).Bytes()
	json.Unmarshal(cityByte, &city)
	return city
}

func StrToInt64(s string) int64 {
	data, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return int64(data)
}
