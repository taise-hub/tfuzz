package tfuzz

import (
	"os"
	"fmt"
)

func ShowError(msg ...interface{}) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf("%s ", s)
	}
	fmt.Printf(tmp)
}

func CheckErr(err error, msg ...interface{}) {
	if err != nil {
		ShowError(msg)
		os.Exit(0)
	}
}