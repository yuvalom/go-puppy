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
