package player

import "github.com/labstack/echo/v4"

func PlayerRoutes(commGroup *echo.Group, h *playerHandler) {
	commGroup.POST("", h.Create())
	commGroup.GET("/:playerID", h.GetByID())
}
