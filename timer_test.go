package metrics_test

import (
	"math"
	"testing"
	"time"

	"github.com/facebookgo/metrics"
)

func TestTimerZero(t *testing.T) {
	tm := metrics.NewTimer()
	if count := tm.Count(); 0 != count {
		t.Errorf("tm.Count(): 0 != %v\n", count)
	}
	if min := tm.Min(); 0 != min {
		t.Errorf("tm.Min(): 0 != %v\n", min)
	}
	if max := tm.Max(); 0 != max {
		t.Errorf("tm.Max(): 0 != %v\n", max)
	}
	if mean := tm.Mean(); 0.0 != mean {
		t.Errorf("tm.Mean(): 0.0 != %v\n", mean)
	}
	if stdDev := tm.StdDev(); 0.0 != stdDev {
		t.Errorf("tm.StdDev(): 0.0 != %v\n", stdDev)
	}
	ps := tm.Percentiles([]float64{0.5, 0.75, 0.99})
	if 0.0 != ps[0] {
		t.Errorf("median: 0.0 != %v\n", ps[0])
	}
	if 0.0 != ps[1] {
		t.Errorf("75th percentile: 0.0 != %v\n", ps[1])
	}
	if 0.0 != ps[2] {
		t.Errorf("99th percentile: 0.0 != %v\n", ps[2])
	}
	if rate1 := tm.Rate1(); 0.0 != rate1 {
		t.Errorf("tm.Rate1(): 0.0 != %v\n", rate1)
	}
	if rate5 := tm.Rate5(); 0.0 != rate5 {
		t.Errorf("tm.Rate5(): 0.0 != %v\n", rate5)
	}
	if rate15 := tm.Rate15(); 0.0 != rate15 {
		t.Errorf("tm.Rate15(): 0.0 != %v\n", rate15)
	}
	if rateMean := tm.RateMean(); 0.0 != rateMean {
		t.Errorf("tm.RateMean(): 0.0 != %v\n", rateMean)
	}
}

func TestTimerExtremes(t *testing.T) {
	tm := metrics.NewTimer()
	tm.Update(math.MaxInt64)
	tm.Update(0)
	if stdDev := tm.StdDev(); 6.521908912666392e18 != stdDev {
		t.Errorf("tm.StdDev(): 6.521908912666392e18 != %v\n", stdDev)
	}
}

func TestTimerStartStop(t *testing.T) {
	tm := metrics.NewTimer()
	func() {
		defer tm.Start().Stop()
		time.Sleep(50e6)
	}()
	if tm.Max() == 0 {
		t.Error("tm.Max() == 0")
	}
}

func TestTimerRate1(t *testing.T) {
	tm := metrics.NewTimer()
	tm.Update(3 * time.Second)
	tm.Tick()
	const expected = 0.2
	if r1 := tm.Rate1(); r1 != expected {
		t.Errorf("tm.Rate1(): %v != %v\n", expected, r1)
	}
}
