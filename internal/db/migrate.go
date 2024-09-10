package db

import (
	"github.com/thiagokufa/loja_virtual/internal/account/account_entity"
	"gorm.io/gorm"
)

func Migrate(postgres *gorm.DB) error {
	err := postgres.AutoMigrate(&account_entity.Account{})
	if err != nil {
		return err
	}
	return nil
}
