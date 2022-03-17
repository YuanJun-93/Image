package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
)

type User struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

//func TestRegisterHandler(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	r := gin.Default()
//	url := "/api/v1/register"
//	r.POST(url, RegisterHandler)
//
//	body := `{
//		"username": "user2",
//		"password": "123456",
//		"confirm_password": "123456"
//	}`
//
//	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
//	//req.Header.Set("Content-Type", "application/json")
//	// 创建回写
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, 200, w.Code)
//
//	// 判断是不是预期返回的
//	// 方法1
//	assert.Contains(t, w.Body.String(), "success")
//
//	// 方法2 将响应内容反序列化到ResponseData 然后判断字段是否与预期一致
//	//res := new(ResponseData)
//	//if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
//	//	t.Fatalf("json unmarshal failed, err:%v\n", err)
//	//}
//	//assert.Equal(t, res.Message, "success")
//}

func TestRegisterHandler(t *testing.T) {
	for i := 1001; i <= 3000; i++ {
		//strconv.Itoa(i)
		body := `{
		"username": "user` + strconv.Itoa(i) + `",
		"password": "123456",
		"confirm_password": "123456"
		}`
		//fmt.Println(body)
		_, err := http.Post("http://127.0.0.1:8080/api/v1/register", "application/json", bytes.NewReader([]byte(body)))
		if err != nil {
			fmt.Println(err)
		}
	}
}

type Data struct {
	Token    string `json:"token"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type ResponseDatas struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	*Data   `json:"data"`
}

func TestLoginHandler(t *testing.T) {
	file, err := os.OpenFile("token.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for i := 1; i <= 3000; i++ {
		rsp := ResponseDatas{}
		//strconv.Itoa(i)
		body := `{
		"username": "user` + strconv.Itoa(i) + `",
		"password": "123456"
		}`
		//fmt.Println(body)
		req, err := http.Post("http://127.0.0.1:8080/api/v1/login", "application/json", bytes.NewReader([]byte(body)))
		if err != nil {
			fmt.Println(err)
		}

		responseBody, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(responseBody, &rsp)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(rsp.Token)
		_, err = file.WriteString(rsp.Token + "\n")
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(count)
	}
}
