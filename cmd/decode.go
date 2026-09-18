/*
Copyright © 2026 Ree-verse. All rights reserved.
Licensed under the MIT License. See LICENSE in the project root for license information.
*/
package cmd

import (
	"github.com/ree-verse/ReSteganographer/resteganographer"
	"github.com/spf13/cobra"
)

var decodeInputImage string

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Extract a hidden message from a PNG image",
	Long: `Decodes and extracts a hidden message from a PNG image that was previously
encoded using this tool.

It reads the least significant bits of the RGB channels from the input image,
reconstructs the payload based on the embedded 4-byte length header, and prints
the recovered secret message to the standard output.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return resteganographer.Run(resteganographer.Config{
			Mode:       resteganographer.ModeDecode,
			InputImage: decodeInputImage,
		})
	},
}

func init() {
	decodeCmd.Flags().StringVarP(&decodeInputImage, "input", "i", "", "Input PNG image path (required)")

	_ = decodeCmd.MarkFlagRequired("input")

	rootCmd.AddCommand(decodeCmd)
}
