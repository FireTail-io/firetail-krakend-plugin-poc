// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
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
	options := firetail.Options{
		DisableRequestValidation:  false,
		DisableResponseValidation: false,
	}
	openapiSpecPath, ok := config["openapi-spec-path"].(string)
	if ok {
		options.OpenapiSpecPath = openapiSpecPath
	}
	logApiKey, ok := config["log-api-key"].(string)
	if ok {
		options.LogApiKey = logApiKey
	}
	logApiUrl, ok := config["log-api-url"].(string)
	if ok {
		options.LogApiUrl = logApiUrl
	}

	// Create firetail middleware
	firetailMiddleware, err := firetail.GetMiddleware(&options)
	if err != nil {
		return nil, err
	}

	// Return the handler wrapped in the firetail middleware
	return firetailMiddleware(h), nil
}
