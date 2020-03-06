package waveform

import (
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"testing"
)

type dummyWave []float32

func newDummyWave() dummyWave {
	w := make([]float32, 48000)
	src := rand.NewSource(0)
	rnd := rand.New(src)
	for n := 0; n < len(w); n++ {
		w[n] = rnd.Float32()
	}
	return w
}

func (dw dummyWave) Len() uint64 {
	return uint64(len(dw))
}

func (dw dummyWave) Rate() uint {
	return 48000
}

func (dw dummyWave) Chans() uint {
	return 1
}

func (dw dummyWave) At(ch uint, offset uint64) float32 {
	if ch != 0 {
		panic("invalid channel")
	}
	return dw[offset]
}

func TestMax(t *testing.T) {
	dw := newDummyWave()
	im := Max(dw, &Options{
		Width: 1800,
		Front: &color.RGBA{
			R: 50,
			G: 100,
			B: 200,
			A: 255,
		},
		Back: &color.RGBA{
			A: 0,
		},
	})
	w, err := os.Create("test-max.png")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	err = png.Encode(w, im)
	if err != nil {
		t.Fatalf("png.Encode failed: %v", err)
	}
}

func TestMaxHalf(t *testing.T) {
	dw := newDummyWave()
	im := Max(dw, &Options{
		Width: 1800,
		Half:  true,
		Front: &color.RGBA{
			R: 10,
			G: 50,
			B: 250,
			A: 255,
		},
		Back: &color.RGBA{
			A: 0,
		},
	})
	w, err := os.Create("test-max-half.png")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	err = png.Encode(w, im)
	if err != nil {
		t.Fatalf("png.Encode failed: %v", err)
	}
}
