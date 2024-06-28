package service

import (
	"card-detect-demo/internal/model"
	"card-detect-demo/internal/util/img"
	"github.com/google/uuid"
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"
)

type Recognizer interface {
	PredictBoxCoord(img image.Image) ([]model.Box, error)
}

type Detector struct {
	recognizer  Recognizer
	pathStorage string
	isLogTime   bool
}

func NewDetector(recognizer Recognizer, pathStorage string, isLogTime bool) *Detector {
	return &Detector{
		recognizer:  recognizer,
		pathStorage: pathStorage,
		isLogTime:   isLogTime,
	}
}

func (d *Detector) Detect(imgPath string) ([]model.Box, string, error) {
	if d.isLogTime {
		start := time.Now()
		defer func() {
			log.Printf(">>> Time detect: %s", time.Since(start))
		}()
	}

	im, err := img.OpenImg(imgPath)
	if err != nil {
		return nil, "", err
	}

	boxes, err := d.recognizer.PredictBoxCoord(im)
	if err != nil {
		return nil, "", err
	}

	outputImgPath := drawBoxes(im, boxes, d.pathStorage)
	return boxes, outputImgPath, nil
}

// drawBoxes - рисует боксы на изображении
// @return путь к сохраненному изображению
func drawBoxes(im image.Image, boxes []model.Box, pathStorage string) string {
	bounds := im.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, im, bounds.Min, draw.Src)

	for _, box := range boxes {
		rect := image.Rect(box.X, box.Y, box.X+box.Width, box.Y+box.Height)
		img.DrawBox(rgba, rect, color.RGBA{255, 0, 0, 255}, 2, box.Label)
	}

	outputFilePath := pathStorage + "/" + uuid.New().String() + ".png"
	img.SaveNRGBA(rgba, outputFilePath)

	return outputFilePath
}
