package usecase

import "github.com/MosPolyNavigation/web-back/internal/entity"

type Usecase interface {
	GetPlan(campus, corpus string, floor int) (entity.Plan, error)
}
