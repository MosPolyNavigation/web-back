package delivery

import (
	"github.com/MosPolyNavigation/web-back/internal/adapters"
	"github.com/MosPolyNavigation/web-back/pkg"

	"github.com/gofiber/fiber/v3"
	log "github.com/sirupsen/logrus"
)

const campusURI = "/campus"

const campusQueryArg = "id"

type Handlers struct {
	client *fiber.App
	uc     adapters.Usecase
	logger *log.Logger
}

func NewHandlers(uc adapters.Usecase, client *fiber.App, logger *log.Logger) *Handlers {
	return &Handlers{
		client: client,
		uc:     uc,
		logger: logger,
	}
}

// RegisterRoute middleware hm...
func (h *Handlers) RegisterRoute() {

	h.client.Get(campusURI, func(c fiber.Ctx) error {
		query := c.Query("id", "0")

		value, err := pkg.PickInt(query)

		if err != nil {
			h.logger.Debugf("query arg not found: %v", err)
			c.Status(fiber.StatusBadRequest)
			return err
		}

		plan, err := h.uc.GetPlan(value)

		if err != nil {
			h.logger.Debugf("plan not found: %v", err)
			c.Status(fiber.StatusNotFound)
			return err
		}

		c.Status(fiber.StatusOK)
		return c.JSON(plan)
	})
}
