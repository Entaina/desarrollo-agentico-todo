package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

// HealthCheck responde 200 si la aplicación y la base de datos están disponibles.
func HealthCheck(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if err := tx.RawQuery("SELECT 1").Exec(); err != nil {
		return c.Render(http.StatusServiceUnavailable, r.JSON(map[string]string{"status": "error"}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]string{"status": "ok"}))
}
