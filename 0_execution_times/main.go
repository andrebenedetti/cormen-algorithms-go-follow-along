package main

import (
	"fmt"
	"math"
	"time"
)

type Timeframe struct {
    duration time.Duration
    label string
}


type Fn struct {
    label string
    fn func(n int) float64
}




func main() {
    // Given the timeframes: second, minute, hour, day, month, year, century

    // determine biggest size 'n' of problem that can be solved in each
    // time frame, considering that each run takes f(n) microseconds

    // f(n):
    // lg(n)
    // sqrt(n)
    // n
    // n * lg(n)
    // n^2
    // n^3
    // 2^n
    // n!

    fns := []Fn{
        {"n", func(n int) float64 { return float64(n)}},
        {"lg(n)", func(n int) float64 { return math.Log2(float64(n))}},
        {"sqrt(n)", func(n int) float64 { return math.Sqrt(float64(n))}},
        // and so on
        // ...
    }

    // We want to know:
    // (timeframe_in_microseconds) / f(n) = 1, find n

    // for example, if f(n) = n and timeframe = 1 second
    //  1000000 / n = 1
    //  n = 1000000

    // if timeframe = 1 minute,
    // 1000000 * 60 / n = 1
    // n = 60000000

    // notice that the largest representable Duration is
    // approximately 290 years, according to the golang docs.
    // Thus, a century is still representable like this.
    const fnDuration = time.Microsecond
    timeframes := [7]Timeframe{
        {time.Second,"second"},
        {time.Minute,"minute"},
        {time.Hour,"hour"},
        {time.Hour * 24, "day"},
        {time.Hour * 24 * 30, "month"},
        {time.Hour * 24 * 30 * 12,"year"},
        {time.Hour * 24 * 30 * 12 * 100, "century"},
    }

    for _, tf := range(timeframes) {
        fmt.Printf("Timeframe: %s\n", tf.label)
        // TODO: computeOps()
    }
}



// Probably need a numeric method here
// to generalize for any function
// func computeOps(fn Fn, frame Timeframe, opDuration time.Duration) int {
//   return
// }
