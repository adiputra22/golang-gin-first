package Models

import (
	"fmt"
	"log"

	"github.com/adiputra22/golang-gin-first/Config"
)

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	log.Println("get all users")

	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}

	return nil
}

func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
