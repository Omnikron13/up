package main

import (
    "strconv"
    "testing"
)

var (
    int_slice   = []byte("123456789")
    float_slice = []byte("1234567.89")
)


// TESTS //


func TestBytesToFloat(t *testing.T) {
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

