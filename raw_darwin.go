package main

import (
	"fmt"
	"runtime"
)

//
func GetRawUptime() (Period, error) {
	//Sysinfo()
	t, _ := syscall.Sysctl("kern.boottime")
	fmt.Println(t)
	var info Sysinfo_t
	unix.Sysinfo(*info)
	fmt.Println(info.uptime)
	return 0, nil
}

