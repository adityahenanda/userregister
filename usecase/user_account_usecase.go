package usecase

import (
	"errors"
	"userregister/models"
	"userregister/util"

	"golang.org/x/crypto/bcrypt"
)

type userAccountUsecase struct {
	userAccountRepo models.UserAccountRepository
}

// NewUserAccountUseCase create new NewUserAccountUseCase representation of models.UserAccountUsecase interface
func NewUserAccountUseCase(m models.UserAccountRepository) models.UserAccountUsecase {
	return &userAccountUsecase{
		userAccountRepo: m,
	}
}

func (u *userAccountUsecase) RegisterUserAccount(req *models.UserAccountRequest) (id uint64, err error) {

	//validate request
	errMsg := ValidateRegister(req)
	if errMsg != "" {
		return 0, errors.New(errMsg)
	}

	//create user account
	id, err = u.userAccountRepo.CreateUserAccount(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ValidateRegister(req *models.UserAccountRequest) (errMsg string) {
	const SEPARATOR = ","
	if req.NewPassword != req.ConfirmPassword {
		errMsg = util.ConcatString(SEPARATOR, errMsg, "password unmatched")
	}

	if !util.ValidateEmail(req.Email) {
		errMsg = util.ConcatString(SEPARATOR, errMsg, "invalid email")
	}

	return errMsg

}

func (u *userAccountUsecase) LoginUserAccount(req *models.UserAccountLoginRequest) (data *models.UserAccount, err error) {

	//validate request

	//get user account by email
	data, err = u.userAccountRepo.GetUserAccountByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	//verify credential
	valid := validatePassword(data.Password, []byte(req.Password))
	if !valid {
		return nil, errors.New("invalid password")
	}

	return data, nil
}

func (u *userAccountUsecase) GetUserAccountById(id uint64) (data *models.UserAccount, err error) {

	data, err = u.userAccountRepo.GetUserAccountById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userAccountUsecase) GetAllUserAccount(limit int, page int) (data []*models.UserAccount, total int, err error) {
	data, total, err = u.userAccountRepo.GetAllUserAccount(limit, page)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (u *userAccountUsecase) DeleteUserAccount(id uint64) (err error) {
	err = u.userAccountRepo.DeleteUserAccount(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userAccountUsecase) UpdateUserAccount(id uint64, req *models.UserAccountUpdateRequest) (err error) {

	//get existing
	data, err := u.userAccountRepo.GetUserAccountById(id)
	if err != nil {
		return errors.New("data not found")
	}

	errMsg := ValidateUpdateAccount(data, req)
	if errMsg != "" {
		return errors.New(errMsg)
	}

	err = u.userAccountRepo.UpdateUserAccount(id, req)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUpdateAccount(data *models.UserAccount, req *models.UserAccountUpdateRequest) (errMsg string) {
	const SEPARATOR = ","

	if !util.ValidateEmail(req.Email) {
		errMsg = util.ConcatString(SEPARATOR, errMsg, "invalid email")
	}

	if req.Username == "" {
		errMsg = util.ConcatString(SEPARATOR, errMsg, "username should not empty")
	}

	if !validatePassword(data.Password, []byte(req.OldPassword)) {
		errMsg = util.ConcatString(SEPARATOR, errMsg, "invalid password")
	}

	if req.NewPassword != req.ConfirmPassword {
		errMsg = util.ConcatString(SEPARATOR, errMsg, "password unmatched")
	}

	return errMsg

}

func validatePassword(userpwd string, plainPassword []byte) (valid bool) {
	byteHash := []byte(userpwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false
	}
	return true
}
