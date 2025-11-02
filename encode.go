package main

import (
	"github.com/cxfredeper/image-thingy/codec";
	"fmt";
	"image/png";
	"io";
	"os";
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

	img, err := codec.Encode(content)
	if err != nil { panic(err) }

	pngFile, err := os.Create(path + ".png")
	if err != nil { panic(err) }

	err = png.Encode(pngFile, img)
	if err != nil { panic(err) }
	pngFile.Close()
}
