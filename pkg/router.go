package pkg

import (
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/namzug16/suxless/pkg/handlers"
	files "github.com/namzug16/suxless/public"
)

func RegisterRoutes(e *echo.Echo) error {
	staticFS, err := fs.Sub(files.Static, "static")
	if err != nil {
		return err
	}

	e.GET("/", handlers.Home)
	e.GET("/health", handlers.Health)
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(http.FS(staticFS)))))

	return nil
}
