# SPDX-FileCopyrightText: 2025 Gergely Nagy
# SPDX-FileContributor: Gergely Nagy
#
# SPDX-License-Identifier: EUPL-1.2
{
  description = "Caddy module match against DNS blocklists";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.11";

    pre-commit-hooks = {
      url = "github:cachix/pre-commit-hooks.nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.nixpkgs-stable.follows = "nixpkgs";
    };
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      systems,
      treefmt-nix,
      ...
    }@inputs:
    let
      inherit (nixpkgs) lib;

      forEachSystem =
        f:
        nixpkgs.lib.genAttrs (import systems) (
          system:
          let
            pkgs = import nixpkgs {
              inherit system;
            };
          in
          f pkgs
        );

      treefmtEval = forEachSystem (pkgs: treefmt-nix.lib.evalModule pkgs ./nix/treefmt.nix);
      treefmtWrapper = pkgs: treefmtEval.${pkgs.system}.config.build.wrapper;
    in
    {
      formatter = forEachSystem treefmtWrapper;
      checks = forEachSystem (pkgs: {
        formatting = treefmtEval.${pkgs.system}.config.build.check self;
        pre-commit-check = inputs.pre-commit-hooks.lib.${pkgs.system}.run {
          src = ./.;
          hooks = import ./nix/pre-commit-check.nix {
            inherit pkgs;
            treefmt = treefmtWrapper pkgs;
          };
        };
      });

      devShells = forEachSystem (pkgs: {
        default = pkgs.mkShell {
          buildInputs = self.checks.${pkgs.system}.pre-commit-check.enabledPackages;
          packages = with pkgs; [
            go
            reuse
            xcaddy
          ];
          inputsFrom = [
            treefmtEval.${pkgs.system}.config.build.devShell
          ];
        };
      });
    };
}
