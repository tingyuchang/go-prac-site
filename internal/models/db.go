package models

import (
	"20210703/internal/types"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

var db_username = "root"
var db_password = "bieH2leabo1eekaehai5Ahch5eishahj"
var db_dbname = "cart"

// CompareUserAndPassword compare user & password is match or not, return error
// when user not found  or invalid password.
func CompareUserAndPassword(username, password string) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@/%v", db_username, db_password, db_dbname))
	if err != nil {
		return err
	}
	defer db.Close()

	var result []byte
	rows, err := db.Query("SELECT password FROM user WHERE name=?", username)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return err
		}
	}

	err = bcrypt.CompareHashAndPassword(result, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid Password")
	}

	return nil
}

// GetUserById return models.User when uid is exist, return error for not found
func GetUserById(uid string) (User, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@/%v?parseTime=true", db_username, db_password, db_dbname))
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT name,email,phone,gender,birthday,regist_at,last_login_at FROM user WHERE id=?", uid)

	if err != nil {
		return User{}, err
	}
	var user User
	var username string
	var email string
	var phone string
	var gender string
	var birthday types.NullTime
	var registAt types.NullTime
	var lastLoginAt types.NullTime
	if !rows.Next() {
		return User{}, fmt.Errorf("user not found")
	}
	err = rows.Scan(&username, &email, &phone, &gender, &birthday, &registAt, &lastLoginAt)

	if err != nil {
		return User{}, err
	}

	user = User{
		Id: uid,
		Name: username,
		Email: email,
		Phone: phone,
		Gender: gender,
		Birthday: birthday,
		Regist_at: registAt,
		LastLogin: lastLoginAt,
	}

	return user, nil

}
// GetUserByUserName return models.User when username is exist, return error for not found
func GetUserByUserName(username string) (User, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@/%v", db_username, db_password, db_dbname))
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id FROM user WHERE name=?", username)
	if err != nil {
		return User{}, err
	}

	var uid int
	rows.Next()
	err = rows.Scan(&uid)

	if err != nil {
		return User{}, err
	}

	return GetUserById(strconv.Itoa(uid))
}

// CreateUser validate imported user data and store to db, return user with uid
func CreateUser(user User) (User, error)  {
	if err := user.validate(); err != nil {
		return User{}, err
	}

	if err := user.checkExistUser(); err != nil {
		return User{}, err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@/%v", db_username, db_password, db_dbname))
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT user SET name=?, password=?, email=?, phone=?, gender=?, regist_at=?")
	if err != nil {
		return User{}, err
	}

	bpw,_ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	res, err := stmt.Exec(user.Name, bpw, user.Email, user.Phone, user.Gender, time.Now())
	if err != nil {
		return User{}, err
	}

	id, err := res.LastInsertId()

	return GetUserById(strconv.Itoa(int(id)))

}

func UpdateUserLastLoginTime(user User) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@/%v", db_username, db_password, db_dbname))
	if err != nil {
		return err
	}
	defer db.Close()

	stmt,err := db.Prepare("UPDATE user SET last_login_at=? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(time.Now(), user.Id)

	if err != nil {
		return err
	}

	return nil

}
