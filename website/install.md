# 📦 Go package
We have a Go package published, so run this command to add it to your go.mod:
```bash
go get vuelto.me@latest
```

# 🔧 Installation

## 🐧 Linux
You have to have X11-dev-packages and C compiler installed.

#### 🖥️ X11
On Debian and derivates like Ubuntu and Linux Mint the xorg-dev meta-package pulls in the development packages for all of X11.
```bash
sudo apt install xorg-dev libasound2-dev gcc
```

On Fedora and derivatives like Red Hat the X11 extension packages libXcursor-devel, libXi-devel, libXinerama-devel and libXrandr-devel required by GLFW pull in all its other dependencies.
```bash
sudo dnf install libXcursor-devel libXi-devel libXinerama-devel libXrandr-devel alsa-lib-devel gcc
```

## 🍎 Mac
You have to have Xcode command line tools installed.
### 📝 Xcode
```bash
xcode-select --install
```

## 🪟 Windows
You have to have a C compiler installed. You can also use WSL, if so follow the Linux instructions.


Download Go from the download page and follow instructions. 

Install one of the available C compilers for windows, the following are tested with Go and Fyne:
- MSYS2 with MingW-w64 - msys2.org
- TDM-GCC - tdm-gcc.tdragon.net
- Cygwin - cygwin.com

In Windows your graphics driver will already be installed, but it is recommended to ensure they are up to date.

The steps for installing with MSYS2 (recommended) are as follows:

1. Install MSYS2 from msys2.org
2. Once installed do not use the MSYS terminal that opens
3. Open "MSYS2 MinGW 64-bit" from the start menu

Execute the following commands (if asked for install options be sure to choose "all"):
```bash
pacman -Syu
pacman -S git mingw-w64-x86_64-toolchain
```

You will need to add /c/Program\ Files/Go/bin and ~/Go/bin to your $PATH, for MSYS2 you can paste the following command into your terminal:
```bash
echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc
```
For the compiler to work on other terminals you will need to set up the windows %PATH% variable to find these tools. Go to the "Edit the system environment variables" control panel, tap "Advanced" and add "C:\msys64\mingw64\bin" to the Path list.


