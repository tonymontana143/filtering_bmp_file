package tasks

import (
	"bitmap/pkg/bmp"
	"fmt"
)

func ApplyRotate(bmpFile *bmp.BMPFile, rotation string) error {
	// Apply the specified rotations to the BMP file
	err := bmpFile.ApplyRotation(rotation)
	if err != nil {
		return fmt.Errorf("error when applying rotation: %w", err)
	}

	return nil
}
