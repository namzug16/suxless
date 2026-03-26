package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	uipages "github.com/namzug16/suxless/pkg/ui/pages"
)

func KitchenSink(c echo.Context) error {
	return c.HTML(http.StatusOK, uipages.KitchenSink().String())
}
