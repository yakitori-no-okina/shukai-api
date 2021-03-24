package usecase

import (
	"log"
	"shukai-api/domain"
	"time"
)

type ApprovalInteractor struct {
	Aw ApprovalWaitRepository
	Ru RecruitmentUsersRepository
	N  NotificationRepository
}

func (interactor *ApprovalInteractor) Decide(id int, should_approval bool) (success bool, err error) {
	recruitment_id, user_id, error_for_get := interactor.Aw.Get(id)
	if error_for_get != nil {
		return false, error_for_get
	}

	error_for_remove := interactor.Aw.Remove(id)
	if error_for_remove != nil {
		return false, error_for_remove
	}

	jst, _ := time.LoadLocation("Asia/Tokyo")
	datetime := time.Now().In(jst)

	n_model := domain.NotificationModel{
		UserID:   user_id,
		About:    "リクエストについて",
		Message:  "",
		HasRead:  false,
		DateTime: datetime,
	}

	if should_approval {
		ru_model := domain.RecruitmentUsersModel{UserID: user_id, RecruitmentID: recruitment_id}
		_, error_for_ru_store := interactor.Ru.Store(ru_model)
		if error_for_ru_store != nil {
			log.Fatal(err)
		}

		n_model.Message = "あなたが申し込んでいた募集のオーナーがあなたのリクエストを承認しました"
	} else {
		n_model.Message = "あなたが申し込んでいた募集のオーナーがあなたのリクエストを拒否しました"
	}

	error_for_n_store := interactor.N.Store(n_model)
	if error_for_n_store != nil {
		log.Fatal(err)
	}

	return true, nil
}
