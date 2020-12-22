//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █       █ █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     Separators are not displayed (second is an even number):
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █        █     █         █   █ █       █ █
//      █    █        ███   █         █   ███       █ █
//      █    █          █   █         █   █ █       █ █
//     ███  ███       ███  ███       ███  ███       ███
//
// ---------------------------------------------------------

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec, nano := now.Hour(), now.Minute(), now.Second(), now.Nanosecond()

		clock := [10]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
			dot,
			digits[nano/1e8],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				// colon blink
				next := clock[index][line]
				if (digit == colon || digit == dot) && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second / 10)
	}
}
