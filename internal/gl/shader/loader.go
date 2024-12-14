//go:build windows || linux || darwin || js || wasm
// +build windows linux darwin js wasm

/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package shader

import (
	"log"
	"os"
)

func LoadShader(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read shader file: %s", err)
		return ""
	}

	return string(bytes)
}
