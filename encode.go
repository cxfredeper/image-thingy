package main

import (
	"fmt"
	"github.com/cxfredeper/image-thingy/codec"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil { panic(err) }

	content, err := io.ReadAll(file)
	if err != nil { panic(err) }
	file.Close()

	pngFile, err := os.Create(path + ".png")
	if err != nil { panic(err) }

	png, err := codec.EncodePNG(content)
	if err != nil { panic(err) }

	_, err = pngFile.Write(png)
	if err != nil { panic(err) }
	pngFile.Close()
}
