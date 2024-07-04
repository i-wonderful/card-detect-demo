package boxes

import (
	. "card-detect-demo/internal/model"
	"card-detect-demo/internal/service/onnx"
)

// MergeCardBoxes - combines all card objects into one card
func MergeCardBoxes(boxes []Box) []Box {
	var cardBox *Box
	result := make([]Box, 0, len(boxes))

	for _, box := range boxes {
		if box.Label == onnx.CLASS_CARD {
			if cardBox == nil {
				cardBox = &Box{
					Label:  onnx.CLASS_CARD,
					X:      box.X,
					Y:      box.Y,
					Width:  box.Width,
					Height: box.Height,
				}
			} else {
				// Расширяем cardBox, чтобы включить текущий box
				minX := min(cardBox.X, box.X)
				minY := min(cardBox.Y, box.Y)
				maxX := max(cardBox.X+cardBox.Width, box.X+box.Width)
				maxY := max(cardBox.Y+cardBox.Height, box.Y+box.Height)

				cardBox.X = minX
				cardBox.Y = minY
				cardBox.Width = maxX - minX
				cardBox.Height = maxY - minY
			}
		} else {
			result = append(result, box)
		}
	}

	if cardBox != nil {
		result = append(result, *cardBox)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
