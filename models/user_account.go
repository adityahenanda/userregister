package models

type UserAccount struct {
	BaseModel
	Username string `json:"username"`
	Email    string `gorm:"unique;type:varchar(255)" json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type AllUserAccountResp struct {
	ResponseResult
	TotalData int `json:"totalData"`
}

type UserAccountRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Address         string `json:"address" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type UserAccountUpdateRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Address         string `json:"address"`
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserAccountLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
userregister the user account usecase : handle a business logic
*/
type UserAccountUsecase interface {
	RegisterUserAccount(req *UserAccountRequest) (id uint64, err error)
	LoginUserAccount(req *UserAccountLoginRequest) (data *UserAccount, err error)
	GetUserAccountById(id uint64) (data *UserAccount, err error)
	GetAllUserAccount(limit int, page int) (data []*UserAccount, total int, err error)
	UpdateUserAccount(id uint64, req *UserAccountUpdateRequest) (err error)
	DeleteUserAccount(id uint64) (err error)
}

/*
userregister the user account repository : handle a datastore or communication layer
*/
type UserAccountRepository interface {
	CreateUserAccount(req *UserAccountRequest) (id uint64, err error)
	GetUserAccountByEmail(email string) (data *UserAccount, err error)
	GetUserAccountById(id uint64) (data *UserAccount, err error)
	GetAllUserAccount(limit int, page int) (data []*UserAccount, total int, err error)
	UpdateUserAccount(id uint64, req *UserAccountUpdateRequest) (err error)
	DeleteUserAccount(id uint64) (err error)
}
