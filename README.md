# Firetail KrakenD Plugin POC

A KrakenD plugin proof of concept using the [firetail-go-lib](https://github.com/FireTail-io/firetail-go-lib)'s http middleware.



## Getting Started

First, build the plugin using the krakend/builder image:

```bash
cd firetail-krakend-plugin
docker run --platform linux/amd64 -it -v "$PWD:/app" -w /app krakend/builder:2.2.1 go build -buildmode=plugin -o firetail-krakend-plugin.so .
```

Next, run KrakenD in docker with the plugin using the [krakend.json](./krakend.json) included in this repo:

```bash
cd ..
docker run --platform linux/amd64 -p 8080:8080 -v $PWD:/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/krakend.json
```

Curling KrakenD's `__health` endpoint should be fine as it's in the included [appspec.yaml](./appspec.yaml):

```bash
curl localhost:8080/__health
```

```json
{"agents":{},"now":"2023-03-24 11:49:27.692165134 +0000 UTC m=+8.586939713","status":"ok"}
```

Curling the `/test/{id}` endpoint defined in the [krakend.json](./krakend.json) should be blocked as it's not included in the [appspec.yaml](./appspec.yaml).

```bash
curl localhost:8080/test/10
```

```json
{"code":404,"title":"the resource \"/test/10\" could not be found","detail":"a path for \"/test/10\" could not be found in your appspec"}
```