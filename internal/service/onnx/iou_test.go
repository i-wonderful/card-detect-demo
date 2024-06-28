package onnx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const CLASS_CARD = "card"

func TestIoU(t *testing.T) {
	tests := []struct {
		name     string
		box1     box
		box2     box
		expected float64
	}{
		{
			name:     "No overlap",
			box1:     box{0.0, 0.0, 1.0, 1.0, CLASS_CARD, 0.7},
			box2:     box{2.0, 2.0, 3.0, 3.0, CLASS_CARD, 0.7},
			expected: 0.0,
		},
		{
			name:     "Complete overlap",
			box1:     box{0.0, 0.0, 1.0, 1.0, CLASS_CARD, 0.7},
			box2:     box{0.0, 0.0, 1.0, 1.0, CLASS_CARD, 0.7},
			expected: 1.0,
		},
		{
			name:     "Partial overlap",
			box1:     box{0.0, 0.0, 1.0, 1.0, CLASS_CARD, 0.7},
			box2:     box{0.5, 0.5, 1.5, 1.5, CLASS_CARD, 0.7},
			expected: 0.14285714285714285,
		},
		{
			name:     "One box inside another",
			box1:     box{0.0, 0.0, 2.0, 2.0, CLASS_CARD, 0.7},
			box2:     box{0.5, 0.5, 1.5, 1.5, CLASS_CARD, 0.7},
			expected: 0.25,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := iou(test.box1, test.box2)
			assert.EqualValues(t, test.expected, result, "Expected IoU to be %f, but got %f", test.expected, result)
		})
	}
}

func TestIntersection(t *testing.T) {
	testCases := []struct {
		box1     box
		box2     box
		expected float64
	}{
		{
			box{0.0, 0.0, 2, 2, CLASS_CARD, 0.7},
			box{1, 1, 3, 3, CLASS_CARD, 0.7},
			1,
		},
		{
			box{0, 0, 2, 2, CLASS_CARD, 0.7},
			box{2, 2, 4, 4, CLASS_CARD, 0.7},
			0,
		},
		{
			box{0, 0, 3, 3, CLASS_CARD, 0.7},
			box{1, 1, 2, 2, CLASS_CARD, 0.7},
			1,
		},
		{
			// Границы прямоугольников совпадают по одной стороне
			box{0, 0, 4, 4, CLASS_CARD, 0.7},
			box{0, 0, 4, 5, CLASS_CARD, 0.7},
			4 * 4,
		}, {
			// Один прямоугольник полностью содержит другой
			box{0, 0, 10, 10, CLASS_CARD, 0.7},
			box{0, 0, 8, 8, CLASS_CARD, 0.7},
			8 * 8,
		}, {
			box1:     box{0.0, 0.0, 1.0, 1.0, CLASS_CARD, 0.7},
			box2:     box{0.5, 0.5, 1.5, 1.5, CLASS_CARD, 0.7},
			expected: 0.25,
		},
	}

	for _, tc := range testCases {
		result := intersection(tc.box1, tc.box2)
		assert.Equal(t, tc.expected, result, "Expected intersection to be %f, but got %f", tc.expected, result)
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		box1     box
		box2     box
		expected float64
	}{
		{
			box1:     box{0.0, 0.0, 1.0, 1.0, CLASS_CARD, 0.7},
			box2:     box{0.5, 0.5, 1.5, 1.5, CLASS_CARD, 0.7},
			expected: 1.75,
		},
	}

	for _, tc := range testCases {
		result := union(tc.box1, tc.box2)
		assert.Equal(t, tc.expected, result, "Expected union to be %f, but got %f", tc.expected, result)
	}
}
