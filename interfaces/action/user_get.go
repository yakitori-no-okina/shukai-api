package action

import (
	"net/http"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/user"
	"shukai-api/usecase"
	"strconv"
)

type UserGetAction struct {
	Interactor usecase.UserInteractor
}

func NewUserGetAction(sqlHandler database.SqlHandler) *UserGetAction {
	return &UserGetAction{
		Interactor: usecase.UserInteractor{
			U: &user.Repository{
				SqlHandler: sqlHandler,
			},
			US: &user.SkillRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (action *UserGetAction) Get(c Context) error {
	// パラメータの抽出
	user_id, error_for_param := strconv.Atoi(c.Param("user_id"))
	if error_for_param != nil {
		return c.JSON(http.StatusBadRequest, error_for_param)
	}
	user_profile, error_for_get := action.Interactor.Get(user_id)
	if error_for_get != nil {
		return c.JSON(http.StatusNotFound, error_for_get)
	}

	return c.JSON(http.StatusOK, user_profile)
}
