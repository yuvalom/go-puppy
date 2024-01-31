package gopuppy

import (
	godog "github.com/yuvalom/go-dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return godog.WhenGrownUp("Woof!")
}

func BigBarks() string {
	return godog.WhenGrownUp("Woof! Woof! Woof!")
}

func From1_1_0() string {
	return godog.WhenGrownUp("I am from version 1.1.0")
}
