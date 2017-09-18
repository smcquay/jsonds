# jsonds

This is a sample application of running a backend for a [Grafana simple json
datasource](https://github.com/grafana/simple-json-datasource).

## build

```bash
$ go get github.com/smcquay/jsonds
```

## run

```bash
$ jsonds
```

You can specify a number of options, as listed in the `-help`, for example:

```bash
# emit new events every 10 seconds
$ jsonds -period 10s

# start with 30 existing events
$ jsonds -hist 30
```
