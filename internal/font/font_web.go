//go:build js || wasm
// +build js wasm

/*
 * Copyright (C) 2025 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package font

import (
	"embed"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"syscall/js"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"vuelto.pp.ua/internal/gl/webgl"
)

type Font struct {
	Text                    string
	Vertices                []float32
	Texture                 js.Value
	Width, Height           float32
	Widthbound, Heightbound int
}

func Load(width, height int, fontPath string, text string, size int, x, y float32) *Font {
	if os.Getenv("VUELTO_DISABLE_BUILD_ERRORS") == "" {
		panic("Load() is not supported in web assembly")
	} else {
		return &Font{}
	}
}

func LoadAsEmbed(width, height int, fs embed.FS, fontPath string, text string, size int, x, y float32) *Font {
	fontBytes, err := fs.ReadFile(fontPath)
	if err != nil {
		log.Fatalf("Failed to read embedded font '%s': %v", fontPath, err)
	}
	return createFontTexture(width, height, fontBytes, size, text, x, y)
}

func LoadAsHTTP(width, height int, fontUrl, text string, size int, x, y float32) *Font {
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

	return createFontTexture(width, height, fontBytes, size, text, x, y)
}

func createFontTexture(w, h int, fontBytes []byte, fontSize int, text string, x, y float32) *Font {
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	fontBounds, _ := font.BoundString(face, text)
	widthBounds := (fontBounds.Max.X - fontBounds.Min.X).Ceil() + 8
	heightBounds := (fontBounds.Max.Y - fontBounds.Min.Y).Ceil() + 8

	img := image.NewRGBA(image.Rect(0, 0, widthBounds, heightBounds))

	drawer := font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
	}

	drawer.Dot = fixed.Point26_6{
		X: fixed.I(0),
		Y: fixed.I(heightBounds),
	}

	drawer.DrawString(text)

	rgbaImg := image.NewRGBA(img.Bounds())
	for y := 0; y < heightBounds; y++ {
		for x := 0; x < widthBounds; x++ {
			rgbaImg.Set(x, heightBounds-1-y, img.At(x, y))
		}
	}

	width := (float32(widthBounds) / float32(w)) * 2.0
	height := (float32(heightBounds) / float32(h)) * 2.0

	return &Font{
		Text:    text,
		Texture: webgl.NewUint8Array(rgbaImg.Pix),

		Width:       width,
		Height:      height,
		Widthbound:  widthBounds,
		Heightbound: heightBounds,
	}
}
