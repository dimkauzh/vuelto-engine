# 🔧 Installation

This guide assumes you have the latest version of Golang installed on your system, if not, [refer to this page and follow the instructions for your platform](https://go.dev/dl/).

## 📦 Go package

We have a Go package published, so run this command to add Vuelto to your `go.mod`:

```bash
go get vuelto.pp.ua@latest
```

## 🐧 Linux

You need to have X11-dev-packages and C compiler installed.

On Debian and derivate distributions like Ubuntu and Linux Mint the xorg-dev meta-package pulls in the development packages for all of X11.

```bash
sudo apt install xorg-dev libasound2-dev gcc
```

On Fedora and derivatives like Red Hat the X11 extension packages libXcursor-devel, libXi-devel, libXinerama-devel and libXrandr-devel required by GLFW pull in all its other dependencies.

```bash
sudo dnf install libXcursor-devel libXi-devel libXinerama-devel libXrandr-devel alsa-lib-devel gcc
```

## 🍎 Mac

You need to have Xcode's command line tools installed.

```bash
xcode-select --install
```

## 🪟 Windows

You need to have a C compiler installed. You can also use WSL, if so follow [the Linux instructions](#-linux). We recommend MSYS2 (with MingW-w64), which you can get from [msys2.org](https://www.msys2.org/#installation), and we'll use this compiler for this manual. [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) and [Cygwin](https://cygwin.com/) are also tested against Golang and Fyne, and should work properly with Vuelto.

In Windows, your graphics driver should be already installed, however we still recommend to ensure it's up to date.

The steps for installing with MSYS2 are as follows:

1. Install MSYS2 from [msys2.org](https://www.msys2.org/#installation)
2. Once installed do not use the MSYS terminal that opens
3. Open "MSYS2 MinGW 64-bit" from the start menu

Execute the following commands (if asked for install options be sure to choose "all"):

```bash
pacman -Syu
pacman -S git mingw-w64-x86_64-toolchain
```

You will need to add `/c/Program\ Files/Go/bin` and `~/Go/bin` to your `$PATH`, for MSYS2 you can paste the following command into your terminal:

```bash
echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc
```

For the compiler to work on other terminals you will need to set up the Windows `%PATH%`. Search up "Edit the system environment variables" in Windows search, click "Environment Variables...", then edit `Path` for either your user or the entire system (select it then click "Edit...") and add "`C:\msys64\mingw64\bin`" (or wherever you installed MSYS32 / MINGW64) to the list (click "New" then paste that).
