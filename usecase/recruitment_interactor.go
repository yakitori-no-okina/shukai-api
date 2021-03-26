package usecase

import (
	"errors"
	"shukai-api/domain"
	"time"
)

type RecruitmentInteractor struct {
	R  RecruitmentRepository
	RU RecruitmentUsersRepository
	RC RecruitmentConditionsRepository
	US UserSkillRepository
	AW ApprovalWaitRepository
	N  NotificationRepository
}

func (interactor *RecruitmentInteractor) GetList(user_id int) (recruitments []domain.Recruitment, err error) {
	us_model, error_for_get := interactor.US.GetWithUserID(user_id)
	if error_for_get != nil {
		return []domain.Recruitment{}, error_for_get
	}

	us := &domain.UserSkill{
		Backend:    us_model.Backend,
		Frontend:   us_model.Frontend,
		Management: us_model.Management,
		Mobile:     us_model.Mobile,
		AI:         us_model.AI,
	}

	r_models, error_for_get_list := interactor.R.GetListWithSkill(us)
	if error_for_get_list != nil {
		return []domain.Recruitment{}, error_for_get_list
	}

	var rs []domain.Recruitment
	for _, r_model := range r_models {
		users, error_for_get_userlist := interactor.getUsersInTeam(r_model.ID)
		if error_for_get_userlist != nil {
			return []domain.Recruitment{}, error_for_get_userlist
		}

		r := domain.Recruitment{
			ID:         r_model.ID,
			Icon:       r_model.Icon,
			NumOfUsers: r_model.NumOfUsers,
			Users:      users,
			Title:      r_model.Title,
			StartDate:  r_model.StartDate,
			EndDate:    r_model.EndDate,
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (interactor *RecruitmentInteractor) GetDetail(recruitment_id int) (recruitment_detail domain.RecruitmentDetail, err error) {
	rc_model, error_for_get := interactor.RC.GetWithRecruitmentID(recruitment_id)
	if error_for_get != nil {
		return domain.RecruitmentDetail{}, error_for_get
	}
	r_model, error_for_get_r := interactor.R.Get(recruitment_id)
	if error_for_get_r != nil {
		return domain.RecruitmentDetail{}, error_for_get_r
	}

	users, error_for_get_userlist := interactor.getUsersInTeam(recruitment_id)
	if error_for_get_userlist != nil {
		return domain.RecruitmentDetail{}, error_for_get_userlist
	}

	rd := domain.RecruitmentDetail{
		Icon:       r_model.Icon,
		EventName:  r_model.EventName,
		EventURL:   r_model.EventURL,
		NumOfUsers: r_model.NumOfUsers,
		Users:      users,
		Conditions: rc_model.Conditions,
		Title:      r_model.Title,
		Message:    r_model.Message,
		StartDate:  r_model.StartDate,
		EndDate:    r_model.EndDate,
		Purpose:    r_model.Purpose,
	}
	return rd, nil
}

func (interactor *RecruitmentInteractor) Add(r_model *domain.RecruitmentModel, rc_model *domain.RecruitmentConditionsModel) (err error) {
	recruitment_id, error_for_store := interactor.R.Store(r_model)
	if error_for_store != nil {
		return error_for_store
	}

	rc_model.RecruitmentID = recruitment_id
	_, error_for_store_conditions := interactor.RC.Store(rc_model)
	if error_for_store_conditions != nil {
		return error_for_store_conditions
	}

	return nil
}

func (interactor *RecruitmentInteractor) Request(recruitment_id int, user_id int) (err error) {
	r_model, error_for_get := interactor.R.Get(recruitment_id)
	if error_for_get != nil {
		return error_for_get
	}
	ru_models, error_for_get_list := interactor.RU.GetList(recruitment_id)
	if error_for_get_list != nil {
		return error_for_get_list
	}

	if r_model.NumOfUsers >= len(ru_models) {
		return errors.New("既に募集人数の上限に達しています。よって、この募集にはリクエストできません")
	}

	aw_id, error_for_store := interactor.AW.Store(&domain.ApprovalWaitModel{
		RecruitmentID: recruitment_id,
		UserID:        user_id,
	})
	if error_for_store != nil {
		return error_for_store
	}

	jst, _ := time.LoadLocation("Asia/Tokyo")
	datetime := time.Now().In(jst)

	n_model := &domain.NotificationModel{
		UserID:         r_model.OwnerID,
		About:          "リクエストについて",
		Message:        "あなたの募集に対してリクエストが届いています。以下のリンクからご確認ください。",
		HasRead:        false,
		DateTime:       datetime,
		ApprovalWaitID: &aw_id,
		RequesterID:    &user_id,
	}
	error_for_store_n := interactor.N.Store(n_model)
	if error_for_store_n != nil {
		return error_for_store_n
	}
	return nil
}

func (interactor *RecruitmentInteractor) Cancel(recruitment_id int, user_id int) (err error) {
	ap_model, error_for_get := interactor.AW.Get(recruitment_id, user_id)
	if error_for_get != nil {
		return error_for_get
	}

	error_for_remove := interactor.AW.Remove(ap_model.ID)
	if error_for_remove != nil {
		return error_for_remove
	}

	r_model, error_for_get := interactor.R.Get(recruitment_id)
	if error_for_get != nil {
		return error_for_get
	}

	jst, _ := time.LoadLocation("Asia/Tokyo")
	datetime := time.Now().In(jst)

	n_model := &domain.NotificationModel{
		UserID:      r_model.OwnerID,
		About:       "リクエストについて",
		Message:     "あなたの募集に対し届いていたリクエストの1件がキャンセルされました。キャンセルした人は以下のリンクから確認してください。",
		HasRead:     false,
		DateTime:    datetime,
		RequesterID: &user_id,
	}
	error_for_store_n := interactor.N.Store(n_model)
	if error_for_store_n != nil {
		return error_for_store_n
	}

	return nil
}

func (interactor *RecruitmentInteractor) getUsersInTeam(recruitment_id int) (users []domain.UserInTeam, err error) {
	ru_models, error_for_get_userlist := interactor.RU.GetList(recruitment_id)
	if error_for_get_userlist != nil {
		return []domain.UserInTeam{}, error_for_get_userlist
	}
	var uts []domain.UserInTeam
	for _, ru_model := range ru_models {
		u := domain.UserInTeam{
			ID:   ru_model.UserID,
			Icon: ru_model.UserModel.Icon,
			Name: ru_model.UserModel.Name,
		}
		uts = append(uts, u)
	}
	return uts, nil
}
