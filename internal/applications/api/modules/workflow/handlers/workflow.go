package handlers

import "github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"

type WorkflowService struct {
	db repository.Storage
}

func NewWorkflowService(db repository.Storage) *WorkflowService {
	return &WorkflowService{db: db}
}
