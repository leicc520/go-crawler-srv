package model

import (
	"time"
)

type SysActivationCode struct {
	Id           uint   `gorm:"primaryKey;autoIncrement"`
	Status       bool   `gorm:"index"`
	Code         string `gorm:"uniqueIndex"`
	ActivateTime *time.Time
	CreateTime   time.Time
	ExpireTime   *time.Time
}
