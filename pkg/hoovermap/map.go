package hoovermap

import "fmt"

var (
	NStep = []int{0, 1}  // NStep is one unit to the North.
	EStep = []int{1, 0}  // EStep is one unit to the East.
	SStep = []int{0, -1} // SStep is one unit to the South.
	WStep = []int{-1, 0} // WStep is one unit to the West.
)

// ComputeMovement will take an input and return a computed output.
func ComputeMovement(inputs *APIInput) (*APIOutput, error) {
	if err := isValid(inputs.Coords, inputs.RoomSize); err != nil {
		return nil, err
	}

	currentPosition := inputs.Coords
	gridSize := inputs.RoomSize
	cleanedPatches := make(map[string]bool)

	// check the initial position for dirt.
	p := formatCoordinates(currentPosition)
	if _, ok := cleanedPatches[p]; !ok {
		for _, val := range inputs.Patches {
			if matchCoordinates(currentPosition, val) {
				cleanedPatches[p] = true
			}
		}
	}

	for _, direction := range inputs.Instructions {
		switch direction {
		case 'N':
			currentPosition = takeStep(currentPosition, NStep, gridSize)
		case 'E':
			currentPosition = takeStep(currentPosition, EStep, gridSize)
		case 'S':
			currentPosition = takeStep(currentPosition, SStep, gridSize)
		case 'W':
			currentPosition = takeStep(currentPosition, WStep, gridSize)
		default:
			return nil, fmt.Errorf("unrecognised movement directive")
		}

		t := formatCoordinates(currentPosition)
		// Check if we have already cleaned this patch, if we
		// have then we don't need to process any further.
		if _, ok := cleanedPatches[t]; ok {
			continue
		} else {
			for _, val := range inputs.Patches {
				if matchCoordinates(currentPosition, val) {
					cleanedPatches[t] = true
				}
			}
		}
	}

	return &APIOutput{Coords: currentPosition, Patches: len(cleanedPatches)}, nil
}

// isValid performs several sanity checks.
func isValid(currentPosition []int, gridSize []int) error {
	// Check that the position is inside the room.
	if currentPosition[0] < 0 || currentPosition[1] < 0 {
		return fmt.Errorf("current position is not inside the room")
	}
	// Check that the grid area is positive.
	if gridSize[0] < 0 || gridSize[1] < 0 {
		return fmt.Errorf("room area is negative")
	}
	// Check we have valid slice.
	if len(gridSize) > 2 {
		return fmt.Errorf("this application only supports two dimensional rooms")
	}
	return nil
}

// IsInBounds will check if a given coordinate set is within another.
// EG:
//  currentPosition: [1,2]
//  gridSize: [2,2]
//
//  Corresponds to:
//      _ _
//     |_|x|
//     |_|_|
//
func isInBounds(currentPosition []int, gridSize []int) bool {
	if isValid(currentPosition, gridSize) != nil {
		return false
	}
	// Check that the current position is inside the room.
	if currentPosition[0] <= gridSize[0] && currentPosition[1] <= gridSize[1] {
		return true
	}
	return false
}

// takeStep will attempt to take a step and then validate whether or not the resultant
// position is valid.
func takeStep(currentPosition []int, movement []int, gridSize []int) []int {
	proposedPosition := []int{currentPosition[0] + movement[0], currentPosition[1] + movement[1]}
	if isInBounds(proposedPosition, gridSize) {
		return proposedPosition
	}
	return currentPosition
}

// formatCoordinates converts [2]int to a string.
func formatCoordinates(coord []int) string {
	return fmt.Sprintf("%d:%d", coord[0], coord[1])
}

// matchCoordinates compares two coordinates for equality.
func matchCoordinates(a []int, b []int) bool {
	if a[0] == b[0] && a[1] == b[1] {
		return true
	}
	return false
}
