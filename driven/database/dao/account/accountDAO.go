package dao

import (
	dto "github.com/devbenatti/rbwallet-go/dto/account"
	"github.com/devbenatti/rbwallet-go/model/valueObject"
)

type AccountDAO interface {
	Create(a dto.AccountDTO)
	FindByID(id valueObject.Uuid) *dto.AccountDTO
}
