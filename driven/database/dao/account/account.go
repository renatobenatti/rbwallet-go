package dao

import (
	dto "github.com/devbenatti/rbwallet-go/dto/account"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

var _ AccountDAO = (*Account)(nil)

type Account struct {
}

func NewAccountDAO() Account {
	return Account{}
}

func (a *Account) Create(ac dto.AccountDTO) {
}

func (a *Account) FindByID(id valueObject.Uuid) *dto.AccountDTO {
	return &dto.AccountDTO{}
}
