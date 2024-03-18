package log

import "fmt"

func Error(err error) {
	fmt.Println(err.Error())
}

func Info(info string) {
	fmt.Println(info)
}

func Debug(debug string) {
	fmt.Println(debug)
}

func Warning(warning string) {
	fmt.Println(warning)
}

func Fatal(fatal error) {
	if fatal != nil {
		fmt.Println(fatal)
	}
}

func Trace(trace string) {
	fmt.Println(trace)
}
