package usecase

import "shukai-api/domain"

type NotificationInteractor struct {
	N NotificationRepository
}

func (interactor *NotificationInteractor) GetList(user_id int) (notifications domain.Notifications, err error) {
	n_models, error_for_get := interactor.N.Getlist(user_id)
	if error_for_get != nil {
		return domain.Notifications{}, error_for_get
	}

	var ns domain.Notifications
	for _, n_model := range n_models {
		n := domain.Notification{
			ID:             n_model.ID,
			UserID:         n_model.UserID,
			About:          n_model.About,
			Message:        n_model.Message,
			HasRead:        n_model.HasRead,
			ApprovalWaitID: n_model.ApprovalWaitID,
			RequesterID:    n_model.RequesterID,
			RecruitmentID:  n_model.RecruitmentID,
			DateTime:       n_model.DateTime,
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func (interactor *NotificationInteractor) Read(notification_id int) (err error) {
	error_for_update := interactor.N.Update(notification_id, true)
	if error_for_update != nil {
		return error_for_update
	}
	return nil
}
