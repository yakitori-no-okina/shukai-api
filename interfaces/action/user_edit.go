package action

import (
	"encoding/json"
	"net/http"
	"shukai-api/domain"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/recruitment"
	"shukai-api/interfaces/database/user"
	"shukai-api/usecase"
	"strconv"
)

type UserEditAction struct {
	Interactor usecase.UserInteractor
}

func NewUserEditAction(sqlHandler database.SqlHandler) *UserEditAction {
	return &UserEditAction{
		Interactor: usecase.UserInteractor{
			U: &user.Repository{
				SqlHandler: sqlHandler,
			},
			US: &user.SkillRepository{
				SqlHandler: sqlHandler,
			},
			R: &recruitment.Repository{
				SqlHandler: sqlHandler,
			},
			RU: &recruitment.UsersRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (action *UserEditAction) Put(c Context) error {
	// bodyデータの抽出
	user_profile := new(domain.UserProfile)
	if error_for_body := c.Bind(user_profile); error_for_body != nil {
		return c.JSON(http.StatusBadRequest, error_for_body)
	}
	json_str := c.FormValue("skills")
	if error_for_skills := json.Unmarshal([]byte(json_str), &(user_profile.Skills)); error_for_skills != nil {
		return c.JSON(http.StatusBadRequest, error_for_skills)
	}

	// パラメータの抽出
	user_id, error_for_param := strconv.Atoi(c.Param("user_id"))
	if error_for_param != nil {
		return c.JSON(http.StatusBadRequest, error_for_param)
	}

	// 更新するuserデータの整形
	links, _ := json.Marshal(user_profile.Links)
	user_for_editting := &domain.UserForEditting{
		Name:    user_profile.Name,
		Icon:    user_profile.Icon,
		Github:  user_profile.Github,
		Twitter: user_profile.Twitter,
		Links:   string(links),
		About:   user_profile.About,
	}

	// 更新するuserskillデータの整形
	userskill := &domain.UserSkill{}
	for _, skill := range user_profile.Skills {
		userskill.Backend += skill.Backend
		userskill.Frontend += skill.Frontend
		userskill.Management += skill.Management
		userskill.Mobile += skill.Mobile
		userskill.AI += skill.AI
	}
	skills, _ := json.Marshal(user_profile.Skills)
	userskill.Skills = string(skills)

	// 更新
	error_for_update := action.Interactor.Put(user_id, user_for_editting, userskill)
	if error_for_update != nil {
		return c.JSON(http.StatusNotFound, error_for_update)
	}

	return c.JSON(http.StatusNoContent, nil)
}
