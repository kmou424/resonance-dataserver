package repositories

import (
	"github.com/kmou424/resonance-dataserver/database/model"
	"gorm.io/gorm"
	"strings"
)

var AuthKey *AuthKeyRepository

type AuthKeyRepository struct {
	db *gorm.DB
}

func (repo *AuthKeyRepository) FindAll() (authKeys []*model.AuthKey) {
	tx := repo.db.Table(model.AuthKeyTable)
	tx.Find(&authKeys)
	return
}

func (repo *AuthKeyRepository) Create(authKey *model.AuthKey) {
	tx := repo.db.Table(model.AuthKeyTable)
	tx.Create(authKey)
}

func (repo *AuthKeyRepository) FindByUser(user string) *model.AuthKey {
	tx := repo.db.Table(model.AuthKeyTable)

	var count int64
	tx.Where("user = ?", user).Count(&count)

	if count == 0 {
		return nil
	}
	var authKey *model.AuthKey
	tx.Where("user = ?", user).First(&authKey)

	return authKey
}

func (repo *AuthKeyRepository) DeleteByUUID(authKey *model.AuthKey) {
	tx := repo.db.Table(model.AuthKeyTable)
	tx.Where("uuid = ?", authKey.UUID).
		Delete(&model.AuthKey{})
}

func (repo *AuthKeyRepository) UpdateByUser(authKey *model.AuthKey) {
	tx := repo.db.Table(model.AuthKeyTable)
	tx.Where("user = ?", authKey.User).
		Updates(authKey)
}

func (repo *AuthKeyRepository) HasUUID(uuid string) bool {
	tx := repo.db.Table(model.AuthKeyTable)

	if uuid == "" {
		return false
	}

	var count int64
	tx.Where("uuid = ?", uuid).Count(&count)

	return count > 0
}

func (repo *AuthKeyRepository) HasUser(user string) bool {
	tx := repo.db.Table(model.AuthKeyTable)

	if user == "" {
		return false
	}

	var count int64
	tx.Where("user = ?", user).Count(&count)

	return count > 0
}

func (repo *AuthKeyRepository) IsAdmin(uuid string) bool {
	tx := repo.db.Table(model.AuthKeyTable)

	if uuid == "" {
		return false
	}

	authKey := model.AuthKey{}
	tx.Where("uuid = ?", uuid).First(&authKey)

	return strings.Contains(authKey.User, "admin")
}
