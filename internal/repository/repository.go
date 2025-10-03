package repository

import "github.com/cuonglv-smartosc/golang-boiler-template/internal/repository/postgres"

type Storage interface {
	postgres.IUserRepository
}
