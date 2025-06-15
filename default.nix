{ stdenv, lib, path, pkgs, makeWrapper, libnotify }:

pkgs.buildGoModule {
  pname = "MyFeedNotification";
  version = "v1.0";

  src = path;
  nativeBuildInputs = [ makeWrapper ];

  postFixup = ''
    wrapProgram $out/bin/MyFeed \
      --prefix PATH : ${lib.makeBinPath [ libnotify ]}
  '';

  vendorHash = "sha256-zZLQ2laQTjB6bFfrhqtnFSWwJCd1Gb+0LmPc5a3rtro=";
}
