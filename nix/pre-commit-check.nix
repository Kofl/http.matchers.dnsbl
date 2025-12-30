# SPDX-FileCopyrightText: 2025 Gergely Nagy
# SPDX-FileContributor: Gergely Nagy
#
# SPDX-License-Identifier: EUPL-1.2

{ pkgs, treefmt, ... }:

{
  treefmt = {
    enable = true;
    always_run = true;
    package = treefmt;
  };
  reuse = {
    enable = true;
    name = "reuse";
    description = "Run REUSE compliance tests";
    entry = "${pkgs.reuse}/bin/reuse lint";
    pass_filenames = false;
    always_run = true;
  };
}
