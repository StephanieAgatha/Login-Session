package models

import (
	"database/sql"
	"learn-session-login-logout/config"
	"learn-session-login-logout/entities"
)

//usermodels ini adalah sebuah model yang digunakan untuk berinteraksi dengan database

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	return &UserModel{db: conn}
}
func (u UserModel) Where(user *entities.User, fieldName string, fieldValue string) error {
	//do query
	row, err := u.db.Query("select * from users where "+fieldName+" = $1 limit 1", fieldValue)
	if err != nil {
		panic("error row")
	}
	//row next
	for row.Next() {
		row.Scan(&user.Id, &user.FullName, &user.Email, &user.Username, &user.Password)
	}
	return nil
}

func (u UserModel) Create(user entities.User) (int64, error) {
	res, err := u.db.Exec("insert into users (fullname, email, username, password) values ($1, $2, $3, $4)", user.FullName, user.Email, user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	lastinsert, _ := res.LastInsertId()
	return lastinsert, nil
}
