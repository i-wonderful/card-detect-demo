package service

import (
	"card-detect-demo/internal/model"
	"card-detect-demo/internal/util/img"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

type Recognizer interface {
	PredictBoxCoord(img image.Image) ([]model.Box, error)
}

type Detector struct {
	recognizer Recognizer
	isLogTime  bool
}

func NewDetector(recognizer Recognizer, isLogTime bool) *Detector {
	return &Detector{
		recognizer: recognizer,
		isLogTime:  isLogTime,
	}
}

func (d *Detector) Detect(imgPath string) error {

	log.Println("Detect: ", imgPath)

	// ----------------------
	im, err := img.OpenImg(imgPath)
	if err != nil {
		return err
	}
	// ----------------------

	boxes, err := d.recognizer.PredictBoxCoord(im)
	if err != nil {
		return err
	}

	// ----------------------
	for _, box := range boxes {
		log.Println(box)
	}

	drawBoxes(im, boxes)
	// ----------------------
	// todo
	return nil
}

func drawBoxes(img image.Image, boxes []model.Box) {
	// Создаем новое изображение RGBA
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	// Рисуем бокс
	for _, area := range boxes {
		// Определяем бокс (пример координат)
		box := image.Rect(area.X, area.Y, area.X+area.Width, area.Y+area.Height)
		drawBox(rgba, box, color.RGBA{255, 0, 0, 255}, 2)
	}

	// Сохраняем результат
	outFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	png.Encode(outFile, rgba)
}

func drawBox(img *image.RGBA, rect image.Rectangle, c color.Color, thickness int) {
	for i := 0; i < thickness; i++ {
		// Верхняя линия
		draw.Draw(img, image.Rect(rect.Min.X-i, rect.Min.Y-i, rect.Max.X+i, rect.Min.Y-i+1), &image.Uniform{c}, image.Point{}, draw.Src)
		// Нижняя линия
		draw.Draw(img, image.Rect(rect.Min.X-i, rect.Max.Y+i-1, rect.Max.X+i, rect.Max.Y+i), &image.Uniform{c}, image.Point{}, draw.Src)
		// Левая линия
		draw.Draw(img, image.Rect(rect.Min.X-i, rect.Min.Y-i, rect.Min.X-i+1, rect.Max.Y+i), &image.Uniform{c}, image.Point{}, draw.Src)
		// Правая линия
		draw.Draw(img, image.Rect(rect.Max.X+i-1, rect.Min.Y-i, rect.Max.X+i, rect.Max.Y+i), &image.Uniform{c}, image.Point{}, draw.Src)
	}
}
