package repository

import (
	"latihanFSE/models/entity"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (repo *UserRepo) CreateUser(user *entity.User) *gorm.DB {
	user.Password, _ = HashPassword(user.Password)

	role := getRoleViewer(repo)
	user.RoleID = role.ID
	result := repo.mysqlConn.Create(&user)
	return result
}

func getRoleViewer(repo *UserRepo) entity.Role {

	role := entity.Role{}
	repo.mysqlConn.Where("title = ?", "viewer").First(&role)
	return role
}
