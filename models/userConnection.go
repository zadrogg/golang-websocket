package models

import (
	"errors"
	log "github.com/sirupsen/logrus"
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

func OnOpen(db *gorm.DB, token string, id string, socket string) error {
	var user UserConnection
	if err := db.Debug().Where("token = ?", token).First(&user).Error; err != nil {
		log.Info("No user found with this token: " + token)
		return errors.New("No user found with this token: " + token)
	}

	db.Model(&user).Updates(UserConnection{
		UpdatedAt:      time.Time{},
		SocketId:       socket,
		UserIdentifier: id,
	})

	return nil
}

func OnMessage(db *gorm.DB, id string) string {
	var user UserConnection
	db.Where("user_identifier = ?", id).First(&user)
	return user.SocketId
}

func OnClose(db *gorm.DB, socket string) {
	db.Where("socket_id = ?", socket).Delete(&UserConnection{})
}
