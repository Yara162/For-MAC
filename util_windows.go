package main

import "syscall"

func syscallStr(s string) *uint16 {
	p, _ := syscall.UTF16PtrFromString(s)
	return p
}
