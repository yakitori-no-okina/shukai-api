package action

import (
	"encoding/json"
	"net/http"
	"shukai-api/domain"
	"shukai-api/interfaces/database"
	"shukai-api/interfaces/database/approvalwait"
	"shukai-api/interfaces/database/notification"
	"shukai-api/interfaces/database/recruitment"
	"shukai-api/interfaces/database/user"
	"shukai-api/usecase"
	"time"
)

type RecruitmentAddAction struct {
	Interactor usecase.RecruitmentInteractor
}

func NewRecruitmentAddAction(sqlHandler database.SqlHandler) *RecruitmentAddAction {
	return &RecruitmentAddAction{
		Interactor: usecase.RecruitmentInteractor{
			R: &recruitment.Repository{
				SqlHandler: sqlHandler,
			},
			RU: &recruitment.UsersRepository{
				SqlHandler: sqlHandler,
			},
			RC: &recruitment.ConditionsRepository{
				SqlHandler: sqlHandler,
			},
			US: &user.SkillRepository{
				SqlHandler: sqlHandler,
			},
			AW: &approvalwait.Repository{
				SqlHandler: sqlHandler,
			},
			N: &notification.Repository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (action *RecruitmentAddAction) Add(c Context) error {
	rfa := new(domain.RecruitmentForAdding)
	if err := c.Bind(rfa); err != nil {
		return c.JSON(http.StatusBadRequest, rfa)
	}

	layout := "2006/01/02"
	start_date, error_for_parse_s := time.Parse(layout, rfa.StartDate)
	if error_for_parse_s != nil {
		return c.JSON(http.StatusBadRequest, error_for_parse_s)
	}
	end_date, error_for_parse_e := time.Parse(layout, rfa.EndDate)
	if error_for_parse_e != nil {
		return c.JSON(http.StatusBadRequest, error_for_parse_e)
	}
	var conditions domain.Conditions
	json.Unmarshal([]byte(rfa.Conditions), &conditions)
	r_model := &domain.RecruitmentModel{
		OwnerID:    rfa.OwnerID,
		EventName:  rfa.EventName,
		EventURL:   rfa.EventURL,
		NumOfUsers: rfa.NumOfUsers,
		Title:      rfa.Title,
		Message:    rfa.Message,
		Icon:       rfa.Icon,
		StartDate:  start_date,
		EndDate:    end_date,
		Purpose:    rfa.Purpose,
		Address:    rfa.Address,
	}

	rc_model := &domain.RecruitmentConditionsModel{
		Conditions: conditions,
	}

	error_for_store := action.Interactor.Add(r_model, rc_model)
	if error_for_store != nil {
		return c.JSON(http.StatusMethodNotAllowed, error_for_store)
	}

	return c.JSON(http.StatusCreated, nil)
}
