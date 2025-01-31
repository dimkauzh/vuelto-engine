package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func pixelToGL(pixelX, pixelY float32, screenWidth, screenHeight int) (float32, float32) {
	glX := 2.0*float32(pixelX)/float32(screenWidth) - 1.0
	glY := 1.0 - 2.0*float32(pixelY)/float32(screenHeight) // Flip Y since OpenGL Y is inverted
	return glX, glY
}

func main() {
	fontUrl := "https://github.com/fusionengine-org/fusion/raw/refs/heads/main/src/fusionengine/external/font.ttf"

	if !(len(fontUrl) > 4 && (fontUrl[:7] == "http://" || fontUrl[:8] == "https://")) {
		panic("LoadFontAsHTTP() only supports HTTP and HTTPS paths")
	}

	resp, err := http.Get(fontUrl)
	if err != nil {
		log.Fatalf("Failed to fetch font '%s': %v", fontUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch font '%s': status code %d", fontUrl, resp.StatusCode)
	}

	fontBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read font data from '%s': %v", fontUrl, err)
	}

	parsedFont, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Create a font face
	const fontSize = 12
	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72, // Standard DPI
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer face.Close()

	// The text to measure
	text := "test"

	// Get text bounding box
	bounds, _ := font.BoundString(face, text)

	// Convert bounds to pixel values
	textWidth := (bounds.Max.X - bounds.Min.X).Ceil()
	textHeight := (bounds.Max.Y - bounds.Min.Y).Ceil()

	// Define OpenGL viewport size (this should be your actual screen size)
	screenWidth, screenHeight := 500, 500 // Example: 800x600 window

	// Convert text pixel size to OpenGL coordinates
	glX1, glY1 := pixelToGL(0.5, 0.5, screenWidth, screenHeight)                                // Bottom-left corner
	glX2, glY2 := pixelToGL(float32(textWidth), float32(textHeight), screenWidth, screenHeight) // Top-right corner

	// Print OpenGL coordinates
	fmt.Printf("OpenGL Coordinates:\n")
	fmt.Printf("Bottom-Left  (%.2f, %.2f)\n", glX1, glY1)
	fmt.Printf("Top-Right    (%.2f, %.2f)\n", glX2, glY2)
}
