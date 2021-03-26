package usecase

import (
	"golang.org/x/crypto/bcrypt"
)

type LoginInteractor struct {
	U UserRepository
}

func (interactor *LoginInteractor) Authentication(mail string, input_pw string) (id int, err error) {
	u_model, error_for_get := interactor.U.GetWithMail(mail)
	if error_for_get != nil {
		return 0, error_for_get
	}
	error_for_authentication := bcrypt.CompareHashAndPassword([]byte(u_model.Password), []byte(input_pw))
	if error_for_authentication != nil {
		return 0, error_for_authentication
	}

	return u_model.ID, nil
}
