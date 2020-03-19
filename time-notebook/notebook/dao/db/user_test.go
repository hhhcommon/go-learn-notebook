package db

import (
	"fmt"
	"testing"
	"time"	
	_ "github.com/go-sql-driver/mysql"
	"github.com/panda8z/notebook/model"
)

func init() {
	dns := "root:123456@tcp(localhost:3306)/notebook?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
		fmt.Println(err.Error())
	}
}


func TestAddUser(t *testing.T) {
	user := &model.User{}
	user.Email = "panda0@126.com"
	user.Password = "123456"
	user.CreateTime = time.Now()
	user.Status = 1
	user.UpdateTime = time.Now()
	userID, err := AddUser(user)
	if err != nil {
		t.Errorf("insert user failed, err:%v\n", err)
		return
	}
	t.Logf("insert user succ, articleId:%d\n", userID)
}

func TestIsUserExistByEmail(t *testing.T) {
	email := "panda8@126.com"

	isExists, err := IsUserExistByEmail(email)

	if err != nil{
		t.Errorf("TestIsUserExistByEmail  failed, err:%v\n", err)
		return
	}

	t.Logf("TestIsUserExistByEmail isExists:%v\n", isExists)
}