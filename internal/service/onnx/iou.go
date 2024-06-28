package onnx

import "math"

// iou calculates the IoU between two boxes.
func iou(box1, box2 box) float64 {
	return intersection(box1, box2) / union(box1, box2)
}

func union(box1, box2 box) float64 {
	box1_area := (box1.x2 - box1.x1) * (box1.y2 - box1.y1)
	box2_area := (box2.x2 - box2.x1) * (box2.y2 - box2.y1)

	return box1_area + box2_area - intersection(box1, box2)
}

func intersection(box1, box2 box) float64 {
	box1_x1, box1_y1, box1_x2, box1_y2 := box1.x1, box1.y1, box1.x2, box1.y2
	box2_x1, box2_y1, box2_x2, box2_y2 := box2.x1, box2.y1, box2.x2, box2.y2
	x1 := math.Max(box1_x1, box2_x1)
	y1 := math.Max(box1_y1, box2_y1)
	x2 := math.Min(box1_x2, box2_x2)
	y2 := math.Min(box1_y2, box2_y2)

	// Если прямоугольники не пересекаются, площадь пересечения равна 0
	if x2 < x1 || y2 < y1 {
		return 0.0
	}

	return (x2 - x1) * (y2 - y1)
}
