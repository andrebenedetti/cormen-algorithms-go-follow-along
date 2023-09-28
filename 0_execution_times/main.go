package main

import "time"

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


    // We want to know:
    // (timeframe_in_microseconds) / f(n) = 1, find n

    // for example, if f(n) = n and timeframe = 1 second
    //  1000000 / n = 1
    //  n = 1000000

    // if timeframe = 1 minute,
    // 1000000 * 60 / n = 1
    // n = 60000000
}
