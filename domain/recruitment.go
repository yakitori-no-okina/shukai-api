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
	ID   int
	Name string
	Icon string
}

type Recruitment struct {
	ID         int
	OwnerID    int
	OwnerIcon  string
	NumOfUsers int
	Member     Member
	Title      string
	Message    string
	StartDate  time.Time
	EndDate    time.Time
}

type RecruitmentForAdding struct {
	OwnerID    int
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
	OwnerID    int
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
	ID         int `gorm:"primaryKey"`
	OwnerID    int
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
	ID               int `gorm:"primaryKey"`
	RecruitmentID    int
	RecruitmentModel RecruitmentModel `gorm:"foreignKey:RecruitmentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Conditions
}

type RecruitmentUsersModel struct {
	ID               int `gorm:"primaryKey"`
	UserID           int
	UserModel        UserModel `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RecruitmentID    int
	RecruitmentModel RecruitmentModel `gorm:"foreignKey:RecruitmentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
