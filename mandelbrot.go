package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	// "strconv"
	"sync"

	"github.com/cheggaaa/pb/v3" // progress bar
)

func mandelbrot(c complex128, iterations int) int {

    var n int = 0
    z := complex(0, 0)

    for cmplx.Abs(z) <= 2 && n < iterations {
        z = z*z + c
        n++
    }
    return n

}


func main() {

    // initialize waitgroup
    var wg sync.WaitGroup

    // create vars
    iterations := 50
    width := 3000
    height := 2000
    img := image.NewRGBA(image.Rect(0, 0, width, height))

    // create progress bar
    bar := pb.StartNew(width)
    defer bar.Finish()

    // for each row in img
    for x := 0; x <= width; x++ {

        // increment progress bar
        bar.Increment()

        // for each column in img
        for y := 0; y <= height; y++ {

            // increment waitgroup
            wg.Add(1)

            // create goroutine to calculate color for (x, y)
            go func(x float64, y float64) {

                // define relative point on complex plane from pixel position
                c := complex( 3.0*x / float64(width) - 2.0 , 2.0*y / float64(height) - 1.0 )

                n := mandelbrot(c, iterations)
                var s uint8 = uint8(255 - 255*n/iterations)
                img.SetRGBA(int(x), int(y), color.RGBA{s, s, s, 255})
                wg.Done()

            }(float64(x), float64(y))

        }
    }

    wg.Wait()

    out, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0777)
	defer out.Close()

	// encode image to file
	png.Encode(out, img)


}
