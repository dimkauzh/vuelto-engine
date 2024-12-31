# 📅 Upcoming release
## V1.1

- [ ] Renderer
    - [ ] SetPixel
    - [ ] LoadText and .Draw()
    - [ ] Move the main API to the Rendering driver
    - [x] Web
        - [x] Web-windowing
        - [x] Support for web using wasm
        - [x] Web-rendering
        - [x] Support for building graphics driver
        - [x] WebGL
    - [x] Desktop
      - [x] Move to OpenGL v3.3

- [x] Windowing
  - [x] Platforms
    - [x] Windowing for web
    - [x] Windowing for desktop
  - [ ] Framerates
    - [ ] Set framerate
    - [ ] Get framerate
    - [ ] Deltatime

- [x] Imaging
    - [x] Load images
    - [x] Process images (to string)

- [x] Events
    - [x] Keyboard
    - [x] Mouse

- [x] 2 Ring Rendering Driver
    - [x] Ring 1
        - [x] Wrapper around OpenGL C api
        - [x] Wrapper around WebGL JS api
    - [x] Ring 2
        - [x] API around OpenGL and WebGL

- [x] 2 Ring Window Driver
    - [x] Ring 1
        - [x] Wrapper around GLFW
        - [x] Wrapper around JS runtime
    - [x] Ring 2
        - [x] API around GLFW and JS windowing

- [x] 2 Ring Input Driver
    - [x] Ring 1 (No ring 1 needed)
    - [x] Ring 2
        - [x] API around GLFW and JS events

- [x] 2 Ring Image Driver
    - [x] Ring 1 (No ring 1 needed)
    - [x] Ring 2
      - [x] API around Go image library

- [ ] Website
    - [ ] Docs
        - [ ] Vuelto Docs
            - [ ] Vuelto API
            - [ ] How to use vuelto
            - [ ] Vuelto's functionality
        - [x] Developer Docs
            - [x] Vuelto's 3 ring structure
            - [x] CONTRIBUTING.md
            - [x] Code of Conduct
    - [ ] Tutorials
        - [ ] Examples
        - [ ] Tutorials to build small games
