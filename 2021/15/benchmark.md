# Stack implementation (by keeping the steps in path)
goos: darwin
goarch: amd64
pkg: github.com/alperkose/adventofcode2021/15
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_Sample_ExtendByFactorOf1-12    	   45925	     26603 ns/op	   64926 B/op	     310 allocs/op
Benchmark_Sample_ExtendByFactorOf3-12    	     591	   2033996 ns/op	 6770211 B/op	   14257 allocs/op
Benchmark_Sample_ExtendByFactorOf5-12    	     100	  19123085 ns/op	70600493 B/op	   96879 allocs/op
PASS
coverage: 73.2% of statements
ok  	github.com/alperkose/adventofcode2021/15	4.916s
---
# Stack implementation (without keeping the steps in path)
goos: darwin
goarch: amd64
pkg: github.com/alperkose/adventofcode2021/15
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_Sample_ExtendByFactorOf1-12    	  100237	     11807 ns/op	   18143 B/op	      11 allocs/op
Benchmark_Sample_ExtendByFactorOf3-12    	    2271	    533391 ns/op	 1087621 B/op	      22 allocs/op
Benchmark_Sample_ExtendByFactorOf5-12    	     310	   3875069 ns/op	 9258347 B/op	      31 allocs/op
PASS
coverage: 71.2% of statements
ok  	github.com/alperkose/adventofcode2021/15	4.258s
---
# With PriorityQueue
goos: darwin
goarch: amd64
pkg: github.com/alperkose/adventofcode2021/15
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_Sample_ExtendByFactorOf1-12    	   67036	     17211 ns/op	    6008 B/op	     207 allocs/op
Benchmark_Sample_ExtendByFactorOf3-12    	    6962	    173407 ns/op	   47217 B/op	    1808 allocs/op
Benchmark_Sample_ExtendByFactorOf5-12    	    2449	    478195 ns/op	  124998 B/op	    5008 allocs/op
PASS
coverage: 72.8% of statements
ok  	github.com/alperkose/adventofcode2021/15	3.927s