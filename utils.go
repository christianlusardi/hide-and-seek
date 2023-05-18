package main

import (
	"bufio"
	"image"
	"os"
)

// OpenImageFromPath returns a image.Image from a file path. A helper function to deal with decoding the image into a usable format. This method is optional.
func OpenImageFromPath(filename string) (image.Image, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, nil
}
