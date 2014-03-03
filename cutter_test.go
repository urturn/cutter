package cutter

import (
	"image"
	"os"
	"testing"
)

func TestCutter_Crop(t *testing.T) {
	img := getGopherImage()

	c := Cutter{512, 400, 0, 0}
	r, err := c.Crop(img)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 512 {
		t.Error("Bad width should be 512 but is", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 400 {
		t.Error("Bad width should be 400 but is", r.Bounds().Dy())
	}
	if r.Bounds().Min.X != 0 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 0 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}

func TestCutter_CenteredCrop(t *testing.T) {
	img := getGopherImage()

	c := Cutter{
		Width:  512,
		Height: 400,
	}
	r, err := c.CropCenter(img)
	if err != nil {
		t.Fatal(err)
	}
	if r.Bounds().Dx() != 512 {
		t.Error("Bad width should be 512 but is", r.Bounds().Dx())
	}
	if r.Bounds().Dy() != 400 {
		t.Error("Bad width should be 512 but is", r.Bounds().Dy())
	}
	if r.Bounds().Min.X != 544 {
		t.Error("Invalid Bounds Min X", r.Bounds().Min.X)
	}
	if r.Bounds().Min.Y != 518 {
		t.Error("Invalid Bounds Min Y", r.Bounds().Min.Y)
	}
}

func getGopherImage() image.Image {
	fi, err := os.Open("fixtures/gopher.jpg")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	img, _, err := image.Decode(fi)
	if err != nil {
		panic(err)
	}
	return img
}
