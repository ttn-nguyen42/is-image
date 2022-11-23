package helper

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"

	hs "github.com/corona10/goimagehash"
)

/*
 * Image types
 */
const (
	PNG  = "image/png"
	JPEG = "image/jpeg"
)

/*
 * Interfaces
 */
type HashInput interface {
	io.Reader
	Seek(offset int64, whence int) (int64, error)
}

/*
 * Get image type based on http.DetectContentType
 */
func GetImageType(buf []byte) (string, error) {
	fileType := http.DetectContentType(buf)
	if fileType == PNG {
		return PNG, nil
	}
	if fileType == JPEG {
		return JPEG, nil
	}
	return "", fmt.Errorf("not %s or %s", PNG, JPEG)
}

/*
 * Generate hashes based on the Average hash algorithm
 * For JPEG format
 */
func GetImageHashJpeg(file HashInput) (string, error) {
	/*
	 * Magic code
	 * Seek back to starting to file to prevent unexpected EOF while parsing PNG/JPEG
	 * Place right before decoding
	 */
	file.Seek(0, 0)
	dec, err := jpeg.Decode(file)
	if err != nil {
		log.Printf("Image type %s failed to DECODE: %s", JPEG, err.Error())
		return "", err
	}
	return getImageHash(dec)
}

/*
 * Generate hashes based on the Average hash algorithm
 * For PNG format
 */
func GetImageHashPng(file HashInput) (string, error) {
	/*
	 * Magic code
	 * Seek back to starting to file to prevent unexpected EOF while parsing PNG/JPEG
	 * Place right before decoding
	 */
	file.Seek(0, 0)
	dec, err := png.Decode(file)
	if err != nil {
		log.Printf("Image type %s failed to DECODE: %s", JPEG, err.Error())
		return "", err
	}
	return getImageHash(dec)
}

func getImageHash(image image.Image) (string, error) {
	res, err := hs.AverageHash(image)
	if err != nil {
		log.Printf("Image type %s failed to HASH: %s", PNG, err.Error())
		return "", err
	}
	return res.ToString(), nil
}
