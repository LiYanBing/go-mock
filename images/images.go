package images

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func GetImageId(size, background, foreground, format, text string) string {
	values := url.Values{
		"s": []string{size},
		"b": []string{background},
		"f": []string{foreground},
		"m": []string{format},
		"t": []string{text},
	}
	return values.Encode()
}

func GenImageWithID(id string, writer io.Writer) (string, error) {
	values, err := url.ParseQuery(id)
	if err != nil {
		return "", err
	}

	size := values.Get("s")
	if size == "" {
		size = "200x100"
	}

	background := values.Get("b")
	if background == "" {
		background = "#FF6600"
	}

	foreground := values.Get("f")
	if foreground == "" {
		foreground = "#FFF"
	}

	format := values.Get("m")
	if format == "" {
		format = "png"
	}

	text := values.Get("t")
	if text == "" {
		text = size
	}

	return "image/" + format, GenImage(size, background, foreground, format, text, writer)
}

func ParseSize(in string) (width, height int64) {
	splits := strings.Split(in, "x")
	if len(splits) >= 2 {
		width, _ = strconv.ParseInt(splits[0], 10, 64)

		height, _ = strconv.ParseInt(splits[1], 10, 64)
	}

	if width <= 0 {
		width = 100
	}

	if height <= 0 {
		height = 100
	}
	return
}

func ParseRGBFromHex(hex string) (r, g, b int) {
	hex = strings.TrimPrefix(hex, "#")
	color64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return
	}

	colorInt := int(color64)
	return colorInt >> 16, (colorInt & 0x00FF00) >> 8, colorInt & 0x0000FF
}

func ColorRGB() string {
	rgbValues := rgbValues()
	return fmt.Sprintf("rgb(%v, %v, %v)",
		rgbValues[0],
		rgbValues[1],
		rgbValues[2],
	)
}

func ColorRGBA() string {
	rgbValues := rgbValues()
	return fmt.Sprintf("rgba(%v, %v, %v, %.2f)",
		rgbValues[0],
		rgbValues[1],
		rgbValues[2],
		rand.Float64(),
	)
}

func ColorHex() string {
	rgbValues := rgbValues()
	return fmt.Sprintf("#%v%v%v",
		rgbValueToHex(rgbValues[0]),
		rgbValueToHex(rgbValues[1]),
		rgbValueToHex(rgbValues[2]),
	)
}

func rgbValueToHex(in uint8) string {
	return fmt.Sprintf("%.2x", in)
}

func rgbValues() [3]uint8 {
	return [3]uint8{
		uint8(rand.Intn(math.MaxUint8 + 1)),
		uint8(rand.Intn(math.MaxUint8 + 1)),
		uint8(rand.Intn(math.MaxUint8 + 1)),
	}
}

func GenImage(size, background, foreground, format, text string, writer io.Writer) error {
	width, height := ParseSize(size)
	canvas := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	r, g, b := ParseRGBFromHex(background)
	canvasBackground := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
	draw.Draw(canvas, canvas.Rect, &image.Uniform{C: canvasBackground}, image.Pt(0, 0), draw.Src)

	imageSize := canvas.Rect.Size()
	fontr, fontg, fontb := ParseRGBFromHex(foreground)
	face := basicfont.Face7x13
	face.Advance = 9

	d := font.Drawer{
		Dst: canvas,
		Src: image.NewUniform(color.RGBA{
			R: uint8(fontr),
			G: uint8(fontg),
			B: uint8(fontb),
			A: 255}),
		Face: face,
		Dot:  fixed.P(imageSize.X/2-face.Width*utf8.RuneCountInString(text)/2, imageSize.Y/2),
	}
	d.DrawString(text)

	switch format {
	case "jpeg":
		return jpeg.Encode(writer, canvas, &jpeg.Options{Quality: 100})
	default:
		return png.Encode(writer, canvas)
	}
}
