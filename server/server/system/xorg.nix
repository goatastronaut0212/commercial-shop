{ config, pkgs, ... }:

{
  services.xserver = {
    # Keymap in X11
    layout = "us";
    xkbVariant = "";

    # Enable the X11 window system
    enable = true;

    videoDrivers = [ "modesetting" ];

    # touchpad support
    libinput.enable = true;

    displayManager.startx.enable = true;
    desktopManager.mate.enable=true;
  };
}
