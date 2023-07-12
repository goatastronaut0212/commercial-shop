{ pkgs, ... }:

{
  environment = {
    systemPackages = with pkgs; [
      # Lang
      nodejs
      go
      openjdk17-bootstrap

      # Screen
      screen

      git
      neofetch
      htop
      unzipNLS
      unrar
      pciutils
      pulsemixer
      xorg.xbacklight
      zip
      zstd
    ];

    variables = {
      BACKEND_PORT="4505"
      NIXOS_FLAKE_DIR="$HOME/flake";
      NVIM_DIR = "$HOME/flake/nvim";
    };
  };
}
