package models

import (
	db "../db"
)

// LoginReq consist fileds to log in
type LoginReq struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// LoginCheck check if the username exists
func LoginCheck(l LoginReq) (resultBool bool, user User, err error) {
	user.Username = l.Username
	resultBool = false
	err = db.Con.QueryRow(
		"select nickname from test where username = ? and password = ?", l.Username, l.Password,
	).Scan(
		&user.Nickname,
	)

	if err == nil {
		resultBool = true
	}
	return
}
