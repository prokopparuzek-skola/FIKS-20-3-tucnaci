package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/spakin/netpbm"
)

const (
	BLACK = false
	WHITE = true
)

const (
	OK = ":)"
	KO = ":("
)

const root = "./data/"
const croppFactor = 2 / 5

func color(img image.Image, x, y int) bool {
	if r, _, _, _ := img.At(x, y).RGBA(); r == 65535 {
		return WHITE
	} else {
		return BLACK
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func cropp(fileName string) float64 {
	var minX, minY, maxX, maxY int
	minX, minY = 120, 120
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			if color(img, x, y) == BLACK {
				minX = min(x, minX)
				minY = min(y, minY)
				maxX = max(x, maxX)
				maxY = max(y, maxY)
			}
		}
	}
	crop1Img := img.(*netpbm.BW).SubImage(image.Rect(minX, minY, maxX+1, maxY+1))
	crop2Img := crop1Img.(*image.Paletted).SubImage(image.Rect(minX+crop1Img.Bounds().Dx()/3, minY+crop1Img.Bounds().Dy()/3, maxX+1-crop1Img.Bounds().Dx()/3, maxY+1-crop1Img.Bounds().Dy()/3))
	var black int
	for y := crop2Img.Bounds().Min.Y; y < crop2Img.Bounds().Max.Y; y++ {
		for x := crop2Img.Bounds().Min.X; x < crop2Img.Bounds().Max.X; x++ {
			if color(crop2Img, x, y) == BLACK {
				black++
			}
		}
	}
	return float64(black) / float64(crop2Img.Bounds().Dx()*crop2Img.Bounds().Dy())
}

func main() {
	var files []string
	var out []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, info.Name())
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range files[1:] {
		tmp := cropp(root + f)
		if tmp >= 0.208 {
			out = append(out, KO)
		} else {
			out = append(out, OK)
		}
	}
	for i, s := range out {
		if i%16 == 0 {
			fmt.Printf("%02d.pbm\n", i/16)
		}
		fmt.Print(s)
		if (i+1)%4 == 0 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
