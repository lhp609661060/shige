package models

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type User struct {
	UserId       string
	Username string
	Password string
	Mobile   string
	Email    string
	Age      int
	Sex      string
	Ip       string
	LastSignInTime string
	CreateTime string
	SignCount int
}

type UserSignInInfo struct {
	UserId string
	CreateTime string
	Ip string
	SignSource string
}


func AuthUser(username string, password string) (bool, string, string) {
	c := MgoDb("shige_hy").C("user")
	user := User{}

	err := c.Find(bson.M{"username": username}).One(&user)
	if err != nil{
		return false, "username", "用户名错误"
	}

	fmt.Println(user.Password)
	fmt.Println(password)

	if user.Password != password {
		return false, "password", "密码输入错误"
	}

	return true, "none", "登录成功"
}