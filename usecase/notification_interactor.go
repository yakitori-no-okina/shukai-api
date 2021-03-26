package usecase

import "shukai-api/domain"

type NotificationInteractor struct {
	N NotificationRepository
}

func (interactor *NotificationInteractor) GetList(user_id int) (notifications domain.Notifications, err error) {
	n_models, error_for_get := interactor.N.Getlist()
	if error_for_get != nil {
		return domain.Notifications{}, error_for_get
	}
	return n_models, nil
}

func (interactor *NotificationInteractor) Read(notification_id int) (err error) {
	error_for_update := interactor.N.Update(notification_id, true)
	if error_for_update != nil {
		return error_for_update
	}
	return nil
}
