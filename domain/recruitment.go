package domain

import "time"

type Conditions struct {
	Backend    int `json:"backend" form:"backend" query"backend"`
	Frontend   int `json:"frontend" form:"frontend" query"frontend"`
	Management int `json:"management" form:"management" query"management"`
	Mobile     int `json:"mobile" form:"mobile" query"mobile"`
	AI         int
}

type UserInTeam struct {
	ID   int
	Name string
	Icon string
}

type Recruitment struct {
	ID         int
	Icon       string
	NumOfUsers int
	Users      []UserInTeam
	Title      string
	StartDate  time.Time
	EndDate    time.Time
}

type RecruitmentForAdding struct {
	OwnerID    int    `json:"owner_id" form:"owner_id"`
	EventName  string `json:"event_name" form:"event_name"`
	EventURL   string `json:"event_url" form:"event_url"`
	NumOfUsers int    `json:"num_of_users" form:"num_of_users"`
	Conditions string `json:"conditions" form:"conditions"`
	Title      string `json:"title" form:"title"`
	Message    string `json:"message" form:"message"`
	Icon       string `json:"icon" form:"icon"`
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	Purpose    int    `json:"purpose" form:"purpose"`
	Address    string `json:"address" form:"address"`
}

type RecruitmentDetail struct {
	Icon       string
	EventName  string
	EventURL   string
	NumOfUsers int
	Users      []UserInTeam
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
	Icon       string    `gorm: "not null"`
	StartDate  time.Time `gorm: "not null"`
	EndDate    time.Time `gorm: "not null"`
	Purpose    int       `gorm: "not null"`
	Address    string    `gorm: "not null"`
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
