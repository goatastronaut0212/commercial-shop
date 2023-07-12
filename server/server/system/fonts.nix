{ config, pkgs, ... }:

{
  fonts.fonts = with pkgs; [
    corefonts
    (nerdfonts.override { fonts = [ "JetBrainsMono" "Terminus" ]; })
    noto-fonts
    noto-fonts-cjk
    noto-fonts-emoji
  ];
}
