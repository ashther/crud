package models

import (
	db "../db"
	"database/sql"
)

// User model
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
}

// GetAll get all users
func (u *User) GetAll() (users []User, err error) {
	// users := make([]User, 0)
	rows, err := db.Con.Query("select username, password, nickname from test")
	if err != nil {
		return
	}
	defer rows.Close()

	// store the result in slice
	for rows.Next() {
		var u User
		rows.Scan(&u.Username, &u.Password, &u.Nickname)
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}

// GetOne get one user
func (u *User) GetOne(uid string) (user User, err error) {
	err = db.Con.QueryRow("select username, password, nickname from test where uid = ?", uid).Scan(
		&user.Username, &user.Password, &user.Nickname,
	)

	switch {
	case err == sql.ErrNoRows:
		user = User{}
	case err != nil:
		return
	}

	return
}

// Create create new user
func (u *User) Create() (uid int64, err error) {
	res, err := db.Con.Exec(
		"insert into test (username, password, nickname) values (?, ?, ?)", u.Username, u.Password, u.Nickname,
	)
	if err != nil {
		return
	}

	uid, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}

// Update update one user
func (u *User) Update(uid string) (user User, err error) {
	err = db.Con.QueryRow("select username, password, nickname from test where uid = ?", uid).Scan(
		&user.Username, &user.Password, &user.Nickname,
	)
	if err != nil {
		return
	}

	if u.Username == "" {
		u.Username = user.Username
	}
	if u.Password == "" {
		u.Password = user.Password
	}
	if u.Nickname == "" {
		u.Nickname = user.Nickname
	}

	_, err = db.Con.Exec(
		"update test set username = ?, password = ?, nickname = ? where uid = ?",
		u.Username, u.Password, u.Nickname, uid,
	)
	if err != nil {
		return
	}

	err = db.Con.QueryRow("select username, password, nickname from test where uid = ?", uid).Scan(
		&user.Username, &user.Password, &user.Nickname,
	)
	if err != nil {
		return
	}

	return
}

// Delete delete one user
func (u *User) Delete(uid string) (affected int64, err error) {
	res, err := db.Con.Exec("delete from test where uid = ?", uid)
	if err != nil {
		return
	}

	affected, err = res.RowsAffected()
	if err != nil {
		return
	}
	return
}
