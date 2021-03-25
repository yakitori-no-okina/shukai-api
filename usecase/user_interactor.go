package usecase

import (
	"encoding/json"
	"shukai-api/domain"
)

type UserInteractor struct {
	U  UserRepository
	US UserSkillRepository
}

func (interactor *UserInteractor) Add(u *domain.UserModel) (id int, err error) {
	user_id, error_for_store := interactor.U.Store(u)
	if error_for_store != nil {
		return 0, error_for_store
	}
	_, error_for_skillstore := interactor.US.Store(&domain.UserSkillModel{UserID: user_id})
	if error_for_skillstore != nil {
		return 0, error_for_skillstore
	}

	return user_id, nil
}

func (interactor *UserInteractor) Get(id int) (user_profile domain.UserProfile, err error) {
	user, error_for_get := interactor.U.Get(id)
	if error_for_get != nil {
		return domain.UserProfile{}, error_for_get
	}

	var links []string
	json.Unmarshal([]byte(user.Links), &links)
	userprofile := domain.UserProfile{
		UserForGetting: domain.UserForGetting{
			Name:    user.Name,
			Icon:    user.Icon,
			Github:  user.Github,
			Twitter: user.Twitter,
			Links:   links,
		},
	}

	userskill, error_for_getskill := interactor.US.GetWithUserID(id)
	if error_for_getskill != nil {
		return userprofile, nil
	}

	var skills []domain.Skill
	json.Unmarshal([]byte(userskill.Skills), &skills)
	userprofile.Skills = skills

	return userprofile, nil
}

func (interactor *UserInteractor) Put(id int, ue *domain.UserForEditting, us *domain.UserSkill) (err error) {
	error_for_update := interactor.U.Update(id, ue)
	if error_for_update != nil {
		return error_for_update
	}

	error_for_skillupdate := interactor.US.Update(id, us)
	if error_for_skillupdate != nil {
		return error_for_skillupdate
	}
	return nil
}
