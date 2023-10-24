package limespeed

import (
	"log/slog"
)

func Main() int {
	slog.Debug("limespeed", "test", true)

	return 0
}

var myPackageVar string

func init() {
	myPackageVar = "Hello from limespeed"
}

func GetMyPackageVar() string {
	return myPackageVar
}

func SetMyPackageVar(newValue string) {
	myPackageVar = newValue
}
