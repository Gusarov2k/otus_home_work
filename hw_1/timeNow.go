// Hello now() Завести Go репозиторий на GitHub, написать программу печатающую
//  текущее время / точное время с использованием библиотеки NTP
package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	timeCurrent, _ := ntp.Time("localhost")
	fmt.Println("Time now:", timeCurrent.Format("09:00:02"))
}
