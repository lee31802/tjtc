package miscsvc

import (
	"github.com/sirupsen/logrus"
	"tjtc/config"
	"tjtc/proto/model"
)

// select by username
var selectByUsernameSql = "SELECT id, `name`,class_no,ct,password FROM tjtc_account.student_account WHERE id=1 LIMIT 1 "

// queryUserByUsername 根据username查询用户信息
func queryStudentById(studentId int64) (*model.StudentAccountSt, error) {
	var student model.StudentAccountSt
	err := config.DB.Get(&student, selectByUsernameSql)
	if err != nil {
		logrus.Errorf("queryStudentById.DB.Get: err:%v", err)
		return nil, err
	}
	return &student, nil
}
