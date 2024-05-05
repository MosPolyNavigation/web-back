package adapters

import "github.com/MosPolyNavigation/web-back/internal/entity"

type Usecase interface {
	GetPlan(id int) (entity.Plan, error)
}
