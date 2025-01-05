# 🔊 Audio docs

Features for playing audio files, supporting `mp3` and `wav`.

## 📜 Usage

First, open a file with `OpenAudioFile()`. This func takes one arg, the file path to your audio file. Returns an `AudioPlayer` struct.

```go
myAudioFile := OpenAudioFile("path/to/audio.mp3")
```

You can now call the `Start()` and `Stop()` methods passing your audio player.

```go
// play!
myAudioFile.Start()

// ...

// stop.
myAudioFile.Stop()
```

If you don't need the audio anymore, you should `myAudioFile.Close()` to close the file.
