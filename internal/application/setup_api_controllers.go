package application

import (
	"net/http"

	"github.com/oking02/signapse-test/internal/core/ports"

	"github.com/oking02/signapse-test/internal/handlers/http/controllers"
)

func setupAPIControllers(orderedNumberFetcher ports.OrderedNumbersFetcher) map[string]http.Handler {
	return map[string]http.Handler{
		"/ordered-numbers": controllers.NewGetOrderedNumbers(orderedNumberFetcher),
	}
}
