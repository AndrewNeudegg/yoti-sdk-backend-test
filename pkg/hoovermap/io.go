package hoovermap

// APIInput specifies the expected JSON body payload.
type APIInput struct {
	RoomSize     []int   `json:"roomSize"`
	Coords       []int   `json:"coords"`
	Patches      [][]int `json:"patches"`
	Instructions string  `json:"instructions"`
}

// APIOutput specifies the output JSON body payload.
type APIOutput struct {
	Coords  []int `json:"coords"`
	Patches int   `json:"patches"`
}
