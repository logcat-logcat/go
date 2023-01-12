package main

import (
	"fmt"

	"./ex_game"
)

func main() {
	user1 := *ex_game.NewUsers("logcat1", "1234")
	user1.Scoreediter(10)
	user2 := *ex_game.NewUsers("logcat2", "1234")
	user2.Scoreediter(20)
	fmt.Println(user1, user2)
}
