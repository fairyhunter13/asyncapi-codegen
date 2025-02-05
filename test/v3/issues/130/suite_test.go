// More info: https://github.com/fairyhunter13/asyncapi-codegen/issues/130

package issue130

import (
	"testing"

	testutil "github.com/fairyhunter13/asyncapi-codegen/test"
	"github.com/fairyhunter13/asyncapi-codegen/test/v3/issues/130/decoupling"
	"github.com/fairyhunter13/asyncapi-codegen/test/v3/issues/130/parameters"
	"github.com/fairyhunter13/asyncapi-codegen/test/v3/issues/130/requestreply"
	"github.com/fairyhunter13/asyncapi-codegen/test/v3/issues/130/trait"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	brokers, cleanup := testutil.BrokerControllers(t)
	defer cleanup()

	for _, b := range brokers {
		suite.Run(t, decoupling.NewSuite(b))
		suite.Run(t, parameters.NewSuite(b))
		suite.Run(t, requestreply.NewSuite(b))
	}
	suite.Run(t, trait.NewSuite())
}
