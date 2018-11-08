package resize

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func NewConfigWithDefaultEncoders(suffix string, ratio float64, resizeWidth bool) Config {
	encoders := map[string]Encoder{
		"gif": func(w io.Writer, img image.Image) error {
			return gif.Encode(w, img, nil)
		},
		"jpeg": func(w io.Writer, img image.Image) error {
			return jpeg.Encode(w, img, nil)
		},
		"png": png.Encode,
	}

	return Config{
		Suffix:      suffix,
		Ratio:       ratio,
		ResizeWidth: resizeWidth,
		Encoders:    encoders,
	}
}

type Encoder func(w io.Writer, m image.Image) error

type Config struct {
	Suffix      string
	Ratio       float64
	ResizeWidth bool
	Encoders    map[string]Encoder
}
