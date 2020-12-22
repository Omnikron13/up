package main

import (
    "math"
    "bytes"
    "errors"
    "io/ioutil"
)


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


// Convert directly from a slice of ASCII digits (and optional
// decimal point) to a Period of nanoseconds.
func BytesToPeriod(s []byte) (Period, error) {
    // Find the decimal point
    n := bytes.IndexByte(s, '.')
    // ...or just process as an int if there isn't one
    if n == -1 {
        i, err := BytesToInt(s)
        return Period(i * Second), err
    }
    // Read the integer section
    i, err := BytesToInt(s[:n])
    if err != nil {
        return 0, err
    }
    p := Period(i * Second)
    // ...and the decimals
    i, err = BytesToInt(s[n+1:])
    if err != nil {
        return 0, err
    }
    return p + i * Decisecond, err
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


