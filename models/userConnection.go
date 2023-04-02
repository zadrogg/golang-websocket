package models

import (
	"gorm.io/gorm"
	"time"
)

type UserConnection struct {
	Id             uint      `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Token          string    `gorm:"index" json:"ws_token"`
	SocketId       string    `gorm:"index;default:null"`
	UserIdentifier string    `gorm:"index;not null" json:"user_identifier"`
}

func OnOpen(db *gorm.DB, token string, id string, socket string) {
	var user UserConnection
	db.Where("token = ?", token).First(&user)
	db.Model(&user).Updates(UserConnection{
		UpdatedAt:      time.Time{},
		SocketId:       socket,
		UserIdentifier: id,
	})
}

func OnMessage(db *gorm.DB, id string) string {
	var user UserConnection
	db.Where("user_identifier = ?", id).First(&user)
	return user.SocketId
}

func OnClose(db *gorm.DB, socket string) {
	db.Where("socket_id = ?", socket).Delete(&UserConnection{})
}
