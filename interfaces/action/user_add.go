package action

import (
	"encoding/json"
	"net/http"
	"shukai-api/domain"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/recruitment"
	"shukai-api/interfaces/database/user"
	"shukai-api/usecase"
)

type UserAddAction struct {
	Interactor usecase.UserInteractor
}

func NewUserAddAction(sqlHandler database.SqlHandler) *UserAddAction {
	return &UserAddAction{
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

func (action *UserAddAction) Add(c Context) error {
	user := new(domain.UserForAdding)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	links, _ := json.Marshal(user.Links)
	usermodel := &domain.UserModel{
		Mail:     user.Name,
		Password: user.Password,
		Name:     user.Name,
		Icon:     user.Icon,
		Github:   user.Github,
		Twitter:  user.Twitter,
		Links:    string(links),
		About:    user.About,
	}

	id, error_for_store := action.Interactor.Add(usermodel)
	if error_for_store != nil {
		return c.JSON(http.StatusMethodNotAllowed, error_for_store)
	}

	return_value := map[string]int{"id": id, "token": "svasdn"}
	return c.JSON(http.StatusCreated, return_value)
}
