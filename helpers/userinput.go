package helpers

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//GetUserInput - A method to get user input
func GetUserInput(message string, r *bufio.Reader) (number int, err error) {
	fmt.Print("Please create the game board : ")

	message, err = r.ReadString('\n')

	if err != nil {
		log.Fatal("Invalid entry", err)
	}

	number, err = strconv.Atoi(strings.TrimSpace(message))

	if err != nil {
		panic(err)
	}

	return number, err

}

//RandomNumber - Generate random numbers and return a slice
func RandomNumber(length int) []int {
	si := make([]int, length)

	rand.Seed(time.Now().UnixNano())

	for i := 1; i < length; i++ {
		si[i] = rand.Intn(i + 1)
	}
	return UniqueValues(si)
}

//UniqueValues - Remove deuplicate numbers and return unique numbers
func UniqueValues(si []int) []int {
	unique := make([]int, 0, len(si))
	gamenum := make(map[int]bool)

	for _, v := range si {
		if _, ok := gamenum[v]; !ok {
			gamenum[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}
