package usecase

import (
	"errors"
	"shukai-api/domain"
	"time"
)

type ApprovalInteractor struct {
	AW ApprovalWaitRepository
	RU RecruitmentUsersRepository
	R  RecruitmentRepository
	N  NotificationRepository
}

func (interactor *ApprovalInteractor) Decide(id int, should_approval bool) (err error) {
	recruitment_id, user_id, error_for_get := interactor.AW.GetProperties(id)
	if error_for_get != nil {
		return error_for_get
	}

	// TODO approvalwaitとrecruitmentusersを統合し、カラムの更新で済ませられるようにする
	error_for_remove := interactor.AW.Remove(id)
	if error_for_remove != nil {
		return error_for_remove
	}

	jst, _ := time.LoadLocation("Asia/Tokyo")
	datetime := time.Now().In(jst)

	n_model := &domain.NotificationModel{
		UserID:   user_id,
		About:    "リクエストについて",
		Message:  "",
		HasRead:  false,
		DateTime: datetime,
	}

	if should_approval {
		ru_model := &domain.RecruitmentUsersModel{UserID: user_id, RecruitmentID: recruitment_id}
		_, error_for_ru_store := interactor.RU.Store(ru_model)
		if error_for_ru_store != nil {
			return error_for_ru_store
		}

		r_model, error_for_get := interactor.R.Get(recruitment_id)
		if error_for_get != nil {
			return error_for_get
		}
		ru_models, error_for_get_list := interactor.RU.GetList(recruitment_id)
		if error_for_get_list != nil {
			return error_for_get_list
		}

		if r_model.NumOfUsers <= len(ru_models)-1 {
			n_model.Message = "リクエストしていた募集の人数が上限に達しました。他の募集に申し込んでください。"
			//TODO teamテーブルを作り、募集削除で関連の募集待ちデータ消せるようにする
			if error_for_remove_list := interactor.AW.RemoveWithRecruitmentID(recruitment_id); error_for_remove_list != nil {
				return error_for_remove_list
			}
			// TODO 非同期処理で実現できるようにする
			for _, model := range ru_models {
				n_model.UserID = model.UserID
				error_for_n_store := interactor.N.Store(n_model)
				if error_for_n_store != nil {
					return error_for_n_store
				}
			}
			return errors.New("すでに募集人数に達しています。他の承認待ちも全て削除しました。")
		}

		n_model.Message = "あなたが申し込んでいた募集のオーナーがあなたのリクエストを承認しました。以下のmailからコンタクトを図ってください。\n mail:" + r_model.Address
	} else {
		n_model.Message = "あなたが申し込んでいた募集のオーナーがあなたのリクエストを拒否しました。"
	}

	error_for_n_store := interactor.N.Store(n_model)
	if error_for_n_store != nil {
		return error_for_n_store
	}

	return nil
}
