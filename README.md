# Smallest server

A small HTTP server written [Golang](https://golang.org/).
It is mostly used for connectivity testing as well as request debugging.
It listens on port 8080
Whatever endpoint you hit, the server will respond with the following payload:

```
{
    "ok": true,
    "path": "THE PATH USED",
    "query": $YOUR QUERY PARAMETERS"
    "header": "THE HEADERS USED"
}
```

Example:

```
> curl --header 'Content-Type: application/json' localhost:8080/endpoint?a=4
{
  "ok": true,
  "path": "/endpoint",
  "query": {"a": ["4"]},
  "header": {
    "Accept": ["*/*"],
    "Content-Type": ["application/json"],
    "User-Agent": ["curl/7.64.1"]
  }
}

```

## build

    make build

## run

    make run

## run from docker

    # after make build
    docker run -p 9000:8080 smallest-server:latest
