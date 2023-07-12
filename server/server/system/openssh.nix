{ config, pkgs, ...}:

{
  # Enable the OpenSSH daemon.
  services.openssh = {
    enable = true;
    settings.PermitRootLogin = "no";
    settings.UseDns = false;
    ports = [ 53 ];
  };
}
