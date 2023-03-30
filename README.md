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



## Configuration

See the KrakenD docs on [injecting plugins](https://www.krakend.io/docs/extending/injecting-plugins/).

See the [example/krakend.json](./example/krakend.json) for an example configuration of the Firetail KrakenD plugin. The following table describes all of the currently supported configuration fields, all of which are optional:

| Field Name                   | Type   | Example                                                      | Description                                                  |
| ---------------------------- | ------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `logs-api-token`             | String | "PS-XX-XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX-XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX" | Your API token for the Firetail SaaS. If unset, no logs will be sent to Firetail. |
| `logs-api-url`               | String | "https://api.logging.eu-west-1.prod.firetail.app/logs/bulk"  | The URL to which logs will be sent via POST requests         |
| `openapi-spec-path`          | String | "/etc/krakend/appspec.yaml"                                  | The absolute path to your appspec. By default, no appspec will be used |
| `enable-request-validation`  | String | "1", "t", "T", "TRUE", "true", "True", "0", "f", "F", "FALSE", "false", "False" | Whether or not requests should be validated against the provided appspec. This is disabled by default and requires `openapi-spec-path` to be defined |
| `enable-response-validation` | String | "1", "t", "T", "TRUE", "true", "True", "0", "f", "F", "FALSE", "false", "False" | Whether or not requests should be validated against the provided appspec. This is disabled by default and requires `openapi-spec-path` to be defined |

