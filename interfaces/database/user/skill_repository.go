package user

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"
)

type SkillRepository struct {
	database.SqlHandler
}

func (repo *SkillRepository) Store(us *domain.UserSkillModel) (id int, err error) {
	result := repo.Create(us)
	return us.ID, result.Error
}

func (repo *SkillRepository) Update(user_id int, us *domain.UserSkill) (err error) {
	userskill := &domain.UserSkillModel{}
	result := repo.Model(userskill).Where("user_id = ?", user_id).Updates(domain.UserSkillModel{
		UserSkill: domain.UserSkill{
			Backend:    us.Backend,
			Frontend:   us.Frontend,
			Management: us.Management,
			Mobile:     us.Mobile,
			AI:         us.AI,
			Skills:     us.Skills,
		},
	})
	return result.Error
}

func (repo *SkillRepository) GetWithUserID(user_id int) (user *domain.UserSkillModel, err error) {
	var us domain.UserSkillModel
	result := repo.First(&us, "user_id = ?", user_id)
	return &us, result.Error
}
