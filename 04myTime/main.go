package main

import (
	"fmt"
	"time"
)

func main() {
	createdAt := time.Now()
	formatedDate := createdAt.Format("01-02-2006")

	fmt.Println(formatedDate)
}