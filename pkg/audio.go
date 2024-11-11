/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the Vuelto License V1.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package vuelto

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

type AudioPlayer struct {
	Streamer beep.StreamSeekCloser
	Format   beep.Format
	File     *os.File
	Done     chan bool
}

// Opens a audio file. It supports two file formats: WAV and MP3.
// Plays the audio file using Start function. Stop the audio file using Stop function.
// Close the audio file using Close function.
func OpenAudioFile(filePath string) *AudioPlayer {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening audio file: ", filePath)
	}

	var streamer beep.StreamSeekCloser
	var format beep.Format

	switch strings.ToLower(filepath.Ext(filePath)) {

	case ".wav":
		streamer, format, err = wav.Decode(f)
	case ".mp3":
		streamer, format, err = mp3.Decode(f)
	default:
		log.Fatal("Unsupported audio format: ", filepath.Ext(filePath))
		f.Close()
	}

	if err != nil {
		log.Fatal("Error decoding audio file: ", filePath)
		f.Close()
	}

	return &AudioPlayer{
		Streamer: streamer,
		Format:   format,
		File:     f,
		Done:     make(chan bool),
	}
}

// Starts playing the audio file.
func (a *AudioPlayer) Start() {
	speaker.Init(a.Format.SampleRate, a.Format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(a.Streamer, beep.Callback(func() {
		a.Done <- true
	})))
}

// Stops playing the audio file.
func (a *AudioPlayer) Stop() {
	speaker.Clear()
	a.Streamer.Seek(0)
}

// Closes the audio file.
func (a *AudioPlayer) Close() {
	a.File.Close()
}
