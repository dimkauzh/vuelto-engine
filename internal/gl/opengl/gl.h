#ifndef GL_H
#define GL_H

#ifdef __APPLE__
#define GL_SILENCE_DEPRECATION
#include <OpenGL/gl.h>

#elif __linux__
#define GL_SILENCE_DEPRECATION
#include <GL/gl.h>

#elif _WIN32
#define GL_CLAMP_TO_EDGE 0x812f
#include <windows.h>
#include <GL/gl.h>

#elif _WIN64
#define GL_CLAMP_TO_EDGE 0x812f
#include <windows.h>
#include <GL/gl.h>

#else
#include <GL/gl.h>

#endif

#endif
