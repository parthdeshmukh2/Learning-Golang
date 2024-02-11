package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("I have pizza")

	reader := bufio.NewReader(os.Stdin);

	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for using", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)
	}




}