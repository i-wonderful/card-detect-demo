package boxes

import (
	"reflect"
	"testing"

	. "card-detect-demo/internal/model"
)

func TestMergeCardBoxes(t *testing.T) {
	tests := []struct {
		name     string
		input    []Box
		expected []Box
	}{
		{
			name: "No card boxes",
			input: []Box{
				{Label: "not_card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "also_not_card", X: 20, Y: 20, Width: 5, Height: 5},
			},
			expected: []Box{
				{Label: "not_card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "also_not_card", X: 20, Y: 20, Width: 5, Height: 5},
			},
		},
		{
			name: "Single card box",
			input: []Box{
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "not_card", X: 20, Y: 20, Width: 5, Height: 5},
			},
			expected: []Box{
				{Label: "not_card", X: 20, Y: 20, Width: 5, Height: 5},
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
			},
		},
		{
			name: "Multiple card boxes",
			input: []Box{
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "not_card", X: 20, Y: 20, Width: 5, Height: 5},
				{Label: "card", X: 5, Y: 5, Width: 15, Height: 15},
			},
			expected: []Box{
				{Label: "not_card", X: 20, Y: 20, Width: 5, Height: 5},
				{Label: "card", X: 0, Y: 0, Width: 20, Height: 20},
			},
		},
		{
			name: "Overlapping card boxes",
			input: []Box{
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "card", X: 5, Y: 5, Width: 10, Height: 10},
			},
			expected: []Box{
				{Label: "card", X: 0, Y: 0, Width: 15, Height: 15},
			},
		},
		{
			name: "Only card boxes",
			input: []Box{
				{Label: "card", X: 0, Y: 0, Width: 10, Height: 10},
				{Label: "card", X: 20, Y: 20, Width: 5, Height: 5},
			},
			expected: []Box{
				{Label: "card", X: 0, Y: 0, Width: 25, Height: 25},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeCardBoxes(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeCardBoxes() = %v, want %v", result, tt.expected)
			}
		})
	}
}
