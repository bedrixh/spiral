package main

func main() {
	boolToString := map[bool]string{
		false: "  ",
		true:  "##",
	}
	const size = 30

	for lineN := range size {
		for charN := range size {
			print(boolToString[getSpiralPixel(size, lineN, charN)])

		}
		println()
	}
}

func getSpiralPixel(size, line, character int) bool {
	// zero'th character does not draw anything
	character++

	//removing one block from the spiral so it is spiral
	if line < size/2 && character < size/2 && line%2 != 0 && line == character {
		return false
	}

	// adding one block to spiral, so it is spiral
	if line < size/2 && character < size/2 && line%2 == 0 && line == character {
		return true
	}

	// if we are over the half of the spiral, we mirror it to the first half
	if line > size/2 {
		line = size/2 - (line - size/2) - 1
	}
	// same as for line but on character
	if character > size/2 {
		character = size/2 - (character - size/2)
	}

	// drawing the full lines
	if line%2 == 0 && character > line && character < size-line {
		return true
	}

	// drawing the full columns
	alreadyWrittenFullLines := line/2 + 1
	if character < alreadyWrittenFullLines*2 && character%2 != 0 {
		return true
	}

	return false
}
