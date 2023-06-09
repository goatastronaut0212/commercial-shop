{ pkgs, ... }:

{
  users.users.nixuser = {
    isNormalUser = true;
    description = "nixuser";
    extraGroups = [ "audio" "networkmanager" "video" "postgres" "wheel" ];
    packages = with pkgs; [
      firefox
      lynx
      qutebrowser
      w3m
    ];
  };
}
