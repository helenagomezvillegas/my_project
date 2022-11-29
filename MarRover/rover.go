package main

func Moves(commands []string) []int {

	xPosition := 0
	yPosition := 0

	for i := 0; i < len(commands); i++ {
		if commands[i] == "R" {
			xPosition += 1
		}

	}

	return []int{xPosition, yPosition}
}
