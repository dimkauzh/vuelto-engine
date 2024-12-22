{
  pkgs ? import <nixpkgs> {},
}:


pkgs.mkShell {
  buildInputs = [
    pkgs.xorg.libX11
    pkgs.xorg.libXext
    pkgs.xorg.libXrandr
    pkgs.xorg.libXinerama
    pkgs.xorg.libXcursor
    pkgs.xorg.libXi
    pkgs.xorg.libXxf86vm

    pkgs.pkg-config
    pkgs.gnumake

    pkgs.glew
    pkgs.freeglut

    pkgs.go
    pkgs.gcc
  ];
}
