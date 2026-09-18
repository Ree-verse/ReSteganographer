/*
Copyright © 2026 Ree-verse. All rights reserved.
Licensed under the MIT License. See LICENSE in the project root for license information.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "resteganographer",
	Short: "PNG steganography tool using the LSB technique",
	Long: `ReSteganographer is a command-line tool for hiding and extracting secret
messages within PNG images. It uses the Least Significant Bit (LSB) technique,
modifying the lowest bit of the Red, Green, and Blue color channels to embed
data without visibly altering the image. The alpha channel is preserved to
maintain image transparency.`,
}

func Execute() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
