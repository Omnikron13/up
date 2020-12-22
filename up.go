package main

import (
    "fmt"
    "math"
)


const (
    MINUTE = 60
    HOUR   = MINUTE * 60
    DAY    = HOUR * 24
    WEEK   = DAY * 7
)


type Duration struct {
    Weeks   uint8
    Days    uint8
    Hours   uint8
    Minutes uint8
    Seconds float32
}


func main() {
    // TODO: deal with errors
    f, _ := GetRawUptime()
    d := NewDuration(float64(f))

    // TODO: offload into more flexible output function(s)
    fmt.Printf("%dd⋅%dh⋅%dm", d.Days + d.Weeks * 7, d.Hours, d.Minutes)
}


// TODO: rework this (and related) to work with ints for cleaner calculations?
// Get number of units of u seconds from t.
func ExtractUnit(t float64, u float64) (uint, float64) {
    i := math.Floor(t / u)
    return uint(i), t - i * u
}


// Split a float of seconds up into human units.
func NewDuration(f float64) *Duration {
    weeks,   f := ExtractUnit(f, WEEK)
    days,    f := ExtractUnit(f, DAY)
    hours,   f := ExtractUnit(f, HOUR)
    minutes, f := ExtractUnit(f, MINUTE)
    seconds := math.Round(f * 100) / 100
    return &Duration {
        uint8(weeks),
        uint8(days),
        uint8(hours),
        uint8(minutes),
        float32(seconds),
    }
}

