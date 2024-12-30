{
  description = "Vuelto development environment for NixOS";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs";

  outputs = { self, nixpkgs }:
    let
      pkgs = import nixpkgs { system = "x86_64-linux"; };
    in {
      devShell.x86_64-linux = pkgs.mkShell {
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

          pkgs.mesa
          pkgs.libglvnd
          pkgs.alsa-lib

          pkgs.go
          pkgs.gcc

          pkgs.gopls
          pkgs.wasmserve
        ];
      };
    };
}
