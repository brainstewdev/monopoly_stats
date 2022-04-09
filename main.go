package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
	AUTOR: brainstewdev

	this software emulates a monopoly game with n player and returns how many times each square has been
	landed on in order to calculate max profit
*/
type player struct {
	// name (optional)
	name string
	// square : the square in which the player is sitting on
	square int
}

var playersArray []player

func movePlayer(squares int, p *player) {
	p.square = (squares + p.square) % 40
}

func main() {
	/* get how many players are requested from the arguments */
	var n int
	var turns int
	var seed int
	var verbose bool
	var jsonOut bool
	flag.IntVar(&turns, "turns", 30, "number of turns")
	flag.IntVar(&n, "n", 4, "number of players")
	flag.BoolVar(&verbose, "v", false, "get output from every roll")
	flag.IntVar(&seed, "s", -1, "seed for random roll generation")
	flag.BoolVar(&jsonOut, "json", false, "get json output instead of 'human readable'")
	flag.Parse()
	// allocates n player in the playersArray array
	playersArray = make([]player, n)
	squaresOccurences := make(map[int]int)
	// initialize random generator for rolls
	if seed != -1 {
		rand.Seed(int64(seed))
	} else {
		rand.Seed(time.Now().UnixNano())
	}
	for k, _ := range playersArray {
		playersArray[k].name = strconv.Itoa(k) // set name as the nth player
		playersArray[k].square = 0             // set player position to square 0 (start)
	}
	// gameplay loop
	// how many turns? -> flag
	for i := 0; i < turns; i++ {
		// execute gameplay loop
		// for each player:
		// roll the dices (value from 2 to 12)
		// move
		for _, v := range playersArray {
			// roll the dices
			// WELP :P if you are lucky and proceed to place  two equals roll then you can roll again
			for double := true; double; {
				// roll one dice at a time
				max := 6
				min := 1
				first := rand.Intn(max-min) + min
				second := rand.Intn(max-min) + min
				if verbose {
					fmt.Println("\tplayer", v.name, "rolled", first+second, "and was in square", v.square)
				}
				// move player
				movePlayer(first+second, &v)
				if verbose {
					fmt.Println("\tplayer", v.name, "now in square", v.square)
				}
				// register new square
				squaresOccurences[v.square]++
				//	if double roll then repeat roll
				if first != second {
					double = false
				} else if verbose {
					fmt.Println("\treroll!!")
				}
			}
		}
	}
	if !jsonOut {
		// print square occurences results
		fmt.Printf("square occurrences:\n")
		for k, v := range squaresOccurences {
			fmt.Printf("- square number %d:\t%d\n", k, v)
		}
	} else {
		file, err := json.MarshalIndent(squaresOccurences, "", " ")
		if err == nil {
			// print in the stdout
			fmt.Print(string(file))
		} else {
			fmt.Println("error occured:", err)
		}
	}

}
