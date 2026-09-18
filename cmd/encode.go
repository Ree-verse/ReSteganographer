/*
Copyright © 2026 Ree-verse. All rights reserved.
Licensed under the MIT License. See LICENSE in the project root for license information.
*/
package cmd

import (
	"github.com/ree-verse/ReSteganographer/resteganographer"
	"github.com/spf13/cobra"
)

var (
	encodeInputImage  string
	encodeOutputImage string
	encodeSecretMsg   string
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Hide a secret message inside a PNG image",
	Long: `Encodes a secret message into a PNG image using LSB steganography. The command
reads the input PNG image, embeds the provided message into the least
significant bits of the RGB channels, and saves the resulting image to the
specified output path.

A 4-byte length header is automatically prepended to the message to facilitate
accurate decoding later. The original image's alpha channel and visual quality
remain completely intact.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return resteganographer.Run(resteganographer.Config{
			Mode:          resteganographer.ModeEncode,
			InputImage:    encodeInputImage,
			OutputImage:   encodeOutputImage,
			SecretMessage: encodeSecretMsg,
		})
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&encodeInputImage, "input", "i", "", "Input PNG image path (required)")
	encodeCmd.Flags().StringVarP(&encodeOutputImage, "output", "o", "", "Output PNG image path (required)")
	encodeCmd.Flags().StringVarP(&encodeSecretMsg, "message", "m", "", "Secret message to hide (required)")

	_ = encodeCmd.MarkFlagRequired("input")
	_ = encodeCmd.MarkFlagRequired("output")
	_ = encodeCmd.MarkFlagRequired("message")

	rootCmd.AddCommand(encodeCmd)
}
