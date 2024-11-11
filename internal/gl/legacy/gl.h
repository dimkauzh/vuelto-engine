#ifndef GL_H
#define GL_H
#define GL_SILENCE_DEPRECATION

#ifdef __APPLE__
#include <OpenGL/gl.h>

#elif __linux__
#include <GL/gl.h>

#elif defined(_WIN32) || defined(_WIN64)
#define GL_CLAMP_TO_EDGE 0x812f
#include <windows.h>
#include <GL/gl.h>

#else
#include <GL/gl.h>

#endif

#endif
