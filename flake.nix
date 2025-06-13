{
  inputs = { nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05"; };

  outputs = { self, nixpkgs }:
    let pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in {
      packages.x86_64-linux.my_feed_notification =
        pkgs.callPackage ./default.nix { path = "${self}"; };
      devShell.x86_64-linux = pkgs.callPackage ./default.nix {
        path = "${self}";
        dev-shell = true;
      };
      packages.x86_64-linux.default =
        self.packages.x86_64-linux.my_feed_notification;

    };
}
