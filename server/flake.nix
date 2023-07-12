{
  description = "My flake";

  inputs = {
    # Link to nix stable packages
    nixpkgs.url = "github:nixos/nixpkgs/nixos-23.05";

    # Link to nix unstable packages
    nixpkgs-unstable.url = "github:nixos/nixpkgs/nixos-unstable";

    # home-manager link packages
    home-manager = {
      url = "github:nix-community/home-manager/release-23.05";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = 
  { self, nixpkgs, nixpkgs-unstable, home-manager, ... }@inputs:
    let
      # Config settings
      username = "nixuser";
      myPlatform = "server";
      system = "x86_64-linux";

      pkgs = import nixpkgs { 
        inherit system;
        config.allowUnfree = true;
      };

      lib = nixpkgs.lib;

      overlay-unstable = final: prev: {
        unstable = import nixpkgs-unstable {
          inherit system;
          config.allowUnfree = true;
        };
      };
    in {
      nixosConfigurations = {
        nixos = lib.nixosSystem {
          inherit system;

          modules = [
            # Overlays-module makes "pkgs.unstable" available in configuration.nix
            ({ config, pkgs, ... }: {
              nixpkgs.overlays = [ 
                overlay-unstable
              ];
            })
            ./${myPlatform}/system/configuration.nix
            
            home-manager.nixosModules.home-manager {
              home-manager.useGlobalPkgs = true;
              home-manager.useUserPackages = true;
              #home-manager.extraSpecialArgs = { inherit inputs; };
              home-manager.users.${username} = {
                imports = [
                  ./${myPlatform}/home-manager/home-manager.nix
                ];
              };
            }
          ];
        };
      };
    };
}
