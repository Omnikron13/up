package main

import (
	"fmt"
	"golang.org/x/sys/windows"
)


const (
	DLL_NAME  = "kernel32.dll"
	PROC_NAME = "GetTickCount64"
)


var (
	_Kernel32 = windows.MustLoadDLL(DLL_NAME)
	_GetTickCount64 = _Kernel32.MustFindProc(PROC_NAME)
)


// Get the raw uptime count.
// Windows favours milliseconds, so it needs multiplying by 1,000 to
// standardise against nanoseconds (the smallest units any system
// will be likely use for anything).
func GetRawUptime() (Period, error) {
	tick, err := GetTickCount64()
	return Period(tick * Millisecond), err
}


// Extra raw system call wrapper
func GetTickCount64() (ms uint64, e error) {
	raw, _, status:= _GetTickCount64.Call()
	ms = uint64(raw)
	if ms == 0 {
		e = fmt.Errorf("Zeroday error! 0 mseconds returned, code: %d", status)
	}
	return ms, e
}

