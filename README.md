# Firetail KrakenD Plugin

A KrakenD plugin for Firetail, built on [firetail-go-lib](https://github.com/FireTail-io/firetail-go-lib)'s http middleware.



## Getting Started

The Firetail KrakenD plugin distributable is a single `.so` file. To build, it, use the KrakenD builder image with a version that matches the version of the KrakenD runtime image you want to use, for example for version `2.2.1`:

```bash
docker run --platform linux/amd64 -it -v "$PWD:/app" -w /app krakend/builder:2.2.1 go build -buildmode=plugin -o firetail-krakend-plugin.so .
```

You should now have a file named `firetail-krakend-plugin.so`. 

An [`appspec.yaml`](./example/appspec.yaml) and [`krakend.json`](./example/krakend.json) is included in the [`example`](./example) directory to test out the plugin. To get this running, first move the `.so` file into the `example` directory and `cd` into it:

```bash
mv firetail-krakend-plugin.so example
cd example
```

You can then run the KrakenD runtime image with the plugin and provided example [`appspec.yaml`](./example/appspec.yaml) and [`krakend.json`](./example/krakend.json):

```bash
docker run --platform linux/amd64 -p 8080:8080 -v $PWD:/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/krakend.json
```

Curling KrakenD's `/__health` endpoint should produce a result similar to the following:

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



## Configuration

See the KrakenD docs on [injecting plugins](https://www.krakend.io/docs/extending/injecting-plugins/).

See the [example/krakend.json](./example/krakend.json) for an example configuration of the Firetail KrakenD plugin. The following table describes all of the currently supported configuration fields, all of which are optional:

| Field Name                   | Type   | Example                                                      | Optional | Description                                                  |
| ---------------------------- | ------ | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `logs-api-token`             | String | `"PS-XX-XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX-XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"` | Yes      | Your API token for the Firetail SaaS. If unset, no logs will be sent to Firetail. |
| `logs-api-url`               | String | `"https://api.logging.eu-west-1.prod.firetail.app/logs/bulk"` | Yes      | The URL to which logs will be sent via POST requests         |
| `openapi-spec-path`          | String | `"/etc/krakend/appspec.yaml"`                                | Yes      | The absolute path to your appspec. By default, no appspec will be used |
| `enable-request-validation`  | Bool   | `true`, `false`                                              | Yes      | Whether or not requests should be validated against the provided appspec. This defaults to `false` and requires `openapi-spec-path` to be defined |
| `enable-response-validation` | Bool   | `true`, `false`                                              | Yes      | Whether or not requests should be validated against the provided appspec. This defaults to `false` and requires `openapi-spec-path` to be defined |
| `debug-errs`                 | Bool   | `true`, `false`                                              | Yes      | Whether or not to include more verbose information in the RFC7807 error responses' `details` member, returned when requests or responses are blocked by validation if enabled. Defaults to `false` |

