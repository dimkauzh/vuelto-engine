//go:build windows || linux || darwin
// +build windows linux darwin

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

package event

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var KeyMap = map[string]Key{
	"A":       {Key: glfw.KeyA},
	"B":       {Key: glfw.KeyB},
	"C":       {Key: glfw.KeyC},
	"D":       {Key: glfw.KeyD},
	"E":       {Key: glfw.KeyE},
	"F":       {Key: glfw.KeyF},
	"G":       {Key: glfw.KeyG},
	"H":       {Key: glfw.KeyH},
	"I":       {Key: glfw.KeyI},
	"J":       {Key: glfw.KeyJ},
	"K":       {Key: glfw.KeyK},
	"L":       {Key: glfw.KeyL},
	"M":       {Key: glfw.KeyM},
	"N":       {Key: glfw.KeyN},
	"O":       {Key: glfw.KeyO},
	"P":       {Key: glfw.KeyP},
	"Q":       {Key: glfw.KeyQ},
	"R":       {Key: glfw.KeyR},
	"S":       {Key: glfw.KeyS},
	"T":       {Key: glfw.KeyT},
	"U":       {Key: glfw.KeyU},
	"V":       {Key: glfw.KeyV},
	"W":       {Key: glfw.KeyW},
	"X":       {Key: glfw.KeyX},
	"Y":       {Key: glfw.KeyY},
	"Z":       {Key: glfw.KeyZ},
	"Up":      {Key: glfw.KeyUp},
	"Down":    {Key: glfw.KeyDown},
	"Left":    {Key: glfw.KeyLeft},
	"Right":   {Key: glfw.KeyRight},
	"Num0":    {Key: glfw.Key0},
	"Num1":    {Key: glfw.Key1},
	"Num2":    {Key: glfw.Key2},
	"Num3":    {Key: glfw.Key3},
	"Num4":    {Key: glfw.Key4},
	"Num5":    {Key: glfw.Key5},
	"Num6":    {Key: glfw.Key6},
	"Num7":    {Key: glfw.Key7},
	"Num8":    {Key: glfw.Key8},
	"Num9":    {Key: glfw.Key9},
	"Space":   {Key: glfw.KeySpace},
	"Enter":   {Key: glfw.KeyEnter},
	"Escape":  {Key: glfw.KeyEscape},
	"Tab":     {Key: glfw.KeyTab},
	"Shift":   {Key: glfw.KeyLeftShift},
	"Control": {Key: glfw.KeyLeftControl},
	"Alt":     {Key: glfw.KeyLeftAlt},
	"F1":      {Key: glfw.KeyF1},
	"F2":      {Key: glfw.KeyF2},
	"F3":      {Key: glfw.KeyF3},
	"F4":      {Key: glfw.KeyF4},
	"F5":      {Key: glfw.KeyF5},
	"F6":      {Key: glfw.KeyF6},
	"F7":      {Key: glfw.KeyF7},
	"F8":      {Key: glfw.KeyF8},
	"F9":      {Key: glfw.KeyF9},
	"F10":     {Key: glfw.KeyF10},
	"F11":     {Key: glfw.KeyF11},
	"F12":     {Key: glfw.KeyF12},
}
