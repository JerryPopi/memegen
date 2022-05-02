package main

import (
	// "fmt"
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"image/color"
	"io/ioutil"
	math_rand "math/rand"
	"os"

	"github.com/fogleman/gg"
)

func init(){
	var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	templatePath := os.Args[1]
	outPath := os.Args[2]
	name := os.Args[3]

	files, err := ioutil.ReadDir(templatePath)
	chk(err)

    randomIndex := math_rand.Intn(len(files))
	fmt.Println(len(files))
	fmt.Println(randomIndex)
    template := files[randomIndex]
    

	bgImage, err := gg.LoadImage(templatePath + template.Name())
	chk(err)
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)
	
	maxWidth := float64(imgWidth) - 20
	
	S := float64(imgWidth)
	n := 6 // "stroke" size
	msg := "Obicham altki mnogo mnogo mnogo mnogo"

	if err := dc.LoadFontFace("impact.ttf", 40); err != nil {
		panic(err)
	}
	
	// Tuka e gorniq text
	dc.SetColor(color.Black)
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := S/2 + float64(dx)
			y := 50 + float64(dy)

			dc.DrawStringWrapped(msg, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
		}
	}
	dc.SetColor(color.White)
	dc.DrawStringWrapped(msg, S/2, 50, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	

	// Tuka e dolniq text
	dc.SetColor(color.Black)
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := S/2 + float64(dx)
			y := S - 50 + float64(dy)

			dc.DrawStringWrapped(msg, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
		}
	}
	dc.SetColor(color.White)
	dc.DrawStringWrapped(msg, S/2, S - 50, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	dc.SavePNG(outPath + name + ".png")
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
