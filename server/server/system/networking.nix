{
  networking = {
    enableIPv6 = true;
    hostName = "nixos";
    networkmanager.enable = true;
    firewall.enable = true;
    firewall.allowedTCPPorts = [ 53 14004 25565 4505 ];
    firewall.allowedUDPPorts = [ 53 14004 25565 4505 ];
  };
}
