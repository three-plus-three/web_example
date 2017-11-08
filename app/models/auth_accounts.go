package models

import (
	"time"

	"github.com/revel/revel"
)

type AuthAccount struct {
	ID          int64     `json:"id" xorm:"id pk autoincr"`
	ManagerID   int64     `json:"manager_id,omitempty" xorm:"manager_id"`
	LeaderID    int64     `json:"leader_id,omitempty" xorm:"leader_id"`
	Name        string    `json:"name" xorm:"name varchar(250) unique notnull"`
	Password    string    `json:"password,omitempty" xorm:"password varchar(250)"`
	Email       string    `json:"email,omitempty" xorm:"email"`
	Sex         string    `json:"sex" xorm:"sex notnull"`
	Level       string    `json:"level" xorm:"level notnull"`
	Description string    `json:"description,omitempty" xorm:"description text"`
	Birthday    time.Time `json:"birthday,omitempty" xorm:"birthday"`
	CreatedAt   time.Time `json:"created_at,omitempty" xorm:"created_at created"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" xorm:"updated_at updated"`
}

func (authAccount *AuthAccount) TableName() string {
	return "tpt_auth_accounts"
}

func (authAccount *AuthAccount) Validate(validation *revel.Validation) bool {
	validation.Required(authAccount.Name).Key("authAccount.Name")
	validation.MinSize(authAccount.Name, 2).Key("authAccount.Name")
	validation.MaxSize(authAccount.Name, 250).Key("authAccount.Name")
	validation.MinSize(authAccount.Password, 8).Key("authAccount.Password")
	validation.MaxSize(authAccount.Password, 250).Key("authAccount.Password")
	if "" != authAccount.Email {
		validation.Email(authAccount.Email).Key("authAccount.Email")
	}
	validation.Required(authAccount.Sex).Key("authAccount.Sex")
	validation.Required(authAccount.Level).Key("authAccount.Level")
	validation.MaxSize(authAccount.Description, 2000).Key("authAccount.Description")
	return validation.HasErrors()
}

func KeyForAuthAccounts(key string) string {
	switch key {
	case "id":
		return "authAccount.ID"
	case "manager_id":
		return "authAccount.ManagerID"
	case "leader_id":
		return "authAccount.LeaderID"
	case "name":
		return "authAccount.Name"
	case "password":
		return "authAccount.Password"
	case "email":
		return "authAccount.Email"
	case "sex":
		return "authAccount.Sex"
	case "level":
		return "authAccount.Level"
	case "description":
		return "authAccount.Description"
	case "birthday":
		return "authAccount.Birthday"
	case "created_at":
		return "authAccount.CreatedAt"
	case "updated_at":
		return "authAccount.UpdatedAt"
	}
	return key
}
