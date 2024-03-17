package log

import "fmt"

func Error(err error) {
	fmt.Println("Error :", err.Error())
}

func Info(info string) {
	fmt.Println("Info : ", info)
}

func Debug(debug string) {
	fmt.Println("Debug :", debug)
}

func Warning(warning string) {
	fmt.Println("Warning :", warning)
}

func Fatal(fatal error) {
	if fatal != nil {
		fmt.Println(fatal)
	}
}

func Trace(trace string) {
	fmt.Println(trace)
}
