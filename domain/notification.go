package domain

import "time"

type Notification struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	About          string    `json:"about"`
	Message        string    `json:"message"`
	HasRead        bool      `json:"has_read"`
	ApprovalWaitID *int      `json:"approval_wait_id"`
	RequesterID    *int      `json:"requester_id"`
	RecruitmentID  *int      `json:"recruitment_id"`
	DateTime       time.Time `json:"date_time"`
}

type NotificationModel struct {
	ID             int `gorm:"primaryKey"`
	UserID         int
	UserModel      UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	About          string
	Message        string
	HasRead        bool `gorm: "not null"`
	ApprovalWaitID *int
	RequesterID    *int
	RecruitmentID  *int
	DateTime       time.Time `gorm: "not null"`
}

type Notifications []Notification
