package account_repository

import (
	"errors"

	"github.com/thiagokufa/loja_virtual/internal/account/account_entity"
	"github.com/thiagokufa/loja_virtual/internal/db"
)

func FindByEmail(email string) (account_entity.Account, error) {

	var account account_entity.Account
	var count int64
	err := db.Postgres.Model(&account_entity.Account{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return account_entity.Account{}, err
	}
	if count == 0 {
		return account_entity.Account{}, errors.New("account not found")
	}
	err = db.Postgres.Where("email = ?", email).First(&account).Error
	if err != nil {
		return account_entity.Account{}, err
	}

	return account, nil
}
