package main

import "math"

var NANOS_IN_SECOND = int64(math.Pow10(9))

type timestamp struct {
	seconds int64
	nanos   int32
}

func (t timestamp) beforeBranched(t2 timestamp) bool {
	return t.seconds < t2.seconds || t.nanos < t2.nanos
}

func (t *timestamp) beforepBranched(t2 *timestamp) bool {
	return t.seconds < t2.seconds || t.nanos < t2.nanos
}

func beforeBranched(t1 timestamp, t2 timestamp) bool {
	return t1.seconds < t2.seconds || t1.nanos < t2.nanos
}

//go:noinline
func beforeBranchedNoinline(t1 timestamp, t2 timestamp) bool {
	return t1.seconds < t2.seconds || t1.nanos < t2.nanos
}

func beforepBranched(t1 *timestamp, t2 *timestamp) bool {
	return t1.seconds < t2.seconds || t1.nanos < t2.nanos
}

func beforeMul(t1 *timestamp, nanos2 int64) bool {
	n1 := t1.seconds*int64(math.Pow10(9)) + int64(t1.nanos)
	return n1 < nanos2
}

func beforeMulConst(t1 timestamp, nanos2 int64) bool {
	n1 := t1.seconds*NANOS_IN_SECOND + int64(t1.nanos)
	return n1 < nanos2
}

//go:noinline
func beforeMulConstNoinline(t1 timestamp, nanos2 int64) bool {
	n1 := t1.seconds*NANOS_IN_SECOND + int64(t1.nanos)
	return n1 < nanos2
}
