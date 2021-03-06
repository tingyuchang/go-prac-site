package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go-prac-site/e"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type DBManager struct {
	once sync.Once
	Db *gorm.DB
}

var dbManager *DBManager = &DBManager{}

func (dbm *DBManager) lazyInit() {
	dbm.once.Do(func() {
		dsn := fmt.Sprintf("%v:%v@/%v?charset=utf8mb4&parseTime=True", viper.Get("db.username"), viper.Get("db.password"), viper.Get("db.database"))
		dbm.Db,_ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		dbm.Db.AutoMigrate(&User{})

		//TODO error handling
	})
}

// CompareUserAndPassword compare user & password is match or not, return error
// when user not found  or invalid password.
func CompareUserAndPassword(username, password string) error {
	dbManager.lazyInit()
	var user User
	result := dbManager.Db.Take(&user, "name = ?", username)
	if result.Error != nil {
		return result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return fmt.Errorf(e.GetErrorMsg(e.ERROR_AUTH_INVALID_PASSWORD))
	}
	return nil
}

// GetUserById return models.User when uid is exist, return error for not found
func GetUserById(uid string) (User, error) {
	dbManager.lazyInit()
	var user User
	result := dbManager.Db.Take(&user, "id = ?", uid)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil

}
// GetUserByUserName return models.User when username is exist, return error for not found
func GetUserByUserName(username string) (User, error) {
	dbManager.lazyInit()
	var user User
	result := dbManager.Db.Take(&user, "name = ?", username)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
// CreateUser validate imported user data and store to db, return user with uid
func CreateUser(user User) (User, error)  {
	dbManager.lazyInit()
	if err := user.validate(); err != nil {
		return User{}, err
	}
	if err := user.checkExistUser(); err != nil {
		return User{}, err
	}

	user.RegisterAt = time.Now()
	bpw,_ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(bpw)
	result := dbManager.Db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	fmt.Println("created: ", user)
	return user, nil

}
func UpdateUserLastLoginTime(user User) error {
	dbManager.lazyInit()

	result := dbManager.Db.Model(&User{}).Where("id = ?", user.Id).Update("last_login_at", time.Now())
	if result.Error != nil {
		return result.Error
	}

	return nil

}
