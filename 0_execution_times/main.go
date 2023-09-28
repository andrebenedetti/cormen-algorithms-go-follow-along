package main

import (
	"fmt"
	"math"
	"time"
)

type Timeframe struct {
	duration time.Duration
	label    string
}

type Fn struct {
	label string
	fn    func(n float64) float64
}

// Something is fundamentally wrong with this code, just don't know what yet
func main() {
	fns := []Fn{
		// {"n", func(n float64) float64 { return float64(n)}},
		{"lg(n)", func(n float64) float64 { return math.Log2(float64(n)) }},
		{"sqrt(n)", func(n float64) float64 { return math.Sqrt(float64(n)) }},
		// and so on
		// ...
	}

	// The largest representable Duration is approximately
	// 290 years, according to the golang docs.
	// Thus, one we're good representing a century as a Duration.
	const fnDuration = time.Microsecond
	// timeframes := [7]Timeframe{
	// 	{time.Second, "second"},
	// 	{time.Minute, "minute"},
	// 	{time.Hour, "hour"},
	// 	{time.Hour * 24, "day"},
	// 	{time.Hour * 24 * 30, "month"},
	// 	{time.Hour * 24 * 30 * 12, "year"},
	// 	{time.Hour * 24 * 30 * 12 * 100, "century"},
	// }

	// Turns out we can't compute for a second in the way this was designed
	// as the max number of operations gets too big.
	// We'll compute it for 50 microseconds and extrapolate it for the other timeframes.
	baseTimeframe := Timeframe{label: "ms", duration: 50 * time.Microsecond}
	for _, fn := range fns {
		fmt.Printf("Computing for function %s\n", fn.label)
		guess := 1000.00
		newtonMethodIterations := 6
		for i := 0; i < newtonMethodIterations; i++ {
			guess = IterateNewtonsMethod(fn, guess, time.Microsecond, baseTimeframe)
		}
		fmt.Printf("Max operations: %f\n", guess)
	}
}

func evaluatePoint(fn Fn, n float64, opDuration time.Duration, tf Timeframe) float64 {
	return fn.fn(n)*float64(opDuration.Microseconds()) - float64(tf.duration.Microseconds())
}

// Return a new guess based on a previous guess
// We're using a simplified Newton's method, where
// instead of the derivative of the fn, we use the
func IterateNewtonsMethod(fn Fn, guess float64, opDuration time.Duration, tf Timeframe) float64 {
	// we have g(n) = fn(n) * duration - timeframe = 0

	// rough approximation to the derivative at the point
	// using two nearby points separated by 'step'
	step := guess / 100
	x1 := guess + step
	x0 := guess - step
	dn := (evaluatePoint(fn, x1, opDuration, tf) - evaluatePoint(fn, x0, opDuration, tf)) / (x1 - x0)

	return guess - evaluatePoint(fn, guess, opDuration, tf)/dn

}

// Probably need a numeric method here
// to generalize for any function
// func computeOps(fn Fn, frame Timeframe, opDuration time.Duration) int {
//   return
// }
