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
~27.4GB: Massive cumulative allocation detected via `pprof`.
- `pprof` shows **27.4GB** allocated in 5 seconds.
- This corresponds to ~**1.7 billion elements** ($27.4\text{GB} \div 16\text{B string header}$).
- **Analysis**: With 100,000 appends per inner loop, the outer loop ran ~**17,000 times**.
- **Cause**: Frequent re-allocations and data copying of the backing array due to `append`.