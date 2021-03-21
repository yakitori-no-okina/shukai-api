package domain

import "time"

type NotificationModel struct {
	ID            uint64 `gorm:"primaryKey"`
	UserID        uint64
	UserModel     UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	About         string
	Message       string
	HasRead       bool `gorm: "not null"`
	ApprovalID    *uint64
	RecruitmentID *uint64
	DateTime      time.Time `gorm: "not null"`
}

type Notifications []NotificationModel
