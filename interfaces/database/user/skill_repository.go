import (
	"shukai-api/domain"
	"shukai-api/interface/datebase"
)

type UserSkillRepository struct {
	database.SqlHandler
}

func (repo *UserSkillRepository) Store(us domain.UserSkillModel) (id int, err error) {
	result := repo.Create(&us)
	return int(us.ID), result.Error
}

func (repo *UserSkillRepository) Update(id int, us domain.UserSkill) (err error) {
	user := domain.UserSkillModel{ID: id}
	result := repo.Model(&user).Updates(&us)
	return result.Error
}

func (repo *UserSkillRepository) Get(id int) (user *domain.UserSkillModel, err error) {
	var u domain.UserModel
	result := repo.First(&u, id)
	return &u, result.Error
}





