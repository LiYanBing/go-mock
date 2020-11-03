package images

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestColorRGB(t *testing.T) {
	t.Log(ColorRGB())
}

func TestColorHex(t *testing.T) {
	t.Log(ColorHex())
}

func TestColorRGBA(t *testing.T) {
	t.Log(ColorRGBA())
}

func TestParseRGBFromHex(t *testing.T) {
	for i := 0; i < 100000; i++ {
		rgbValues := rgbValues()
		hex := fmt.Sprintf("#%v%v%v",
			rgbValueToHex(rgbValues[0]),
			rgbValueToHex(rgbValues[1]),
			rgbValueToHex(rgbValues[2]),
		)
		r, g, b := ParseRGBFromHex(hex)
		assert.Condition(t, func() (success bool) {
			success = uint8(r) == rgbValues[0] && uint8(g) == rgbValues[1] && uint8(b) == rgbValues[2]
			return
		}, fmt.Sprintf("%v", rgbValues), fmt.Sprintf("%v %v %v", r, g, b))
	}
}

func TestGenImage(t *testing.T) {
	ff, err := os.Create("image1.png")
	if err != nil {
		t.Error(err)
		return
	}

	err = GenImage("200x100", "#894FC4", "#FFFFFF", "png", "200 x 100", ff)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGenImageWithID(t *testing.T) {

}
