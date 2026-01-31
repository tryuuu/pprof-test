# pprof-test

A simple Go project to experiment slice memory allocation strategies using `pprof`.

## How to Run

Run the unoptimized dynamic append version:
```bash
go run main.go -type=bad
```

Run the optimized pre-allocation version:
```bash
go run main.go -type=good
```

Visualize the memory profile in your browser:
```bash
go tool pprof -http=:8080 mem.prof
```
