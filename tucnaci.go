package main

import (
	"image"
	"log"
	"os"

	"github.com/spakin/netpbm"
)

const (
	BLACK = false
	WHITE = true
)

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

func main() {
	var minX, minY, maxX, maxY int
	minX, minY = 120, 120
	file, err := os.Open("data/0000.pbm")
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
	cropImg := img.(*netpbm.BW).SubImage(image.Rect(minX, minY, maxX+1, maxY+1))
	w, err := os.OpenFile("test.pbm", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer w.Close()
	netpbm.Encode(w, cropImg, &netpbm.EncodeOptions{Format: netpbm.PBM, MaxValue: 0, Plain: false, TupleType: "", Comments: nil})
}
