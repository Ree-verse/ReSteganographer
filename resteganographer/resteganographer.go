/*
Copyright © 2026 Ree-verse. All rights reserved.
Licensed under the MIT License. See LICENSE in the project root for license information.
*/
package resteganographer

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/ree-verse/ReSteganographer/resteganographer/internal"
)

type SteganographyMode int

const (
	ModeEncode SteganographyMode = iota
	ModeDecode
)

type Config struct {
	Mode          SteganographyMode
	InputImage    string
	OutputImage   string
	SecretMessage string
}

func Run(cfg Config) error {
	switch cfg.Mode {
	case ModeEncode:
		return encode(cfg)
	case ModeDecode:
		return decode(cfg)
	default:
		return fmt.Errorf("unknown steganography mode")
	}
}

func encode(cfg Config) error {
	srcImage, err := decodePNG(cfg.InputImage)
	if err != nil {
		return fmt.Errorf("decoding input image: %w", err)
	}

	encodedImage, err := internal.EncodeMessage(srcImage, []byte(cfg.SecretMessage))
	if err != nil {
		return fmt.Errorf("encoding message: %w", err)
	}

	if err := encodePNG(cfg.OutputImage, encodedImage); err != nil {
		return fmt.Errorf("encoding output image: %w", err)
	}

	fmt.Println("Message encoded successfully into", cfg.OutputImage)
	return nil
}

func decode(cfg Config) error {
	srcImage, err := decodePNG(cfg.InputImage)
	if err != nil {
		return fmt.Errorf("decoding input image: %w", err)
	}

	hiddenMessage, err := internal.DecodeMessage(srcImage)
	if err != nil {
		return fmt.Errorf("decoding message: %w", err)
	}

	fmt.Println("Hidden message:", string(hiddenMessage))
	return nil
}

func decodePNG(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return png.Decode(file)
}

func encodePNG(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
