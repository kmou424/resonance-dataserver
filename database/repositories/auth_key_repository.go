package repositories

import (
	"github.com/kmou424/resonance-dataserver/database/model"
	"gorm.io/gorm"
)

var AuthKey *AuthKeyRepository

type AuthKeyRepository struct {
	db *gorm.DB
}

func (repo *AuthKeyRepository) Has(uuid string) bool {
	tx := repo.db.Table(model.AuthKeyTable)

	if uuid == "" {
		return false
	}

	var count int64
	tx.Where("uuid = ?", uuid).Count(&count)

	return count > 0
}
