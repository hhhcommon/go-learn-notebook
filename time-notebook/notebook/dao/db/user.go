package db

import (
	"github.com/panda8z/notebook/model"
)

// CREATE TABLE  IF NOT EXISTS `user` (
// 	`user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
// 	`email` varchar(128) NOT NULL COMMENT '邮箱',
// 	`password` varchar(256) NOT NULL COMMENT '密码',
// 	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
// 	`update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
// 	 `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1，会员为2，试用3，过期4',
// 	  PRIMARY KEY ( `user_id` )
//   )ENGINE=InnoDB DEFAULT CHARSET=utf8;

// AddUser 插入用户
func AddUser(user *model.User) (userID int64, err error) {
	if user == nil {
		return
	}

	sqlstr := `insert into 
		user( 
		email, password, 
		create_time, update_time, 
		status) 
		values (?,?,?,?,?)`

	result, err := DB.Exec(sqlstr,
		user.Email,
		user.Password,
		user.CreateTime,
		user.UpdateTime,
		user.Status,
	)

	if err != nil {
		return
	}

	userID, err = result.LastInsertId()
	return

}

// IsUserExistByEmail 使用email判断用户是否存在
func IsUserExistByEmail(email string) (isExists bool, err error) {
	if email == string("") {
		return
	}
	isExists = false

	sqlstr := `select user_id from user where email=?`

	result, err := DB.Exec(sqlstr, email)

	id, err := result.LastInsertId()
	
	if err != nil {
		return 
	}
	if id != 0 {
		isExists = true
	}
	return
}
