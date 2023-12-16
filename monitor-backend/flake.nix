{
  description = "this is the backend for a monitoring application";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin" ];

      perSystem = { config, self', inputs', pkgs, system, ... }:
        let
          name = "monitor-backend";
          version = "latest";
          vendorHash = "sha256-4DHILTsQOdGv47l+awkQfFpENfZD+xmsqMexq07OYQQ="; # update whenever go.mod changes
        in
        {
          devShells = {
            default = pkgs.mkShell {
              inputsFrom = [ self'.packages.default ];
            };
          };

          packages = {
            default = pkgs.buildGoModule {
              inherit name vendorHash;
              src = ./.;
              subPackages = [ "cmd/monitor-backend" ];
            };

            docker = pkgs.dockerTools.buildImage {
              inherit name;
              tag = version;
              config = {
                Cmd = "${self'.packages.default}/bin/${name}";
              };
            };
          };
        };
    };
}
