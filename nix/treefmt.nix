# SPDX-FileCopyrightText: 2025 Gergely Nagy
# SPDX-FileContributor: Gergely Nagy
#
# SPDX-License-Identifier: EUPL-1.2

_:

{
  config = {
    projectRootFile = "./flake.nix";
    programs = {
      nixfmt.enable = true; # nix
      statix.enable = true; # nix static analysis
      deadnix.enable = true; # find dead nix code
      gofmt.enable = true; # go
      taplo.enable = true; # toml
    };
    settings.formatter = {
      taplo.options = [
        "format"
        "--option"
        "indent_tables=true"
        "--option"
        "indent_entries=true"
      ];
    };
  };
}
