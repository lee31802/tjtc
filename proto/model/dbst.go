package model

type StudentAccountSt struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	ClassNo  int64  `db:"class_no"`
	Ct       int64  `db:"ct"`
	Password string `db:"password"`
}
