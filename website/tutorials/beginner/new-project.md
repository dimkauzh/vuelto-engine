<!-- markdownlint-disable md010 -->
# 📂 Creating a new project

Hello and welcome to the first tutorial in the beginner series! In this tutorial, we will create a new project using Vuelto. Let's get started!

!!! note
    This tutorial assumes you have a basic understanding of Go. If you are new to Go, you can learn more about it.
    If you want to follow along, the example we are building is available in the [Github repo](https://github.com/vuelto-org/vuelto/blob/stable/examples/basic-window/main.go).

## 📦 Prerequisites

Before we start, make sure you have the following installed:

- 🖥️ A C compiler
- 🔧 A Go compiler (Go 1.18 and above)
- 🪟 Xorg/Wayland development packages (For Linux only)

Also make sure you're on a supported platform (Windows, macOS, or Linux), and you're targeting supported platforms too (Windows, macOS, Linux, and the web).

Without this, vuelto might not function. For a installation guide, [go here](../../get-started.md).

## 🚀 Creating a new project

First, create a new directory for your project. You can name it whatever you want. For the sake of this tutorial, we will name it `my-vuelto-project`.

On any shell, run:

```bash
mkdir my-vuelto-project
cd my-vuelto-project
```

Then init a new Go module:

```bash
go mod init my-vuelto-project
```

To get started lets create a new file called `main.go`. Here out game will be built in, and this can be expanded into multiple files later on.

On Linux/macOS:

```bash
touch main.go
```

On Windows (PowerShell):

```powershell
New-Item main.go -type file
```

Or, with Windows CMD

```cmd
type nul > main.go
```

Now open the file in your favorite text editor. To check if everything is setup correctly lets print "Hello, Vuelto!" to the console.
Put the following go code in the `main.go` file:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Vuelto!")
}
```

Now you can run the project by running:

```bash
go run main.go
```

This should print `Hello, Vuelto!` to the console. If it does, you have successfully created a new project! 🎉

## 📚 Setup vuelto

In the steps above, we have created a new project and printed "Hello, Vuelto!" to the console. Now we will setup vuelto in our project.

First, we need to get the vuelto package. Run the following command in your terminal:

```bash
go get vuelto.pp.ua@latest
```

This will download the latest version of vuelto and add it to your `go.mod` file.
Now that we have the package installed, we can start with using it in our project!

First up, import the vuelto package in your `main.go` file:

```go
// ...
import (
	// ..
	vuelto "vuelto.pp.ua/pkg"
	// ..
)
// ...
```

Next, we need to initialize Vuelto and create a new window. Add the following code to your `main.go` file:

```go
// ...
func main() {
	// ...
	win := vuelto.NewWindow("My Vuelto Window", 800, 600, false)
	// ...
}
// ...
```

This will create a window with the following properties:

| Property | Value |
| :--- | ---: |
| Window's title | `My Vuelto Window` |
| Window's width | `800` |
| Window's height | `600` |
| Is the window resizable? | `false` |

Now we need to keep the game running. This is done by creating a so called game loop. This loop will keep the game running until the window is closed. Add the following code to your `main.go` file:

```go
// ...
func main() {
	// ...
	for !win.Close() {
		// ...
		win.Refresh()
	}
	// ...
}
// ...
```

Now you can run the project (the same way as before, `go run main.go`) and you should see a window pop up with the title `My Vuelto Window`. If you do, you've successfully setup Vuelto in your project! 🎉
