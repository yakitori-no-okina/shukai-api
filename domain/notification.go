package domain

import "time"

type NotificationModel struct {
	ID            int `gorm:"primaryKey"`
	UserID        int
	UserModel     UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	About         string
	Message       string
	HasRead       bool `gorm: "not null"`
	ApprovalID    *int
	RecruitmentID *int
	DateTime      time.Time `gorm: "not null"`
}

type Notifications []NotificationModel
