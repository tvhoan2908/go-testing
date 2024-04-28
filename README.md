## Getting started
- Run testing
```sh
$ go test -v
```
- Run testing with a number cpu
```sh
$ go test -v -cpu 1
```
- Run benchmark:
```sh
$ go test -run x -bench . -cpu 1
```