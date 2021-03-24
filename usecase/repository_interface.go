package usecase

import "shukai-api/domain"

type ApprovalWaitRepository interface {
	Store(a domain.ApprovalWaitModel) (id int, err error)
	Remove(id int) (err error)
	Get(id int) (recruitment_id int, user_id int, err error)
}

type NotificationRepository interface {
	Store(n domain.NotificationModel) (err error)
	Getlist() (notifications *domain.Notifications, err error)
}

type RecruitmentConditionsRepository interface {
	Store(rc domain.RecruitmentConditionsModel) (id int, err error)
	Get(id int) (condition *domain.RecruitmentConditionsModel, err error)
	Getlist() (conditionList *[]domain.RecruitmentConditionsModel, err error)
}

type RecruitmentRepository interface {
	Store(r domain.RecruitmentModel) (id int, err error)
	Remove(id int) (err error)
	Get(id int) (recruitment *domain.RecruitmentModel, err error)
}

type RecruitmentUsersRepository interface {
	Store(ru domain.RecruitmentUsersModel) (id int, err error)
	Get(recruitment_id int) (users *[]domain.RecruitmentUsersModel, err error)
}

type UserRepository interface {
	Store(u domain.UserModel) (id int, err error)
	Update(id int, ue domain.UserForEditting) (err error)
	Get(id int) (user *domain.UserModel, err error)
}

type UserSkillRepository interface {
	Store(us domain.UserSkillModel) (id int, err error)
	Update(id int, us domain.UserSkill) (err error)
	Get(id int) (user *domain.UserSkillModel, err error)
}
