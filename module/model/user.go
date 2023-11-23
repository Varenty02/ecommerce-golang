package usermodel

import (
	"Ecommerce/commons"
	"errors"
	"gorm.io/gorm"
	"time"
)

var EntityName = "USER"

type User struct {
	Id        int        `json:"id" gorm:"column:id"`
	Email     string     `json:"email" gorm:"column:email"`
	Password  string     `json:"-" gorm:"column:password"`
	Salt      string     `json:"-" gorm:"column:salt;"`
	LastName  string     `json:"last_name,omitempty" gorm:"column:last_name;"`
	FirstName string     `json:"first_name,omitempty" gorm:"column:first_name;"`
	Phone     string     `json:"phone,omitempty" gorm:"column:phone;"`
	Role      string     `json:"role,omitempty" gorm:"column:role;"`
	Status    int        `json:"status,omitempty" gorm:"column:status"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
type UserCreate struct {
	Id        int        `json:"id" gorm:"column:id"`
	Email     string     `json:"email" gorm:"column:email"`
	Password  string     `json:"password" gorm:"column:password"`
	Salt      string     `json:"-" gorm:"column:salt;"`
	LastName  string     `json:"last_name" gorm:"column:last_name;"`
	FirstName string     `json:"first_name" gorm:"column:first_name;"`
	Phone     string     `json:"phone" gorm:"column:phone;"`
	Role      string     `json:"-" gorm:"column:role;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"create d_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
type UserUpdate struct {
	Email     *string `json:"email" gorm:"column:email;"`
	Password  *string `json:"password" gorm:"column:password;"`
	Salt      *string `json:"-" gorm:"column:salt;"`
	LastName  *string `json:"last_name" gorm:"column:last_name;"`
	FirstName *string `json:"first_name" gorm:"column:first_name;"`
	Phone     *string `json:"phone" gorm:"column:phone;"`
	Role      *string `json:"role" gorm:"column:role;"`
	Status    *int    `json:"status" gorm:"column:status"`
}
type UserDelete struct {
	Email    string         `json:"email" gorm:"column:email;"`
	DeleteAt gorm.DeletedAt `json:"delete_at" gorm:"column:deleted_at;"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (User) TableName() string       { return "users" }
func (UserUpdate) TableName() string { return "users" }
func (UserCreate) TableName() string { return "users" }
func (UserLogin) TableName() string  { return "users" }
func (UserDelete) TableName() string { return "users" }
func (u *User) GetUserId() int {
	return u.Id
}
func (u *User) GetUserEmail() string {
	return u.Email

}
func (u *User) GetUserRole() string {
	return u.Role
}

var (
	ErrEmailOrPasswordInvalid = commons.NewCustomError(
		errors.New("Email or password invalid"),
		"Email or password invalid",
		"ErrEmailOrPasswordInvalid")
	ErrEmailExisted = commons.NewCustomError(
		errors.New("Email has already existed"),
		"Email has already existed",
		"ErrEmailExisted")
)
