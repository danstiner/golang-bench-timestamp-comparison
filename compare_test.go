package main

import (
	"math/rand"
	"testing"
	"time"
)

// always store the result to a package level variable
// so the compiler cannot eliminate the Benchmark itself.
var result bool

func BenchmarkMethodBeforepBranched(b *testing.B) {
	t1 := &timestamp{seconds: 0, nanos: 0}
	t2 := &timestamp{seconds: 0, nanos: 1}

	var r bool
	for n := 0; n < b.N; n++ {
		r = t1.beforepBranched(t2)
	}
	result = r
}

func BenchmarkMethodBeforeBranched(b *testing.B) {
	t1 := timestamp{seconds: 0, nanos: 0}
	t2 := timestamp{seconds: 0, nanos: 1}

	var r bool
	for n := 0; n < b.N; n++ {
		r = t1.beforeBranched(t2)
	}
	result = r
}

func BenchmarkStaticBeforeBranched(b *testing.B) {
	t1 := timestamp{seconds: 0, nanos: 0}
	t2 := timestamp{seconds: 0, nanos: 1}

	var r bool
	for n := 0; n < b.N; n++ {
		r = beforeBranched(t1, t2)
	}
	result = r
}

func BenchmarkStaticBeforeBranchedNoinline(b *testing.B) {
	t1 := timestamp{seconds: 0, nanos: 0}
	t2 := timestamp{seconds: 0, nanos: 1}

	var r bool
	for n := 0; n < b.N; n++ {
		r = beforeBranchedNoinline(t1, t2)
	}
	result = r
}

func BenchmarkStaticBeforepBranched(b *testing.B) {
	t1 := timestamp{seconds: 0, nanos: 0}
	t2 := timestamp{seconds: 0, nanos: 1}

	var r bool
	for n := 0; n < b.N; n++ {
		r = beforepBranched(&t1, &t2)
	}
	result = r
}

func BenchmarkStaticBeforepBranchedPointerVars(b *testing.B) {
	t1 := &timestamp{seconds: 0, nanos: 0}
	t2 := &timestamp{seconds: 0, nanos: 1}

	var r bool
	for n := 0; n < b.N; n++ {
		r = beforepBranched(t1, t2)
	}
	result = r
}

func BenchmarkStaticBeforeMul(b *testing.B) {
	t1 := &timestamp{seconds: 0, nanos: 0}
	nanos2 := int64(1)

	var r bool
	for n := 0; n < b.N; n++ {
		r = beforeMul(t1, nanos2)
	}
	result = r
}

func BenchmarkStaticBeforeMulConst(b *testing.B) {
	t1 := timestamp{seconds: 0, nanos: 0}
	nanos2 := int64(1)

	var r bool
	for n := 0; n < b.N; n++ {
		r = beforeMulConst(t1, nanos2)
	}
	result = r
}

func BenchmarkStaticBeforeMulConstNoinline(b *testing.B) {
	t1 := timestamp{seconds: 0, nanos: 0}
	nanos2 := int64(1)

	var r bool
	for n := 0; n < b.N; n++ {
		r = beforeMulConstNoinline(t1, nanos2)
	}
	result = r
}

func BenchmarkBeforeRandomInputs(b *testing.B) {
	t1 := timestamp{seconds: rand.Int63(), nanos: rand.Int31()}
	t2 := timestamp{seconds: rand.Int63(), nanos: rand.Int31()}

	var r bool
	for n := 0; n < b.N; n++ {
		r = beforeBranched(t1, t2)
	}
	result = r
}

func BenchmarkTimeBefore(b *testing.B) {
	t1 := time.Unix(0, 0)
	t2 := time.Unix(0, 1)

	var r bool
	for n := 0; n < b.N; n++ {
		r = t1.Before(t2)
	}
	result = r
}
