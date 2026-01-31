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

## Result
### Good
~2MB: Uses a fixed-size slice and reuses memory.
- A Go string header is 16 bytes (pointer + length).
- $16 \text{ bytes} \times 100,000 \text{ elements} \approx 1.6 \text{ MB}$.
- Since the size is constant, memory usage never exceeds this baseline.
### Bad
~27GB: Uses append, causing the slice to grow indefinitely for 5 seconds.
- If the loop executes ~17,000 times, the slice reaches 1.7 billion elements.
- $16 \text{ bytes} \times 1.7\text{B elements} \approx 27.2 \text{ GB}$.
- The massive consumption comes from the ever-expanding backing array.