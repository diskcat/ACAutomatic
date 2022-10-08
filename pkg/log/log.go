package log

import "fmt"

func Error(a ...any) {
	fmt.Print("Error: ")
	fmt.Println(a...)
}

func Warn(a ...any) {
	fmt.Print("Warn: ")
	fmt.Println(a...)
}

func Info(a ...any) {
	fmt.Print("Info: ")
	fmt.Println(a...)
}

func Debug(a ...any) {
	fmt.Print("Debug: ")
	fmt.Println(a...)
}
