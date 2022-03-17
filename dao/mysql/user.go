package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"image/model"
)

const secret = "weibo"

func encryPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	err = db.Get(&count, sqlStr, username)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count > 0 {
		// 用户已经存在
		return ErrorUserExit
	}
	return
}

func InsertUser(user *model.User) (err error) {
	// 对密码进行md5加密
	user.PassWord = encryPassword([]byte(user.PassWord))
	// 执行sql入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?);`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.PassWord)
	return
}

func Login(user *model.User) (err error) {
	// 记录一下原始密码
	originPassword := user.PassWord
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		// 查询数据出错
		return
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return ErrorUserNotExit
	}
	// 生成加密密码与查询到的密码比较
	password := encryPassword([]byte(originPassword))
	if password != user.PassWord {
		return ErrorPasswordWrong
	}
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(userId int64) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, userId)
	fmt.Println(err)
	return
}
