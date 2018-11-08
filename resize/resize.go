package resize

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func All(config Config, files []string) error {
	files = unique(files)

	if files == nil || len(files) == 0 {
		return nil
	}

	for _, file := range files {
		if err := Single(config, file); err != nil {
			return err
		}
	}

	return nil
}

func Single(config Config, file string) error {
	reader, err := os.Open(file)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, kind, err := image.Decode(reader)
	if err != nil {
		return err
	}

	encoder, ok := config.Encoders[kind]
	if !ok {
		return fmt.Errorf("format %q has no encoder", kind)
	}

	originalWidth, originalHeight := img.Bounds().Max.X, img.Bounds().Max.Y
	var resizedWidth, resizedHeight uint

	if config.ResizeWidth {
		resizedWidth = uint(float64(originalWidth) * config.Ratio)
		resizedHeight = uint(originalHeight)
	} else {
		resizedWidth = uint(originalWidth)
		resizedHeight = uint(float64(originalHeight) * (1.0 / config.Ratio))
	}

	resizedImage := resize.Resize(resizedWidth, resizedHeight, img, resize.Lanczos3)

	ext := filepath.Ext(file)
	outputFilepath := strings.TrimSuffix(file, ext) + config.Suffix + ext

	writer, err := os.Create(outputFilepath)
	if err != nil {
		return err
	}
	defer writer.Close()

	return encoder(writer, resizedImage)
}
