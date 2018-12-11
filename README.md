# golang-bench-compare

Benchmark various techniques for comparing timestamps in Go.

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
make
./bin/golang-bench-compare
```

### Testing

``make test``

### Benchmarking

``make bench``

https://stackimpact.com/blog/practical-golang-benchmarks/

With Go 1.11.2, Linux, and a Intel® Core™ i7-8550U CPU @ 1.80GHz × 4 with hyperthreading disabled I obtained the following results:

| Benchmark                                   | Iterations  | Time       |
| ------------------------------------------- | ----------- | ---------- |
| BenchmarkMethodBeforepBranched-4            | 10000000000 | 1.69 ns/op |
| BenchmarkMethodBeforeBranched-4             | 10000000000 | 0.56 ns/op |
| BenchmarkStaticBeforeBranched-4             | 10000000000 | 0.56 ns/op |
| BenchmarkStaticBeforeBranchedNoinline-4     | 10000000000 | 4.56 ns/op |
| BenchmarkStaticBeforepBranched-4            | 10000000000 | 1.75 ns/op |
| BenchmarkStaticBeforepBranchedPointerVars-4 | 10000000000 | 1.92 ns/op |
| BenchmarkStaticBeforeMul-4                  | 10000000000 | 3.94 ns/op |
| BenchmarkStaticBeforeMulConst-4             | 10000000000 | 1.12 ns/op |
| BenchmarkStaticBeforeMulConstNoinline-4     | 10000000000 | 3.83 ns/op |
| BenchmarkBeforeRandomInputs-4               | 10000000000 | 1.15 ns/op |
| BenchmarkTimeBefore-4                       | 2000000000  | 12.9 ns/op |
