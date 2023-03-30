package main

import (
	"testing"

	firetail "github.com/FireTail-io/firetail-go-lib/middlewares/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractOptionsEmptyConfig(t *testing.T) {
	options, err := extractOptions(map[string]interface{}{})
	require.Nil(t, err)
	assert.Equal(t, *options, firetail.Options{})
}

func TestExtractOptionsExtraneousConfig(t *testing.T) {
	options, err := extractOptions(map[string]interface{}{
		"extraneous-field": "extraneous-value",
	})
	require.Nil(t, err)
	assert.Equal(t, *options, firetail.Options{})
}

func TestExtractValidLogsApiToken(t *testing.T) {
	const TestApiToken = "TEST_API_TOKEN"
	options, err := extractOptions(
		map[string]interface{}{
			"logs-api-token": TestApiToken,
		},
	)
	require.Nil(t, err)
	assert.Equal(t, *options, firetail.Options{
		LogsApiToken: TestApiToken,
	})
}

func TestExtractInvalidLogsApiToken(t *testing.T) {
	options, err := extractOptions(
		map[string]interface{}{
			"logs-api-token": 3.14159,
		},
	)
	assert.Equal(t,
		"logs-api-token must be of type string; got '3.14159' (float64)",
		err.Error(),
	)
	assert.Nil(t, options)
}

func TestExtractValidLogsApiUrl(t *testing.T) {
	const TestApiUrl = "TEST_API_URL"
	options, err := extractOptions(
		map[string]interface{}{
			"logs-api-url": TestApiUrl,
		},
	)
	require.Nil(t, err)
	assert.Equal(t, *options, firetail.Options{
		LogsApiUrl: TestApiUrl,
	})
}

func TestExtractInvalidLogsApiUrl(t *testing.T) {
	options, err := extractOptions(
		map[string]interface{}{
			"logs-api-url": 3.14159,
		},
	)
	assert.Equal(t,
		"logs-api-url must be of type string; got '3.14159' (float64)",
		err.Error(),
	)
	assert.Nil(t, options)
}

func TestExtractValidOpenapiSpecPath(t *testing.T) {
	const TestOpenapiSpecPath = "TEST_OPENAPI_SPEC_PATH"
	options, err := extractOptions(
		map[string]interface{}{
			"openapi-spec-path": TestOpenapiSpecPath,
		},
	)
	require.Nil(t, err)
	assert.Equal(t, *options, firetail.Options{
		OpenapiSpecPath: TestOpenapiSpecPath,
	})
}

func TestExtractInvalidOpenapiSpecPath(t *testing.T) {
	options, err := extractOptions(
		map[string]interface{}{
			"openapi-spec-path": 3.14159,
		},
	)
	assert.Equal(t,
		"openapi-spec-path must be of type string; got '3.14159' (float64)",
		err.Error(),
	)
	assert.Nil(t, options)
}

func TestExtractEnableRequestValidationValidValues(t *testing.T) {
	for _, val := range []bool{true, false} {
		options, err := extractOptions(
			map[string]interface{}{
				"enable-request-validation": val,
			},
		)
		require.Nil(t, err)
		assert.Equal(t, *options, firetail.Options{
			EnableRequestValidation: val,
		})
	}
}

func TestExtractEnableRequestValidationInvalidValues(t *testing.T) {
	options, err := extractOptions(
		map[string]interface{}{
			"enable-request-validation": "true",
		},
	)
	require.Equal(
		t,
		"enable-request-validation must be of type bool",
		err.Error(),
	)
	assert.Nil(t, options)
}

func TestExtractEnableResponseValidationValidValues(t *testing.T) {
	for _, val := range []bool{true, false} {
		options, err := extractOptions(
			map[string]interface{}{
				"enable-response-validation": val,
			},
		)
		require.Nil(t, err)
		assert.Equal(t, *options, firetail.Options{
			EnableResponseValidation: val,
		})
	}
}

func TestExtractEnableResponseValidationInvalidValue(t *testing.T) {
	options, err := extractOptions(
		map[string]interface{}{
			"enable-response-validation": "true",
		},
	)
	require.Equal(
		t,
		"enable-response-validation must be of type bool",
		err.Error(),
	)
	assert.Nil(t, options)
}

func TestExtractDebugErrsValidValues(t *testing.T) {
	for _, val := range []bool{true, false} {
		options, err := extractOptions(
			map[string]interface{}{
				"debug-errs": val,
			},
		)
		require.Nil(t, err)
		assert.Equal(t, *options, firetail.Options{
			DebugErrs: val,
		})
	}
}

func TestExtractDebugErrsInvalidValue(t *testing.T) {
	options, err := extractOptions(
		map[string]interface{}{
			"debug-errs": "true",
		},
	)
	require.Equal(
		t,
		"debug-errs must be of type bool",
		err.Error(),
	)
	assert.Nil(t, options)
}
