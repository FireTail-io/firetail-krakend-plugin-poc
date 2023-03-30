// A KrakenD server plugin which uses the Firetail Go Library to forward request and response data to the Firetail SaaS.
// Derived from the example KrakenD http server plugin provided here:
// https://www.krakend.io/docs/extending/http-server-plugins/

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	firetail "github.com/FireTail-io/firetail-go-lib/middlewares/http"
)

const pluginName = "firetail-krakend-plugin"

type registerer string

// the symbol the plugin loader will try to load
var HandlerRegisterer = registerer(pluginName)

func (r registerer) RegisterHandlers(f func(name string, handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error))) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {
	// This plugin requires config, so return an err if not found in the krakend.json's extra_config
	config, ok := extra[pluginName].(map[string]interface{})
	if !ok {
		return h, errors.New("configuration not found")
	}

	// Extract options from config
	options, err := extractOptions(config)
	if err != nil {
		return nil, err
	}

	// Create firetail middleware
	firetailMiddleware, err := firetail.GetMiddleware(options)
	if err != nil {
		return nil, err
	}

	// Return the handler wrapped in the firetail middleware
	return firetailMiddleware(h), nil
}

func extractOptions(config map[string]interface{}) (*firetail.Options, error) {
	// This is ugly. Slightly annoying KrakenD gives us map[string]interface{}

	options := firetail.Options{}

	logsApiToken, hasLogsApiToken := config["logs-api-token"]
	if hasLogsApiToken {
		logsApiTokenString, ok := logsApiToken.(string)
		if !ok {
			return nil, fmt.Errorf(
				"logs-api-token must be of type string; got '%#v' (%T)",
				logsApiToken, logsApiToken,
			)
		}
		options.LogsApiToken = logsApiTokenString
	}

	logsApiUrl, hasLogsApiUrl := config["logs-api-url"]
	if hasLogsApiUrl {
		logsApiUrlString, ok := logsApiUrl.(string)
		if !ok {
			return nil, fmt.Errorf(
				"logs-api-url must be of type string; got '%#v' (%T)",
				logsApiUrl, logsApiUrl,
			)
		}
		options.LogsApiUrl = logsApiUrlString
	}

	openapiSpecPath, hasOpenApiSpecPath := config["openapi-spec-path"]
	if hasOpenApiSpecPath {
		openapiSpecPathString, ok := openapiSpecPath.(string)
		if !ok {
			return nil, fmt.Errorf(
				"openapi-spec-path must be of type string; got '%#v' (%T)",
				openapiSpecPath, openapiSpecPath,
			)
		}
		options.OpenapiSpecPath = openapiSpecPathString
	}

	enableRequestValidation, hasEnableRequestValidation := config["enable-request-validation"]
	if hasEnableRequestValidation {
		enableRequestValidationBool, ok := enableRequestValidation.(bool)
		if !ok {
			return nil, errors.New("enable-request-validation must be of type bool")
		}
		options.EnableRequestValidation = enableRequestValidationBool
	}

	enableResponseValidation, hasEnableResponseValidation := config["enable-response-validation"]
	if hasEnableResponseValidation {
		enableResponseValidationBool, ok := enableResponseValidation.(bool)
		if !ok {
			return nil, errors.New("enable-response-validation must be of type bool")
		}
		options.EnableResponseValidation = enableResponseValidationBool
	}

	debugErrs, hasDebugErrs := config["debug-errs"]
	if hasDebugErrs {
		debugErrsBool, ok := debugErrs.(bool)
		if !ok {
			return nil, errors.New("debug-errs must be of type bool")
		}
		options.DebugErrs = debugErrsBool
	}

	return &options, nil
}
