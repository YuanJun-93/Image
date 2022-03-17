package logic

import (
	"image/dao/mysql"
	"image/model"
	"image/pkg/jwt"
	"image/pkg/snowflake"
)

func Register(p *model.ParamRegisterUser) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.UserName); err != nil {
		return err
	}
	//if p.PassWord != p.ConfirmPassword {
	//	err = errors.New("两次输入的密码不正确")
	//}
	// 2. 生成userID
	userID := snowflake.GenID()
	// 构造一个用户实例
	var user = &model.User{
		UserID:   userID,
		UserName: p.UserName,
		PassWord: p.PassWord,
	}
	// 3. 保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *model.User) (user *model.User, err error) {
	user = &model.User{
		UserName: p.UserName,
		PassWord: p.PassWord,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}
