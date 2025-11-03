package main

import (
	"fmt"
	"github.com/cxfredeper/image-thingy/codec"
	"image"
	"image/png"
	"os"
	"strings"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s FILE.png\n", os.Args[0])
		os.Exit(1)
	}

	// Open input file.
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil { panic(err) }

	// Prepare output file.
	outPath, found := strings.CutSuffix(path, ".png")
	if !found {
		outPath = path + ".decode"
	}
	flags := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	outFile, err := os.OpenFile(outPath, flags, 0644)
	if err != nil { panic(err) }

	// Decode
	img, err := png.Decode(file)
	if err != nil { panic(err) }
	file.Close()

	payload, err := codec.Decode(img.(*image.NRGBA))
	if err != nil { panic(err) }

	_, err = outFile.Write(payload)
	if err != nil { panic(err) }
	outFile.Close()
}
