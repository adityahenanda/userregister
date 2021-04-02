package repository

import (
	"errors"
	"userregister/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type userAccountRepository struct {
	DB *gorm.DB
}

// NewUserAccountRepository will create an implementation of models.UserAccountRepository
func NewUserAccountRepository(db *gorm.DB) models.UserAccountRepository {
	return &userAccountRepository{
		DB: db,
	}
}

//create user account
func (m *userAccountRepository) CreateUserAccount(req *models.UserAccountRequest) (id uint64, err error) {

	//bcrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.ConfirmPassword), 5)
	if err != nil {
		return 0, errors.New("Error While Hashing Password, Try Again")
	}
	req.ConfirmPassword = string(hash)

	var userAccount models.UserAccount
	userAccount.Username = req.Username
	userAccount.Address = req.Address
	userAccount.Email = req.Email
	userAccount.Password = req.ConfirmPassword
	res := m.DB.Create(&userAccount)
	if res.Error != nil {
		return 0, err
	}
	return userAccount.Id, nil

}

//get user by email
func (m *userAccountRepository) GetUserAccountByEmail(email string) (data *models.UserAccount, err error) {

	var rows []*models.UserAccount
	err = m.DB.Where("status <> ?", "deleted").Find(&rows, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("data not found")
	}

	return rows[0], nil
}

//get user by id
func (m *userAccountRepository) GetUserAccountById(id uint64) (data *models.UserAccount, err error) {
	var rows []*models.UserAccount
	err = m.DB.Where("status <> ?", "deleted").Find(&rows, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("data not found")
	}

	return rows[0], nil
}

func (m *userAccountRepository) GetAllUserAccount(limit int, page int) (data []*models.UserAccount, total int, err error) {
	//default limit, offset
	offset := 0
	if limit == 0 {
		limit = 100
	}
	if page > 0 {
		offset = (page - 1) * limit
	}

	rows, err := m.DB.Raw(`SELECT * FROM user_accounts where status <> 'deleted' order by id desc LIMIT ? OFFSET ?`, limit, offset).Rows()
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.UserAccount
		err = m.DB.ScanRows(rows, &temp)
		if err != nil {
			return nil, 0, err
		}

		data = append(data, &temp)

	}
	return data, len(data), err
}

//soft delete data
func (m *userAccountRepository) DeleteUserAccount(id uint64) (err error) {
	err = m.DB.Exec(`Update user_accounts set status = 'deleted' where status <> 'deleted' and id = ?`, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *userAccountRepository) UpdateUserAccount(id uint64, req *models.UserAccountUpdateRequest) (err error) {

	data, err := m.GetUserAccountById(id)
	if err != nil {
		return err
	}
	data.Username = req.Username
	data.Address = req.Address
	data.Email = req.Email
	data.Password = req.NewPassword
	err = m.DB.Save(&data).Error
	if err != nil {
		return err
	}
	return nil
}
