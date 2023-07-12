{ config, pkgs, ... }:

{
  boot = {
    # Boot loader
    loader.grub = {
      enable = true;
      device = "/dev/sda";
    };

    kernel.sysctl = {
      "net.ipv4.conf.all.forwarding" = true;
    };
  };
}
