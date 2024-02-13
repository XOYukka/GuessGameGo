package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	fmt.Println("I`ve chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// fmt.Println(target)

	var success bool

	for i := 0; i < 10; i++ {
		fmt.Println("You have", 10-i, "guesses left.")

		fmt.Println()
		fmt.Print("Make a guess: ")

		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if i == 9 && target != guess {
			break
		}

		if guess < target {
			fmt.Println("Oops. Your guess is LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess is HIGH")
		} else {
			fmt.Println("Good job! My number was", guess, "!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn`t guess my nubmer. It was", target, ":(")
	}
}
