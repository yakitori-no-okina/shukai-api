import (
	"shukai-api/domain"
	"shukai-api/interface/datebase"
)

type UserSkillRepository struct {
	database.SqlHandler
}

func (repo *UserSkillRepository) Update(us domain.UserSkill) (id int, err error) {
	result := repo.Create(&us)
	return &us.ID, result.Error
}

func (repo *UserSkillRepository) Update(id int, us domain.UserSkill) (err error) {
	user := domain.UserSkillModel{ID: id}
	result := repo.Model(&user).Update(&us)
	return result.Error
}

func (repo *UserSkillRepository) Get(id int) (user *domain.UserSkillModel, err error) {
	var user domain.UserModel
	result := repo.First(&user, id)
	return &user, result.Error
}





