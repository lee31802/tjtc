package miscsvc

import (
	"tjtc/proto/model"
)

func loginSvc(userId int64, password string, typ int) (user *model.StudentAccountSt, err error) {
	user, err = queryStudentById(userId)
	return user, err
}
