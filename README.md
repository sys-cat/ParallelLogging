# ParallelLogging
Test Logging when run webserver on goroutine

## Before

plz install glide

## Build

```
> go build main.go
```

or

```
> cd PATH_TO_THIS_REPO/
> go install
```

## Usage

run main binary with flag `-p`

```
> ./main -p=8080
> ./main -p=8090
> ./main -p=9000
```

## Benchmark

```
> ab -n 1000000 -c 100 http://127.0.0.1:8080/
> ab -n 10000 -c 100 http://127.0.0.1:8090/
> ab -n 100 -c 100 http://127.0.0.1:9000/
```

### what use apace-bench-multi-url?

** cant install apache-bench-multi-url to macOS !!!**