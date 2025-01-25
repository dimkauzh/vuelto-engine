//go:build windows || linux || darwin
// +build windows linux darwin

/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package gl

func Uint16ToInt32(u16 []uint16) []int32 {
	i32 := make([]int32, len(u16))

	for i, v := range u16 {
		i32[i] = int32(v)
	}
	return i32
}
