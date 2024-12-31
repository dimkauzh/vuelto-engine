//go:build js && wasm
// +build js,wasm

package main

import (
	"embed"
	_ "image/png"
	"syscall/js"

	vimage "vuelto.pp.ua/internal/image"
)

//go:embed tree.png
var embeddedFiles embed.FS

func renderImage() {
	doc := js.Global().Get("document")
	canvas := doc.Call("createElement", "canvas")
	canvas.Set("width", 800)
	canvas.Set("height", 600)
	doc.Get("body").Call("appendChild", canvas)

	gl := canvas.Call("getContext", "webgl")
	if gl.IsNull() {
		panic("WebGL not supported")
	}

	vertexShaderSource := `
		attribute vec4 position;
		attribute vec2 texCoord;
		varying vec2 vTexCoord;
		void main() {
			gl_Position = position;
			vTexCoord = texCoord;
		}
	`

	fragmentShaderSource := `
		precision mediump float;
		varying vec2 vTexCoord;
		uniform sampler2D texture;
		void main() {
			gl_FragColor = texture2D(texture, vTexCoord);
		}
	`

	program := createProgram(gl, vertexShaderSource, fragmentShaderSource)
	gl.Call("useProgram", program)

	data := []float32{
		-1, 1, 0, 0, 1, // Top-left
		-1, -1, 0, 0, 0, // Bottom-left
		1, 1, 0, 1, 1, // Top-right
		1, -1, 0, 1, 0, // Bottom-right
	}

	indices := []uint16{
		0, 1, 2,
		2, 1, 3,
	}

	vertexBuffer := gl.Call("createBuffer")
	gl.Call("bindBuffer", gl.Get("ARRAY_BUFFER"), vertexBuffer)

	vertices := js.Global().Get("Float32Array").New(len(data))
	for i, v := range data {
		vertices.SetIndex(i, v)
	}
	gl.Call("bufferData", gl.Get("ARRAY_BUFFER"), vertices, gl.Get("STATIC_DRAW"))

	indexBuffer := gl.Call("createBuffer")
	gl.Call("bindBuffer", gl.Get("ELEMENT_ARRAY_BUFFER"), indexBuffer)

	indicesTyped := js.Global().Get("Uint16Array").New(len(indices))
	for i, v := range indices {
		indicesTyped.SetIndex(i, v)
	}
	gl.Call("bufferData", gl.Get("ELEMENT_ARRAY_BUFFER"), indicesTyped, gl.Get("STATIC_DRAW"))

	// position := gl.Call("getAttribLocation", program, "position")
	gl.Call("enableVertexAttribArray", 1)
	gl.Call("vertexAttribPointer", 0, 3, gl.Get("FLOAT"), false, 20, 0)

	// texCoord := gl.Call("getAttribLocation", program, "texCoord")
	gl.Call("enableVertexAttribArray", 0)
	gl.Call("vertexAttribPointer", 1, 2, gl.Get("FLOAT"), false, 20, 12)

	texture := gl.Call("createTexture")
	gl.Call("bindTexture", gl.Get("TEXTURE_2D"), texture)

	img := vimage.LoadAsEmbed(embeddedFiles, "tree.png")
	gl.Call("texImage2D", gl.Get("TEXTURE_2D"), 0, gl.Get("RGBA"), img.Width, img.Height, 0, gl.Get("RGBA"), gl.Get("UNSIGNED_BYTE"), img.Texture)

	gl.Call("texParameteri", gl.Get("TEXTURE_2D"), gl.Get("TEXTURE_MIN_FILTER"), gl.Get("LINEAR"))
	gl.Call("texParameteri", gl.Get("TEXTURE_2D"), gl.Get("TEXTURE_WRAP_S"), gl.Get("CLAMP_TO_EDGE"))
	gl.Call("texParameteri", gl.Get("TEXTURE_2D"), gl.Get("TEXTURE_WRAP_T"), gl.Get("CLAMP_TO_EDGE"))

	gl.Call("clearColor", 0.0, 0.0, 0.0, 1.0)
	gl.Call("clear", gl.Get("COLOR_BUFFER_BIT"))
	gl.Call("drawElements", gl.Get("TRIANGLES"), len(indices), gl.Get("UNSIGNED_SHORT"), 0)
}

func createShader(gl js.Value, shaderType int, source string) js.Value {
	shader := gl.Call("createShader", shaderType)
	gl.Call("shaderSource", shader, source)
	gl.Call("compileShader", shader)
	if !gl.Call("getShaderParameter", shader, gl.Get("COMPILE_STATUS")).Bool() {
		panic(gl.Call("getShaderInfoLog", shader).String())
	}
	return shader
}

func createProgram(gl js.Value, vertexShaderSource, fragmentShaderSource string) js.Value {
	vertexShader := createShader(gl, gl.Get("VERTEX_SHADER").Int(), vertexShaderSource)
	fragmentShader := createShader(gl, gl.Get("FRAGMENT_SHADER").Int(), fragmentShaderSource)

	program := gl.Call("createProgram")
	gl.Call("attachShader", program, vertexShader)
	gl.Call("attachShader", program, fragmentShader)
	gl.Call("linkProgram", program)
	if !gl.Call("getProgramParameter", program, gl.Get("LINK_STATUS")).Bool() {
		panic(gl.Call("getProgramInfoLog", program).String())
	}
	return program
}

func main() {
	renderImage()
}
