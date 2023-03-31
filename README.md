# Firetail KrakenD Plugin

A KrakenD plugin for Firetail, built on [firetail-go-lib](https://github.com/FireTail-io/firetail-go-lib)'s http middleware.



## Getting Started

The Firetail KrakenD plugin distributable is a single `.so` file. To build it, use the KrakenD builder image with a version that matches the version of the KrakenD runtime image you want to use, for example for version `2.2.1`:

```bash
docker run --platform linux/amd64 -it -v "$PWD:/app" -w /app krakend/builder:2.2.1 go build -buildmode=plugin -o firetail-krakend-plugin.so .
```

You should now have a file named `firetail-krakend-plugin.so`. 

An [`appspec.yaml`](./example/appspec.yaml) and [`krakend.json`](./example/krakend.json) is included in the [`example`](./example) directory to test out the plugin. To get this running, first move the `.so` file into the [`example`](./example) directory and `cd` into it:

```bash
mv firetail-krakend-plugin.so example
cd example
```

You can then run the KrakenD runtime image with the plugin and provided example [`appspec.yaml`](./example/appspec.yaml) and [`krakend.json`](./example/krakend.json):

```bash
docker run --platform linux/amd64 -p 8080:8080 -v $PWD:/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/krakend.json
```

Curling KrakenD's [`/__health`](http://localhost:8080/__health) endpoint should produce a result similar to the following:

```bash
curl localhost:8080/__health
```

```json
{"agents":{},"now":"2023-03-30 09:29:12.17746026 +0000 UTC m=+6.320137296","status":"ok"}
```

The provided [`krakend.json`](./example/krakend.json) has request & response validation and debug errs enabled, and defines an endpoint `/test/{id}` which is not included in the provided [`appspec.yaml`](./example/appspec.yaml). We can see the Firetail KrakenD plugin in action by curling the `/test/{id}` endpoint:

```bash
curl localhost:8080/test/1
```

```json
{"code":404,"title":"the resource \"/test/1\" could not be found","detail":"a path for \"/test/1\" could not be found in your appspec"}
```



## Public Releases

You do not have to build the Firetail KrakenD plugin yourself. Versioned releases of the Firetail KrakenD plugin are available as [release assets from this repository](https://github.com/FireTail-io/firetail-krakend-plugin-poc/releases). The naming convention is:

```bash
firetail-krakend-plugin-${ARCH}-$(subst :,-,${KRAKEND_BUILDER_IMAGE})-${RELEASE_VERSION}.so
```

- `${ARCH}` is either `amd64` or `arm64`.
- `${KRAKEND_BUILDER_IMAGE}` correlates to the name of the KrakenD builder image used to build the plugin, which you will need to match to your KrakenD runtime image version. For example:
  - For runtime image `krakend/krakend:2.2.1`, `KRAKEND_BUILDER_IMAGE` would be `builder-2.2.1`
  - For runtime image `krakend/krakend-ee:2.2`, `KRAKEND_BUILDER_IMAGE` would be `builder-ee-2.2`
- `${RELEASE_VERSION}` is the release version of the Firetail KrakenD plugin.



## Setup With Firetail SaaS

To get the Firetail KrakenD plugin to send logs to the Firetail SaaS, you need to create an API or app token to authenticate with the Firetail SaaS' logs API. This can be done via the Firetail SaaS' web UI at [firetail.app](https://firetail.app/). See the docs for [creating an API token](https://firetail.io/docs/create-an-api-token) or [creating an app token](https://firetail.io/docs/create-app-token).

Once you have a token for the Firetail SaaS' logs API, you will need to add it to your KrakenD configuration file as `logs-api-token` under the plugin's configuration, for example:

```json
{
  "version": 3,
  "plugin": {},
  "endpoints": [],
  "extra_config": {
    "plugin/http-server": {
      "name": ["firetail-krakend-plugin"],
      "firetail-krakend-plugin": {
        "logs-api-token": "YOUR-FIRETAIL-SAAS-API-TOKEN",
      }
    }
  }
}
```

ℹ️ See the [Configuration](#configuration) section for information on the other available config fields.

ℹ️ Logs are sent to the Firetail SaaS in batches, so logs may not appear immediately on the Firetail SaaS' web UI.



## Installation Into Existing KrakenD Instances

The Firetail KrakenD plugin is a HTTP server plugin. See the KrakenD docs on [injecting plugins](https://www.krakend.io/docs/extending/injecting-plugins/) for how to load it into your KrakenD instances, and the following section on [Configuration](#configuration) for how to configure the Firetail KrakenD plugin once you have successfully injected it.



## Configuration

See the [example/krakend.json](./example/krakend.json) for an example configuration of the Firetail KrakenD plugin. The following table describes all of the currently supported configuration fields:

| Field Name                   | Type   | Example                                                      | Optional | Description                                                  |
| ---------------------------- | ------ | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `logs-api-token`             | String | `"PS-XX-XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX-XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"` | Yes      | Your API token for the Firetail SaaS. If unset, no logs will be sent to Firetail. |
| `logs-api-url`               | String | `"https://api.logging.eu-west-1.prod.firetail.app/logs/bulk"` | Yes      | The URL to which logs will be sent via POST requests. Defaults to the Firetail SaaS' bulk logs endpoint. |
| `openapi-spec-path`          | String | `"/etc/krakend/appspec.yaml"`                                | Yes      | The absolute path to your appspec. By default, no appspec will be used. |
| `enable-request-validation`  | Bool   | `true`, `false`                                              | Yes      | Whether or not requests should be validated against the provided appspec. This defaults to `false` and requires `openapi-spec-path` to be defined. |
| `enable-response-validation` | Bool   | `true`, `false`                                              | Yes      | Whether or not requests should be validated against the provided appspec. This defaults to `false` and requires `openapi-spec-path` to be defined. |
| `debug-errs`                 | Bool   | `true`, `false`                                              | Yes      | Whether or not to include more verbose information in the RFC7807 error responses' `details` member, returned when requests or responses are blocked by validation if enabled. Defaults to `false`. |

