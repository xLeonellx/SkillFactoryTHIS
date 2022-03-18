package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"
)

const DEFAULT_MAX_WIDTH float64 = 320
const DEFAULT_MAX_HEIGHT float64 = 240

// Рассчитываем размер изображения после масштабирования
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// Создание миниатюры
func makeMiniature(imagePath, savePath string) error {

	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// Вызов библиотеки изменения размера для увеличения изображения
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// файл для сохранения
	imgfile, _ := os.Create(savePath)
	defer imgfile.Close()

	// Сохраняем файл в формате PNG
	err = png.Encode(imgfile, m)
	if err != nil {
		return err
	}

	return nil
}
