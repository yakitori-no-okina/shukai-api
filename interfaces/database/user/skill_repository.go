import (
	"shukai-api/domain"
	"shukai-api/interface/datebase"
)

type SkillRepository struct {
	database.SqlHandler
}

func (repo *SkillRepository) Store(us domain.UserSkillModel) (id int, err error) {
	result := repo.Create(&us)
	return us.ID, result.Error
}

func (repo *SkillRepository) Update(id int, us domain.UserSkill) (err error) {
	user := domain.UserSkillModel{ID: id}
	result := repo.Model(&user).Updates(&us)
	return result.Error
}

func (repo *SkillRepository) Get(id int) (user *domain.UserSkillModel, err error) {
	var u domain.UserModel
	result := repo.First(&u, id)
	return &u, result.Error
}





