package application

import (
	"net/http"

	"github.com/oking02/signapse-test/internal/datasources"
	"github.com/oking02/signapse-test/internal/handlers/http/controllers"
)

func setupAPIControllers(orderedNumberFetcher datasources.OrderedNumbersFetcher) map[string]http.Handler {
	return map[string]http.Handler{
		"/ordered-numbers": controllers.NewGetOrderedNumbers(orderedNumberFetcher),
	}
}
