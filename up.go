package main

import (
    "bytes"
    "errors"
    "fmt"
    "io/ioutil"
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
    d := NewDuration(f)

    // TODO: offload into more flexible output function(s)
    fmt.Printf("%dd⋅%dh⋅%dm", d.Days + d.Weeks * 7, d.Hours, d.Minutes)
}


// Convert a byte slice of ASCII numerals into a float.
func BytesToFloat(s []byte) (float64, error) {
    // Find the decimal point
    n := bytes.IndexByte(s, '.')
    // ...or just process as an int if there isn't one
    if n == -1 {
        i, err := BytesToInt(s)
        return float64(i), err
    }
    // ...and count decimal places
    dp := (len(s) - n) - 1

    // Read the integer section
    i, err := BytesToInt(s[:n])
    if err != nil {
        return 0, err
    }
    f := float64(i)
    // ...and the decimals
    i, err = BytesToInt(s[n+1:])
    if err != nil {
        return 0, err
    }
    f += float64(i) / math.Pow10(dp)

    return f, nil
}


// Convert a byte slice of ASCII numerals into an integer.
func BytesToInt(s []byte) (int, error) {
    var i int
    for _, v := range s {
        if v < '0' || v > '9' {
            return 0, errors.New("NaN")
        }
        i *= 10
        i += int(v - '0')
    }
    return i, nil
}


// TODO: rework this (and related) to work with ints for cleaner calculations?
// Get number of units of u seconds from t.
func ExtractUnit(t float64, u float64) (uint, float64) {
    i := math.Floor(t / u)
    return uint(i), t - i * u
}


// Get the raw uptime seconds count.
// If OS other than linux are to be supported, it would be here.
func GetRawUptime() (float64, error) {
    s, err := ioutil.ReadFile("/proc/uptime")
    if err != nil {
        return 0, err
    }
    f, err := BytesToFloat(s[:bytes.IndexByte(s, ' ')])
    return f, err
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

