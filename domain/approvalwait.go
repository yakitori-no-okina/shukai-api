package domain

type ApprovalWaitModel struct {
	ID               uint64 `gorm:"primaryKey"`
	UserID           uint64
	UserModel        UserModel `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RecruitmentID    uint64
	RecruitmentModel RecruitmentModel `gorm:"foreignKey:RecruitmentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
