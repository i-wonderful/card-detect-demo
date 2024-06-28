package img

import (
	"image"
	"image/png"
	"log"
	"os"
)

func SaveImg(outputPath string, content []byte) error {
	file, err := os.Create(outputPath)
	if err != nil {
		log.Printf("Error creating image file: %v", err)
		return err
	}
	defer file.Close()

	if _, err = file.Write(content); err != nil {
		log.Printf("Error writing image to file: %v", err)
		return err
	}
	log.Println("Image has been saved:", outputPath)
	return nil
}

func SaveNRGBA(img *image.RGBA, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
