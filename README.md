# ReSteganographer

[![Go Reference](https://pkg.go.dev/badge/github.com/ree-verse/ReSteganographer.svg)](https://pkg.go.dev/github.com/ree-verse/ReSteganographer)

A minimalist command-line tool for hiding and extracting secret messages within PNG images using Least Significant Bit (LSB) steganography. It embeds data into the RGB channels while preserving the alpha channel to maintain image transparency. A 4-byte length header is automatically prepended to the payload to facilitate decoding.

## Installation

```bash
go install github.com/ree-verse/ReSteganographer@latest
```

## Usage

Encode a secret message:
```bash
resteganographer encode -i input.png -o output.png -m "Your secret message"
```

Decode a hidden message:
```bash
resteganographer decode -i output.png
```

## Example

| Original | Encoded |
|:---:|:---:|
| ![Original image](https://raw.githubusercontent.com/ree-verse/ReSteganographer/main/examples/original.png) | ![Encoded image](https://raw.githubusercontent.com/ree-verse/ReSteganographer/main/examples/encoded.png) |
| - | Hidden message: `Hey, this tool works (really) well, right?` |

## Code note

Certain functions and variables are intentionally only used once, for readability purposes.

## Support

If you have questions, suggestions, or want to hang out with other developers, join my Discord server: [Ree-verse GitHub Support](https://discord.gg/ZZfqH9Z4uQ).

## License

[MIT](LICENSE) © 2026 Ree-verse
