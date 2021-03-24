package domain

type ApprovalWaitModel struct {
	ID               int `gorm:"primaryKey"`
	UserID           int
	UserModel        UserModel `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RecruitmentID    int
	RecruitmentModel RecruitmentModel `gorm:"foreignKey:RecruitmentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
