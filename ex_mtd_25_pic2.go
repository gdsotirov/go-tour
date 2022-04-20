/* Exercise: Images
 * See https://go.dev/tour/methods/25
 * Description: Generate and display image with implementation of image.Image
 */

package main

import (
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

type Image struct {
    width  int
    height int
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
    red  := uint8(x)
    green:= uint8(y)
    /* You could experiment with values of blue between 0 and 255... */
    var blue uint8 = 255
    /* and value sof alpha in the same range */
    var alpha uint8 = 255
    return color.RGBA{red, green, blue, alpha}
}

func main() {
    m := Image{255, 255}
    pic.ShowImage(m)
}
