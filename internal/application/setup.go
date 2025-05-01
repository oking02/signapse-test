package application

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/oking02/signapse-test/internal/core/services"
	rest "github.com/oking02/signapse-test/internal/handlers/http"
)

func Setup() (*App, error) {

	rawNumberFetcher, err := setupRawNumberFetcher()
	if err != nil {
		return nil, fmt.Errorf("failed to setup raw number fetcher: %w", err)
	}

	orderedNumberFetcher := services.NewOrderedNumbers(rawNumberFetcher)

	apiServer := rest.NewServer(
		IntEnvar("HTTP_PORT", 3000),
	)

	apiServer.SetupRoutes(setupAPIControllers(orderedNumberFetcher))

	return &App{
		components: []Component{
			apiServer,
		},
	}, nil
}

// StringEnvar is a helper function for looking up environment variables
// and returning the value if found, or a default value if not.
// will panic if the environment variable is required and not found.
// this is to ensure that the application can't bootstrap without the required
// environment variables.
func StringEnvar(name string, defaultValue ...string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}
	return val
}

// IntEnvar is a helper function for looking up environment variables
// and returning the value if found, or a default value if not.
// will panic if the environment variable is required and not found.
// this is to ensure that the application is not started without the required
// environment variables.
func IntEnvar(name string, defaultValue ...int) int {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(name + " is not a valid int: " + err.Error())
	}
	return intVal
}

func hasDefault(value any) bool {
	if value == nil {
		return false
	}
	return reflect.ValueOf(value).Len() > 0
}
