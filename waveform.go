package waveform

import (
	"image"
	"image/color"
	"math/rand"
)

// WaveReader interface ...
type WaveReader interface {
	Len() uint64
	Rate() uint
	Chans() uint
	At(ch uint, offset uint64) float32
}

// Options represents ...
type Options struct {
	Width  int
	Height int
	Half   bool
	Zoom   float32
	Back   *color.RGBA
	Front  *color.RGBA
}

func initOptions(o *Options) *Options {
	no := &Options{
		Width:  800,
		Height: 250,
		Front: &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Back: &color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		},
	}
	if o != nil {
		if o.Half {
			no.Half = true
		}
		if o.Width > 0 {
			no.Width = o.Width
		}
		if o.Height > 0 {
			no.Height = o.Height
		}
		if o.Back != nil {
			no.Back = o.Back
		}
		if o.Front != nil {
			no.Front = o.Front
		}
		if o.Zoom != 0 {
			no.Zoom = o.Zoom
		}
	}
	return no
}

func newRGBA(o *Options) *image.RGBA {
	if o == nil {
		panic("options 'o' is nil")
	}
	rc := image.Rect(0, 0, o.Width, o.Height)
	im := image.NewRGBA(rc)
	for y := 0; y < o.Height; y++ {
		for x := 0; x < o.Width; x++ {
			im.SetRGBA(x, y, *o.Back)
		}
	}
	return im
}

// Max function ...
func Max(w WaveReader, o *Options) *image.RGBA {
	o = initOptions(o)
	wf := newRGBA(o)
	if uint64(o.Width) < w.Len() {
		for x := 0; x < o.Width; x++ {
			h := rand.Float32()
			m := float32(o.Height) / 2
			H := h * m
			t := m - H*o.Zoom
			b := m + H*o.Zoom
			if o.Half {
				b = float32(o.Height)
				t = b - h*b*o.Zoom
			}
			for y := int(t); y < int(b); y++ {
				wf.SetRGBA(x, y, *o.Front)
			}
		}
		return wf
	}
	return wf
}

// Rms function ...
func Rms(w WaveReader, o *Options) *image.RGBA {
	o = initOptions(o)
	wf := newRGBA(o)
	return wf
}
