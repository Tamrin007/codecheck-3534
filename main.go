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
	// 各入力を数値に変換
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

	// 対象の日付をパースして数値に変換
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
	// その年までの閏月の回数
	leapMonth := stockDaysPerYear * (target.year - 1) / daysInMonth
	// 月の上限
	limitMonth := daysInYear / daysInMonth
	// 閏年判定
	// その年を終えた時点で「ストック日数」が「1年でストックされる日数」を
	// 下回っていたら月が追加されている
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

	// 対象の日付までの日数
	pastDays :=
		(daysInYear-stockDaysPerYear)*(target.year-1) +
			daysInMonth*(target.month-1+leapMonth) +
			target.day

	// 曜日を算出
	dayOfWeek := pastDays % daysInWeek

	// AはASCIIコード65なので、64に曜日の剰余を足す、割り切れる場合は1週間の日数を足す
	// 1 -> A曜日、2 -> B曜日、、、0 -> 最後の曜日
	if dayOfWeek != 0 {
		fmt.Println(string(64 + dayOfWeek))
		return
	}
	fmt.Println(string(64 + daysInWeek))
}
