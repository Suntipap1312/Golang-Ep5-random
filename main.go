package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// fmt.Println(time.Now().UnixNano())
	// answer := 10
	// rand.Seed(time.Now().UnixNano() / 25)

	// for n := 0; n != guess; {
	// 	n = rand.Intn(11) // จำนวน n ตัวเลขที่ต้องการให้สุ่ม (0<n)
	// 	fmt.Printf("%-2d", n)
	// }

	// for turn := 0; turn < 5; turn++ {
	// 	n := rand.Intn(guess + 1)
	// 	if n == guess {
	// 		fmt.Println("Congrats!!!, You guess right!!")
	// 		return
	// 	} else {
	// 		fmt.Printf("YOU LOST : Your guess is %d\n", n)
	// 	}
	// }

	// var myguess int
	// var err error
	// myguess, err = strconv.Atoi(os.Args[1])
	// rand.Seed(time.Now().UnixNano() / 25)
	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }
	// if len(os.Args) != 2 {
	// 	fmt.Println("Wrong input")
	// 	return
	// }
	// // for i := 0; i < 5; i++ {
	// n := rand.Intn(answer + 1)
	// if myguess == n {
	// 	fmt.Printf("Your guess is : %d\n", myguess)
	// 	fmt.Printf("Congrate : ANSWER IS %d\n", n)
	// 	return
	// } else {
	// 	fmt.Printf(" Wrong your guess is : %d\n", myguess)
	// 	fmt.Printf(" Random answer is : %d\n", n)

	// }

	args := os.Args[1:]
	guess, err := strconv.Atoi(args[0])
	rand.Seed(time.Now().UnixNano() * 2)
	if err != nil {
		fmt.Println("Not a number.", err)
		return
	}

	if len(args) != 1 {
		fmt.Println("Pick a number.")
		return

	}

	if guess < 0 {
		fmt.Println("Plese pick positive number.")
		return
	}

	for turn := 0; turn < 5; turn++ {
		n := rand.Intn(11)
		fmt.Printf("Your %d attempt : %d\n", turn+1, n)
		if n == guess {
			fmt.Printf("You Win!!(your %dth attempt)", turn+1)
			if turn == 0 {
				fmt.Println("!!!First turn win +1 Special Bonus!!!")
			}
			return
		}
	}
	fmt.Println("You Lost, try again")
}
