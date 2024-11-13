<h1 align="center">
<p align="center">
<img width="1400" alt="banner" src="/logo/banner-dev.png">
  <a href="https://github.com/vuelto-org/vuelto"><img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/vuelto-org/vuelto"></a>
  <a href="https://github.com/vuelto-org/license"><img alt="License" src="https://img.shields.io/badge/license-VLv1-blue"></a>
  <a href="https://github.com/vuelto-org/vuelto"><img alt="CI Check" src="https://github.com/vuelto-org/vuelto/actions/workflows/ci_check.yml/badge.svg"></a>
  <a href="https://github.com/vuelto-org/vuelto"><img alt="Lines of code" src="https://www.aschey.tech/tokei/github/vuelto-org/vuelto"></a>
  <a href="https://goreportcard.com/report/github.com/vuelto-org/vuelto"><img alt="Report card" src="https://goreportcard.com/badge/github.com/vuelto-org/vuelto"></a>
  <a href="https://www.opengl.org/Documentation/Specs.html"><img alt="Powered By" src="https://img.shields.io/badge/powered_by-GL_3.3-blue"></a>
</p>
</h1>

Vuelto is an open-source, fast, and light game engine, based on Golang, CGo, and OpenGL. It's really easy to use, yet very powerful, and it also supports cross-platform compiling.

## ✨ Features
- 🌍 Cross Platform
- 🛠️ Open-Source
- 📚 Easy to learn
- 🚀 Fully built using CGo (and some other libraries)

## 📦 Installation

### 📋 Requirements
You need to have the following installed on your system:
- 🖥️ A C compiler
- 🔧 A Go compiler
- 🪟 Xorg/Wayland development packages (For Linux only)
- 🖱️ Supported platform

For an installation guide, [go here](https://vuelto-org.github.io/vuelto/install/).

### 🐹 Gopkg

You can get the latest Go package by running this command:
```sh
go get vuelto.me@latest
```

## 🖼️ Vuelto Example

```go
package main

import (
	vuelto "vuelto.me/pkg"
)

func main() {
	w := vuelto.NewWindow("Image Example - Vuelto", 800, 600, false)
	ren := w.NewRenderer2D()

	image := ren.LoadImage("test/image.png", 300, 300, 250, 250)
	image1 := ren.LoadImage("test/image.png", 100, 100, 150, 150)

	for !w.Close() {
		image.Draw()
		image1.Draw()
		w.Refresh()
	}
}
```

## 🖥️ Platform Support

| Platform | Status |
| :---- | :---- |
| Windows | ✅ |
| macOS (Darwin) | ✅ |
| Linux | ✅ |
| Web | ❌ (*Work In Progress, V1.1*) |

## 📖 Docs

You can check out Vuelto's documentation at [Vuelto's website](https://vuelto.me/docs/).

> [!TIP]
> In case the documentation is missing something or there is something wrong, use the [GoDoc](https://pkg.go.dev/vuelto.me) page for API documentation. Use the [(TODO: ADD LINK) examples](http://examples) directory for usage examples.

### 🛣️ Roadmap

Our roadmap is available on the website (URL here) or in the ROADMAP.md file in the root of the GitHub repo.

### 🤝 Contributing

We're really thankful for your contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

1. 🍴 Fork the repository
2. 🌟 Create your feature branch (`git checkout -b feature/amazing-feature`)
3. 📝 Commit your changes (`git commit -m 'Add some amazing feature'`)
4. 🚀 Push to the branch (`git push origin feature/amazing-feature`)
5. 🔄 Open a Pull Request

## 🛡️ Support & Security

### 🐛 Issues

See the [Issues](https://github.com/vuelto-org/vuelto/issues) page for current bugs and feature requests. In case you find any issues, please open an issue or search for any other form of contact to submit a bug report.

#### 🔒 Security Issues

If you find a security vulnerability, please follow the instructions in [SECURITY.md](SECURITY.md) to safely report it.

### 🔐 License

Vuelto is licensed under a customized version of [VLv1 Licence](LICENSE.md). Any PRs that primarily focus on changing the license won't be accepted.

### 🌐 Community and Contact

You can contact us via our Discord community or at our email:

- 🗨️ [Discord server](https://discord.gg/gZqdRXbbqg)
- ✉️ [Email](mailto:dima@vuelto.me)

### 🙌 Thanks To

A special thanks to:
- **Dimkauzh** for the initial idea and development.
- **ZakaHaceCosas** for the great improvements on top of vuelto.

Your support has helped make Vuelto even better! ❤️


---

<h2 style="text-align: center;">Made with ❤️ by the Vuelto Team </h2>
