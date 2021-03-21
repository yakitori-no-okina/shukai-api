package domain

import "time"

type Conditions struct {
	Backend    int
	Frontend   int
	Management int
	Mobile     int
	AI         int
}

type Member struct {
	ID   uint64
	Name string
	Icon string
}

type Recruitment struct {
	ID         uint64
	OwnerID    uint64
	OwnerIcon  string
	NumOfUsers int
	Member     Member
	Title      string
	Message    string
	StartDate  time.Time
	EndDate    time.Time
}

type RecruitmentForAdding struct {
	OwnerID    uint64
	EventName  string
	EventURL   string
	NumOfUsers int
	Conditions Conditions
	Title      string
	Message    string
	StartDate  time.Time
	EndDate    time.Time
	Purpose    int
}

type RecruitmentDetail struct {
	OwnerID    uint64
	OwnerIcon  string
	EventName  string
	EventURL   string
	NumOfUsers int
	Users      []Member
	Conditions Conditions
	Title      string
	Message    string
	StartDate  time.Time
	EndDate    time.Time
	Purpose    int
}

type RecruitmentModel struct {
	ID         uint64 `gorm:"primaryKey"`
	OwnerID    uint64
	UserModel  UserModel `gorm:"foreignKey:OwnerID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventName  string    `gorm: "not null"`
	EventURL   string    `gorm: "not null"`
	NumOfUsers int       `gorm: "not null"`
	Title      string    `gorm: "not null"`
	Message    string
	StartDate  time.Time `gorm: "not null"`
	EndDate    time.Time `gorm: "not null"`
	purpose    int       `gorm: "not null"`
	address    string    `gorm: "not null"`
}

type RecruitmentConditionsModel struct {
	ID               uint64 `gorm:"primaryKey"`
	RecruitmentID    uint64
	RecruitmentModel RecruitmentModel `gorm:"foreignKey:RecruitmentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Conditions
}

type RecruitmentUsersModel struct {
	ID               uint64 `gorm:"primaryKey"`
	UserID           uint64
	UserModel        UserModel `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RecruitmentID    uint64
	RecruitmentModel RecruitmentModel `gorm:"foreignKey:RecruitmentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
