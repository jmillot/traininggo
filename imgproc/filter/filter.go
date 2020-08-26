package filter

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

// Filter is an interface
type Filter interface {
	Process(srcPath, dstPath string) error
}

// Grayscale is an empty structure
type Grayscale struct{}

// Process implemente the process method of Filter Interface
func (g Grayscale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	dst := imaging.Grayscale(src)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}

// Blur is a struct for Blur filter
type Blur struct{}

// Process implemente the process method of Filter Interface
func (b Blur) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	dst := imaging.Blur(src, 3)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}
