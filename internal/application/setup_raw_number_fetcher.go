package application

import (
	"fmt"
	"github.com/oking02/signapse-test/internal/core/ports"
	"net/http"

	"github.com/oking02/signapse-test/internal/datasources/http/signapsesolutions"
)

func setupRawNumberFetcher() (ports.RawNumbersFetcher, error) {

	rawNumberSource := StringEnvar("NUMBERS_SOURCE", "signapse")

	if rawNumberSource == "signapse" {
		return signapsesolutions.NewClient(
			http.DefaultClient,
			StringEnvar("SIGNAPSE_BASE_URL", "https://technical-test.api.production.signapsesolutions.com"),
			StringEnvar("SIGNAPSE_TOKEN"),
		), nil
	}

	return nil, fmt.Errorf("invalid number source %s", rawNumberSource)
}
