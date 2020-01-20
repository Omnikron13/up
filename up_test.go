package main

import (
    "fmt"
    "strconv"
    "testing"
)

var (
    int_slice   = []byte("123456789")
    float_slice = []byte("1234567.89")
)


// TESTS //


func TestBytesToFloat(t *testing.T) {
    // TODO: rework these into table based tests
    t.Run("Test on int slice", func(t *testing.T) {
        i, err := BytesToFloat(int_slice)
        if err != nil {
            t.Error("Threw error")
        }
        if i != 123456789 {
            t.Error("Incorrect output")
        }
    })

    t.Run("Test on float slice", func(t *testing.T) {
        f, err := BytesToFloat(float_slice)
        if err != nil {
            t.Error("Threw error")
        }
        if f != 1234567.89 {
            t.Error("Incorrect output")
        }
    })

    t.Run("Test on non numeric (int) slice", func(t *testing.T) {
        _, err := BytesToFloat([]byte("123abc"))
        if err == nil {
            t.Error("Didn't throw error")
        }
    })

    t.Run("Test on non numeric (float) slice 1", func(t *testing.T) {
        _, err := BytesToFloat([]byte("abc.123"))
        if err == nil {
            t.Error("Didn't throw error")
        }
    })

    t.Run("Test on non numeric (float) slice 2", func(t *testing.T) {
        _, err := BytesToFloat([]byte("123.abc"))
        if err == nil {
            t.Error("Didn't throw error")
        }
    })
}


func TestBytesToInt(t *testing.T) {
    t.Run("Test on int slice", func(t *testing.T) {
        i, err := BytesToInt(int_slice)
        if err != nil {
            t.Error("Threw error")
        }
        if i != 123456789 {
            t.Error("Incorrect output")
        }
    })

    t.Run("Test on non numeric slice", func(t *testing.T) {
        _, err := BytesToInt([]byte("123abc"))
        if err == nil {
            t.Error("Didn't throw error")
        }
    })
}


func TestExtractUnit(t *testing.T) {
    tests := []struct {
        input_time float64
        input_unit float64
        output_units uint
        output_time  float64
    } {
        {
            input_time: 1000.1,
            input_unit: 500,
            output_units: 2,
            output_time:  0.10000000000002274, // Floats, amiright?
        },
        {
            input_time: 1000.1,
            input_unit: 5000,
            output_units: 0,
            output_time:  1000.1,
        },
    }

    for _, tc := range tests {
        t.Run(
            fmt.Sprintf("Time %f Unit %f", tc.input_time, tc.input_unit),
            func(t *testing.T) {
                u, f := ExtractUnit(tc.input_time, tc.input_unit)
                if u != tc.output_units {
                    t.Errorf("Incorrect unit output; got %d expected %d", u, tc.output_units)
                }
                if f != tc.output_time {
                    t.Errorf("Incorrect time output; got %f expected %f", f, tc.output_time)
                }
        })
    }
}


// BENCHMARKS //


func BenchmarkBytesToInt(b *testing.B) {
    for n := 0; n < b.N; n++ {
        BytesToInt(int_slice)
    }
}

func BenchmarkBytesToInt_ViaString(b *testing.B) {
    for n := 0; n < b.N; n++ {
        BytesToInt_ViaString(int_slice)
    }
}

func BenchmarkBytesToFloat(b *testing.B) {
    for n := 0; n < b.N; n++ {
        BytesToFloat(float_slice)
    }
}

func BenchmarkBytesToFloat_ViaString(b *testing.B) {
    for n := 0; n < b.N; n++ {
        BytesToFloat_ViaString(float_slice)
    }
}


// BENCHMARK COMPARISONS //


func BytesToInt_ViaString(s []byte) (int, error) {
    i, _ := strconv.Atoi(string(s))
    return i, nil
}

func BytesToFloat_ViaString(s []byte) (float64, error) {
    f, _ := strconv.ParseFloat(string(s), 64)
    return f, nil
}

