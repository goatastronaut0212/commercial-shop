{ pkgs, ... }:

{
  home.stateVersion = "22.11";
  home.username = "pepsiman";
  home.homeDirectory = "/home/pepsiman";

  imports = [
    ./neovim.nix
  ];
}
