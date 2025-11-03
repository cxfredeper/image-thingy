package main

import (
	"fmt"
	"github.com/cxfredeper/image-thingy/codec"
	"io"
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
	imgFile, err := os.Open(path)
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
	imgData, err := io.ReadAll(imgFile)
	if err != nil { panic(err) }
	imgFile.Close()

	payload, err := codec.DecodePNG(imgData)
	if err != nil { panic(err) }

	_, err = outFile.Write(payload)
	if err != nil { panic(err) }
	outFile.Close()
}
