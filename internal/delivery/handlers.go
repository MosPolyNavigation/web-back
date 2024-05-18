package delivery

import (
	"errors"
	"github.com/MosPolyNavigation/web-back/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v3"
	log "github.com/sirupsen/logrus"
)

const planURI = "/plan"

const campusQueryArg = "id"

type Handlers struct {
	client *fiber.App
	uc     usecase.Usecase
	logger *log.Logger
}

func NewHandlers(uc usecase.Usecase, client *fiber.App, logger *log.Logger) *Handlers {
	return &Handlers{
		client: client,
		uc:     uc,
		logger: logger,
	}
}

// RegisterRoute middleware hm...
func (h *Handlers) RegisterRoute() {
	h.client.Get(planURI, func(c fiber.Ctx) error {
		campus := c.Query("campus")
		if campus == "" {
			c.Status(fiber.StatusBadRequest)
			return errors.New("campus is required")
		}

		corpus := c.Query("corpus")
		if corpus == "" {
			c.Status(fiber.StatusBadRequest)
			return errors.New("corpus is required")
		}

		floorQuery := c.Query("floor")
		if floorQuery == "" {
			c.Status(fiber.StatusBadRequest)
			return errors.New("floor is required")
		}

		floor, err := strconv.Atoi(c.Query("floor"))
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return err
		}

		if err != nil {
			h.logger.Debugf("query arg not found: %v", err)
			c.Status(fiber.StatusBadRequest)
			return err
		}

		plan, err := h.uc.GetPlan(campus, corpus, floor)

		if err != nil {
			h.logger.Debugf("plan not found: %v", err)
			c.Status(fiber.StatusNotFound)
			return err
		}

		c.Status(fiber.StatusOK)
		return c.JSON(plan)
	})
}
