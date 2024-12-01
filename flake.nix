{
  description = "ZK frontend dev environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";
  };

  outputs = {nixpkgs, ...}: let
    # system should match the system you are running on
    system = "x86_64-linux";
  in {
    devShells."${system}".default = let
      pkgs = import nixpkgs {
        inherit system;
      };
    in
      pkgs.mkShell {
        # create an environment with nodejs_18, pnpm, and yarn
        packages = with pkgs; [
          go
          gopls
          gotools
          delve
          air
          watchexec
          templ
        ];

        shellHook = ''
          export GOPATH=$HOME/go
          export PATH=$GOPATH/bin:$PATH
          # Print versions
          echo "Go version: $(go version)"
          echo "Development environment ready!"
          exec zsh
        '';
      };
  };
}
