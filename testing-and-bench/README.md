### run test and bench

```
$ go test
PASS
ok      github.com/ktrysmt/golang-samplecodes/testing-and-bench 0.005s
$ go test -bench .
BenchmarkHitValue       2000000000               0.45 ns/op
BenchmarkHitPointer     1000000000               1.84 ns/op
PASS
ok      github.com/ktrysmt/golang-samplecodes/testing-and-bench 3.032s
```

