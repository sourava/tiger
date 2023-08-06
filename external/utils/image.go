package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

var ErrImageTypeUnsupported error = errors.New("error unsupported image type, supported types are PNG and JPEG")

const PngImageType string = "png"
const JpegImageType string = "jpeg"

func imageToBytes(img image.Image, imageType string) ([]byte, error) {
	buf := new(bytes.Buffer)

	switch imageType {
	case PngImageType:
		err := png.Encode(buf, img)
		if err != nil {
			return nil, err
		}
	case JpegImageType:
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrImageTypeUnsupported
	}

	return buf.Bytes(), nil
}

func ResizeImage(base64Image string, finalWidth uint, finalHeight uint) (string, error) {
	base64ImageBytes, err := base64.StdEncoding.DecodeString(base64Image[strings.IndexByte(base64Image, ',')+1:])
	if err != nil {
		return "", err
	}

	image, imageType, err := image.Decode(bytes.NewReader(base64ImageBytes))
	if err != nil {
		return "", err
	}

	resizedImage := resize.Resize(finalWidth, finalHeight, image, resize.Lanczos2)

	imageBytes, err := imageToBytes(resizedImage, imageType)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(imageBytes), nil
}
