package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type date struct {
	year  int
	month int
	day   int
}

func main() {
	run(os.Args[1:])
}

func run(args []string) {
	daysInYear, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	daysInMonth, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	daysInWeek, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	input := strings.Split(args[3], "-")
	var target date
	target.year, err = strconv.Atoi(input[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	target.month, err = strconv.Atoi(input[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	target.day, err = strconv.Atoi(input[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	// 1年でストックされる日数
	stockDaysPerYear := daysInYear % daysInMonth
	// その年を終えた時のストック日数
	stockDays := stockDaysPerYear * target.year % daysInMonth
	// 閏月の回数
	leapMonth := stockDaysPerYear * (target.year - 1) / daysInMonth
	// 月の上限
	limitMonth := daysInYear / daysInMonth
	// 閏年判定
	if stockDays < stockDaysPerYear {
		limitMonth++
	}

	// 無効な日付をチェック
	if target.day > daysInMonth {
		fmt.Println("-1")
		return
	}
	// 無効な月をチェック
	if target.month > limitMonth {
		fmt.Println("-1")
		return
	}

	pastDaysInYears := (daysInYear - stockDaysPerYear) * (target.year - 1)
	pastDaysInMonths := daysInMonth * (target.month - 1 + leapMonth)

	pastDays := pastDaysInYears + pastDaysInMonths + target.day
	dayOfWeek := pastDays % daysInWeek

	if dayOfWeek != 0 {
		fmt.Println(string(65 + dayOfWeek - 1))
		return
	}
	fmt.Println(string(65 + dayOfWeek + daysInWeek - 1))
}
