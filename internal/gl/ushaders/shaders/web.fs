#version 300 es

precision mediump float;

in vec2 TexCoord;
out vec4 FragColor;

uniform sampler2D ourTexture;
uniform bool useTexture;
uniform vec4 uniformColor;

void main() {
    if (useTexture) {
        FragColor = texture(ourTexture, TexCoord);
    } else {
        FragColor = uniformColor;
    }
}
