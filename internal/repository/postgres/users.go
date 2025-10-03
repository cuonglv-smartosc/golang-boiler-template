package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository/models"
)

type IUserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
}

func (d *Database) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := d.Gorm.
		WithContext(ctx).
		Where("email = ?", strings.ToLower(email)).
		First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to find  userby email: %w", err)
	}

	return &user, nil
}

func (d *Database) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	if err := d.Gorm.
		WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	return &user, nil
}

func (d *Database) CreateUser(user *models.User) error {
	return nil
}

func (d Database) UpdateUser(user *models.User) error {
	return nil
}
