package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pallagel/game/helpers"
)

//Welcome message
const (
	welcomeMessage = `Hello Everyone. Welcome to my sample game 
	design to improve my Golang development skills 
	ARE YOU READY TO PLAY????????`
)

//Main method
func main() {
	numbers := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Welcome to the Game!!!!!")
	fmt.Println("**********************************************")
	fmt.Println(welcomeMessage)

	length, _ := helpers.GetUserInput("Create Numbers to Play the Game : ", reader)

	numbers = helpers.RandomNumber(length)
	fmt.Print("The numbers to play in the game : ", numbers)

}
