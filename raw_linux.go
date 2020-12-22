package main

import (
    "bytes"
    "io/ioutil"
)


// Get the raw uptime seconds count.
// If OS other than linux are to be supported, it would be here.
func GetRawUptime() (Period, error) {
    s, err := ioutil.ReadFile("/proc/uptime")
    if err != nil {
        return 0, err
    }
    return BytesToPeriod(s[:bytes.IndexByte(s, ' ')])
}

