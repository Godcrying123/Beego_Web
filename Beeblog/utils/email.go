package utils

import (
	"fmt"
	"time"
)

func main() {
	var currenttime string
	currenttime = time.Now().Format("2019 Jan 01")
	fmt.Println("The current time" + currenttime)
}
