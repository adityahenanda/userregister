package models

import "time"

type ModelInterface interface {
	SetCreatedAt()
}

type BaseModel struct {
	Id        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy int       `json:"createdBy"`
	UpdatedBy int       `json:"updatedBy"`
	Status    string    `json:"status"`
}

type ResponseResult struct {
	Status  int
	Message string
	Result  interface{}
}

func (b *BaseModel) SetCreatedAt() {
	b.CreatedAt = time.Now()
}
