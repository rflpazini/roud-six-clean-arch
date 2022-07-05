package player

import (
	"net/http"
	"rflpazini/round-six/config"
	"rflpazini/round-six/internal/model"

	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

type playerHandler struct {
	cfg    *config.Config
	logger *zap.Logger
}

func NewPlayerHandler(cfg *config.Config, log *zap.Logger) *playerHandler {
	return &playerHandler{cfg, log}
}

// Create a new player with current request
func (h *playerHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(model.Player)
		if err := c.Bind(p); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, p)
	}
}

// GetByID returns a player by id
func (h *playerHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("playerID")
		return c.JSON(http.StatusOK, id)
	}

}
