package models

import (
	"time"

	"github.com/revel/revel"
)

type OnlineUser struct {
	ID            int64     `json:"id" xorm:"id pk autoincr"`
	AuthAccountID int64     `json:"auth_account_id" xorm:"auth_account_id notnull"`
	Hostaddress   string    `json:"hostaddress,omitempty" xorm:"hostaddress"`
	Macaddress    string    `json:"macaddress,omitempty" xorm:"macaddress varchar(200)"`
	CreatedAt     time.Time `json:"created_at,omitempty" xorm:"created_at created"`
}

func (onlineUser *OnlineUser) TableName() string {
	return "tpt_online_users"
}

func (onlineUser *OnlineUser) Validate(validation *revel.Validation) bool {
	validation.Required(onlineUser.AuthAccountID).Key("onlineUser.AuthAccountID")
	validation.MaxSize(onlineUser.Macaddress, 200).Key("onlineUser.Macaddress")
	return validation.HasErrors()
}

func KeyForOnlineUsers(key string) string {
	switch key {
	case "id":
		return "onlineUser.ID"
	case "auth_account_id":
		return "onlineUser.AuthAccountID"
	case "hostaddress":
		return "onlineUser.Hostaddress"
	case "macaddress":
		return "onlineUser.Macaddress"
	case "created_at":
		return "onlineUser.CreatedAt"
	}
	return key
}
