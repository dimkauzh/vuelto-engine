{
  description = "Vuelto development environment for NixOS";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs";

  outputs = { self, nixpkgs }:
    let
      pkgs = import nixpkgs { system = "x86_64-linux"; };
    in {
      devShell.x86_64-linux = pkgs.mkShell {
        buildInputs = [
          pkgs.libX11
          pkgs.gcc
          pkgs.gnumake
          pkgs.go
        ];
      };
    };
}
