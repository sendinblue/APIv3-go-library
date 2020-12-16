package test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/sendinblue/APIv3-go-library/client"
)

func getAPIClient(tb testing.TB) *client.SendinBlue {
	tb.Helper()
	if testing.Short() {
		tb.Skip("short")
	}
	sib := client.NewHTTPClient(nil)
	sib.Transport.(*httptransport.Runtime).DefaultAuthentication = httptransport.APIKeyAuth("api-key", "header", getAPIKey(tb))
	return sib
}

const apiKeyEnvVariable = "TEST_API_KEY"

func getAPIKey(tb testing.TB) string {
	tb.Helper()
	apiKey := os.Getenv(apiKeyEnvVariable)
	if apiKey == "" {
		tb.Skipf("environment variable %q is not defined", apiKeyEnvVariable)
	}
	return "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"
}

const (
	testEmail       = "test@sendinblue.com"
	testEmailUnique = "test+%d@sendinblue.com"
)

func getTestEmailUnique() string {
	return fmt.Sprintf(testEmailUnique, rand.Int63())
}

var spewConfig *spew.ConfigState

func init() {
	spewConfig = spew.NewDefaultConfig()
	spewConfig.Indent = "\t"
	spewConfig.DisablePointerAddresses = true
	spewConfig.DisableCapacities = true
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
