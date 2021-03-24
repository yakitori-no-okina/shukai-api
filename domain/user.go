package domain

type Skill struct {
	Icon       string
	Backend    int
	Frontend   int
	Management int
	Mobile     int
	AI         int
}

type UserForAdding struct {
	Mail     string
	Password string
	Name     string
	Icon     string
	Github   string
	Twitter  string
	Links    []string
	About    string
}

type UserForEditting struct {
	Name    string
	Icon    string
	Github  string
	Twitter string
	Links   []string
}

type UserProfile struct {
	UserForEditting
	Skills []Skill
}

type UserModel struct {
	ID       int    `gorm:"primaryKey"`
	Mail     string `gorm: "not null"`
	Password string `gorm: "not null"`
	Name     string `gorm: "not null"`
	Icon     string `gorm: "not null"`
	Github   string
	Twitter  string
	Links    string
	About    string
}

type UserSkill struct {
	Backend    int
	Frontend   int
	Management int
	Mobile     int
	AI         int
	Skills     string
}

type UserSkillModel struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	UserModel UserModel `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserSkill
}
