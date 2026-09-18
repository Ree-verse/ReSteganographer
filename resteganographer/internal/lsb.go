/*
Copyright © 2026 Ree-verse. All rights reserved.
Licensed under the MIT License. See LICENSE in the project root for license information.
*/
package internal

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
)

const messageLengthHeaderSizeBytes = 4 // uint32 for message length
const bitsPerPixel = 3                 // R, G, B (skip alpha to preserve transparency)
const bitsPerByte = 8

// EncodeMessage hides a message inside a (PNG) image using LSB steganography.
// It prepends a 4-byte big-endian length header before the message payload.
func EncodeMessage(img image.Image, message []byte) (*image.NRGBA, error) {
	messageWithHeader := make([]byte, messageLengthHeaderSizeBytes+len(message))
	binary.BigEndian.PutUint32(messageWithHeader[:messageLengthHeaderSizeBytes], uint32(len(message)))
	copy(messageWithHeader[messageLengthHeaderSizeBytes:], message)

	messageBits := bytesToBits(messageWithHeader)

	maxCapacityBits := img.Bounds().Dx() * img.Bounds().Dy() * bitsPerPixel
	if len(messageBits) > maxCapacityBits {
		return nil, fmt.Errorf("message too large: need %d bits, image capacity is %d bits", len(messageBits), maxCapacityBits)
	}

	encodedImage := image.NewNRGBA(img.Bounds())
	bitIndex := 0

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)

			if bitIndex < len(messageBits) {
				c.R = setLSB(c.R, messageBits[bitIndex])
				bitIndex++
			}
			if bitIndex < len(messageBits) {
				c.G = setLSB(c.G, messageBits[bitIndex])
				bitIndex++
			}
			if bitIndex < len(messageBits) {
				c.B = setLSB(c.B, messageBits[bitIndex])
				bitIndex++
			}
			// (alpha channel untouched to preserve transparency)

			encodedImage.SetNRGBA(x, y, c)

			if bitIndex == len(messageBits) {
				// Fill remaining pixels unchanged
				fillRemainingPixels(encodedImage, img, x+1, y)
				return encodedImage, nil
			}
		}
	}

	return nil, fmt.Errorf("failed to encode message")
}

// DecodeMessage extracts a hidden message from a (PNG) image encoded with LSB steganography.
func DecodeMessage(img image.Image) ([]byte, error) {
	availableBits := img.Bounds().Dx() * img.Bounds().Dy() * bitsPerPixel
	headerSizeBits := messageLengthHeaderSizeBytes * bitsPerByte

	if availableBits < headerSizeBits {
		return nil, fmt.Errorf("image too small to contain a message header")
	}

	// Extract all LSBs from R, G, B channels
	allBits := extractLSBBits(img)

	// Read the 4-byte length header
	messageLength := binary.BigEndian.Uint32(bitsToBytes(allBits[:headerSizeBits]))

	if messageLength == 0 {
		return nil, fmt.Errorf("no hidden message found (length header is zero)")
	}

	totalRequiredBits := headerSizeBits + int(messageLength)*bitsPerByte
	if totalRequiredBits > len(allBits) {
		return nil, fmt.Errorf("message length header claims %d bytes, but image doesn't contain enough data", messageLength)
	}

	messageBits := allBits[headerSizeBits:totalRequiredBits]
	return bitsToBytes(messageBits), nil
}

// setLSB clears the LSB of value and sets it to bit.
func setLSB(value, bit uint8) uint8 {
	return (value & 0xFE) | (bit & 0x01)
}

// bytesToBits unpacks a byte slice into a bit slice (MSB-first per byte).
func bytesToBits(bytes []byte) []uint8 {
	bits := make([]uint8, len(bytes)*bitsPerByte)
	for i, b := range bytes {
		for j := range bitsPerByte {
			bits[i*bitsPerByte+j] = (b >> (bitsPerByte - 1 - j)) & 1
		}
	}
	return bits
}

// bitsToBytes packs a bit slice back into bytes (MSB-first per byte).
// Pads with zeros if the length isn't a multiple of 8.
func bitsToBytes(bits []uint8) []byte {
	bytes := make([]byte, (len(bits) + bitsPerByte - 1) / bitsPerByte)
	for i, b := range bits {
		if b == 1 {
			bytes[i/bitsPerByte] |= 1 << (bitsPerByte - 1 - (i % bitsPerByte))
		}
	}
	return bytes
}

// extractLSBBits reads the LSB of every R, G, B channel across all pixels.
func extractLSBBits(img image.Image) []uint8 {
	var bits []uint8
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			bits = append(bits, c.R&1, c.G&1, c.B&1)
		}
	}
	return bits
}

// fillRemainingPixels copies untouched pixels from src to dst starting after (startX, y).
func fillRemainingPixels(dst *image.NRGBA, src image.Image, startX, y int) {
	// Finish the current row
	for x := startX; x < src.Bounds().Max.X; x++ {
		dst.Set(x, y, src.At(x, y))
	}
	// Fill all subsequent rows
	for y++; y < src.Bounds().Max.Y; y++ {
		for x := src.Bounds().Min.X; x < src.Bounds().Max.X; x++ {
			dst.Set(x, y, src.At(x, y))
		}
	}
}
