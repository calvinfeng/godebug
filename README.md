# Go Debug

## Use Prometheus

```
$ curl localhost:1234/metrics
```

## Use GOPS

Install `gops` on your machine, assuming that you have Go installed already.

```
$ go get -u github.com/google/gops
```

```
$ gops pprof-heap <pid>
```

Then generate call graphs as pdf

```
$ pdf
```

Or use mem stats for a quick snapshot

```
$ gops memstats <pid>
```