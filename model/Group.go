package model

import "chatroom/global"

type Group struct {
	global.GVA_MODEL

	UserId    uint `gorm:"userId" json:"userId"`
	GroupName string
	Notice    string
}
