package main

var VertexShaderSource string = `
		#version 330 core
		layout(location = 0) in vec3 aPos;
		void main() {
			gl_Position = vec4(aPos, 1.0);
		}
	`

var FragmentShaderSource string = `
		#version 330 core
		out vec4 FragColor;
		void main() {
			FragColor = vec4(1.0, 0.5, 0.2, 1.0);
		}
	`

var WebVertexShaderSource string = `#version 300 es
precision mediump float;

layout(location = 0) in vec3 aPos;

void main() {
    gl_Position = vec4(aPos, 1.0);
}`

var WebFragmentShaderSource string = `#version 300 es
precision mediump float;

out vec4 FragColor;

void main() {
    FragColor = vec4(1.0, 0.5, 0.2, 1.0);
}`
