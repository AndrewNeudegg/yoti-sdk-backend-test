package hoovermap

import (
	"reflect"
	"testing"
)

// TestComputeMovement performs some basic offline checks.
func TestComputeMovement(t *testing.T) {
	type args struct {
		inputs *APIInput
	}
	tests := []struct {
		name    string
		args    args
		want    *APIOutput
		wantErr bool
	}{
		{
			name: "T0",
			args: args{
				inputs: &APIInput{
					RoomSize:     []int{2, 2},
					Coords:       []int{1, 1},
					Patches:      [][]int{[]int{1, 1}},
					Instructions: "",
				}},
			want:    &APIOutput{Coords: []int{1, 1}, Patches: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ComputeMovement(tt.args.inputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeMovement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeMovement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_isInBounds sanity checks the isInBounds method.
func Test_isInBounds(t *testing.T) {
	type args struct {
		currentPosition []int
		gridSize        []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Small Grid", args: args{currentPosition: []int{0, 0}, gridSize: []int{1, 1}}, want: true},
		{name: "Large Grid", args: args{currentPosition: []int{0, 0}, gridSize: []int{100000000, 10000000}}, want: true},
		{name: "Off Grid Position", args: args{currentPosition: []int{-1, 0}, gridSize: []int{1, 1}}, want: false},
		{name: "Negative grid size", args: args{currentPosition: []int{0, 0}, gridSize: []int{-1, -1}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInBounds(tt.args.currentPosition, tt.args.gridSize); got != tt.want {
				t.Errorf("isInBounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
