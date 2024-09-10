package account_repository

import (
	"github.com/thiagokufa/loja_virtual/internal/account/account_entity"
	"github.com/thiagokufa/loja_virtual/internal/db"
)

func Create(account account_entity.Account) error {
	err := db.Postgres.Create(&account).Error
	if err != nil {
		return err
	}
	return nil
}
