{ stdenv, lib, path, pkgs, libnotify }:

pkgs.buildGoModule {
  pname = "MyFeedNotification";
  version = "v1.0";

  src = path;
  depsHostTarget = [ libnotify ];

  vendorHash = "sha256-zZLQ2laQTjB6bFfrhqtnFSWwJCd1Gb+0LmPc5a3rtro=";
}
