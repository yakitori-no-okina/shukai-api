package domain

type Skill struct {
	Icon       string `json:"icon" form:"icon" query"icon"`
	Backend    int    `json:"backend" form:"backend" query"backend"`
	Frontend   int    `json:"frontend" form:"frontend" query"frontend"`
	Management int    `json:"management" form:"management" query"management"`
	Mobile     int    `json:"mobile" form:"mobile" query"mobile"`
	AI         int
}

type UserForAdding struct {
	Mail     string   `json:"mail" form:"mail" query"mail"`
	Password string   `json:"password" form:"password" query"password"`
	Name     string   `json:"name" form:"name" query"name"`
	Icon     string   `json:"icon" form:"icon" query"icon"`
	Github   string   `json:"github" form:"github" query"github"`
	Twitter  string   `json:"twitter" form:"twitter" query"twitter"`
	Links    []string `json:"links" form:"links" query"links"`
	About    string   `json:"about" form:"about" query"about"`
}

type UserForEditting struct {
	Name    string `json:"name" form:"name" query"name"`
	Icon    string `json:"icon" form:"icon" query"icon"`
	Github  string `json:"github" form:"github" query"github"`
	Twitter string `json:"twitter" form:"twitter" query"twitter"`
	Links   string `json:"links" form:"links" query"links"`
	About   string `json:"about" form:"about" query"about"`
}

type UserForGetting struct {
	Name    string   `json:"name" form:"name" query"name"`
	Icon    string   `json:"icon" form:"icon" query"icon"`
	Github  string   `json:"github" form:"github" query"github"`
	Twitter string   `json:"twitter" form:"twitter" query"twitter"`
	Links   []string `json:"links" form:"links" query"links"`
	About   string   `json:"about" form:"about" query"about"`
}

type UserProfile struct {
	UserForGetting
	Skills   []Skill `json:"skills" form:"skill_detail" query"skills"`
	TeamID   *int    `json:"team_id"`
	TeamIcon string  `json:"team_icon"`
}

type UserModel struct {
	ID       int    `gorm:"primaryKey" json:"id" form:"id" query"id"`
	Mail     string `gorm: "not null;unique" json:"mail" form:"mail" query"mail"`
	Password string `gorm: "not null;unique" json:"password" form:"password" query"password"`
	Name     string `gorm: "not null" json:"name" form:"name" query"name"`
	Icon     string `gorm: "not null" json:"icon" form:"icon" query"icon"`
	Github   string `json:"github" form:"github" query"github"`
	Twitter  string `json:"twitter" form:"twitter" query"twitter"`
	Links    string `json:"links" form:"links" query"links"`
	About    string `json:"about" form:"about" query"about"`
}

type UserSkill struct {
	Backend    int `json:"backend" form:"backend" query"backend"`
	Frontend   int `json:"frontend" form:"frontend" query"frontend"`
	Management int `json:"management" form:"management" query"management"`
	Mobile     int `json:"mobile" form:"mobile" query"mobile"`
	AI         int
	Skills     string `json:"skills" form:"skills" query"skills"`
}

type UserSkillModel struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	UserModel UserModel `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserSkill
}
