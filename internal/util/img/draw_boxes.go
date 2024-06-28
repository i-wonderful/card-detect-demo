package img

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

func DrawBox(img *image.RGBA, rect image.Rectangle, c color.Color, thickness int, label string) {
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

	// Добавляем надпись
	point := fixed.Point26_6{
		X: fixed.Int26_6(rect.Min.X * 64),
		Y: fixed.Int26_6((rect.Min.Y - 5) * 64), // Немного выше бокса
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
