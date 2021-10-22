package main

import (
	"fmt"
)

func main() {
	var (
		drob   int64
		znam   int64
		del    int64
		znamtd int64
		deltd  int64
		//cel     int
		//cel_tec int
	)
	fmt.Scan(&drob)
	for i := 0; i < int(drob); i++ {
		fmt.Scan(&znamtd)
		fmt.Scan(&deltd)
		if znam == 0 {
			znam = znamtd
			del = deltd
		} else {
			znam = znam*deltd + znamtd*del
			del = del * deltd
		}
	}
outer:
	for i := 1; i <= int(znam); i++ {
		if znam%int64(i) == 0 && del%int64(i) == 0 && del != 1 {
			znam = znam / int64(i)
			del = del / int64(i)
			i = 1
		}
		if i > 1000 {
			break outer
		}
	}
	//znam = znam + int64(cel)*del
	fmt.Print(znam, "/", del)
}
