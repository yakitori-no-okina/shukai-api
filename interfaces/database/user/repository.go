package user

import (
	"shukai-api/domain"
	"shukai-api/interfaces/database"

	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	database.SqlHandler
}

func (repo *Repository) Store(u *domain.UserModel) (id int, err error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return 0, err
	}

	u.Password = string(pw)
	result := repo.Create(u)
	return u.ID, result.Error
}

func (repo *Repository) Update(id int, ue *domain.UserForEditting) (err error) {
	user := &domain.UserModel{ID: id}
	result := repo.Model(user).Updates(domain.UserModel{
		Name:    ue.Name,
		Icon:    ue.Icon,
		Github:  ue.Github,
		Twitter: ue.Twitter,
		Links:   ue.Links,
		About:   ue.About,
	})
	return result.Error
}

func (repo *Repository) Get(id int) (user *domain.UserModel, err error) {
	var u domain.UserModel
	result := repo.First(&u, id)
	return &u, result.Error
}

func (repo *Repository) GetWithMail(mail string) (user *domain.UserModel, err error) {
	var u domain.UserModel
	result := repo.Where("mail = ?", mail).First(&u)
	return &u, result.Error
}
