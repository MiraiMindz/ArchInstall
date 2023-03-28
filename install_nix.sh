#!/bin/sh

# This script installs the Nix package manager on your system by
# downloading a binary distribution and running its installer script
# (which in turn creates and populates /nix).

{ # Prevent execution if this script was only partially downloaded
oops() {
    echo "$0:" "$@" >&2
    exit 1
}

umask 0022

tmpDir="$(mktemp -d -t nix-binary-tarball-unpack.XXXXXXXXXX || \
          oops "Can't create temporary directory for downloading the Nix binary tarball")"
cleanup() {
    rm -rf "$tmpDir"
}
trap cleanup EXIT INT QUIT TERM

require_util() {
    command -v "$1" > /dev/null 2>&1 ||
        oops "you do not have '$1' installed, which I need to $2"
}

case "$(uname -s).$(uname -m)" in
    Linux.x86_64)
        hash=afa6bf82bacb18ec0d6a1fd992b718192f895023d6cea8b94b2fb35fce352ff5
        path=1sh0llzrvkah17hypza89fcd5bmm35cl/nix-2.13.3-x86_64-linux.tar.xz
        system=x86_64-linux
        ;;
    Linux.i?86)
        hash=579a79a155c93c07f23e58de1dfa766d905239315b6b8c116e0ad7ffb7b14135
        path=bg4vn61787q0l6wny81laa7al4sw5hda/nix-2.13.3-i686-linux.tar.xz
        system=i686-linux
        ;;
    Linux.aarch64)
        hash=dfbbaa86404c3efe74d87c8d4c39081affc673c8b5f64cfefc4ab0eaf84ed39b
        path=zdq06xpx6a4a1hkaqcvjj9lbd7vqd9c3/nix-2.13.3-aarch64-linux.tar.xz
        system=aarch64-linux
        ;;
    Linux.armv6l)
        hash=debcb8a09744e83fca06efc6c5eee3f6066653da27dda8e0fb861e565b93f5f5
        path=60gh3w8ikkhf547d3scyhsn7xlpp1k01/nix-2.13.3-armv6l-linux.tar.xz
        system=armv6l-linux
        ;;
    Linux.armv7l)
        hash=07a1e718e9ebfd11b0c8ba947029b5dc50a81f4e6f8ac5b8b8e3572e1ce436cb
        path=cr6g8xc66yc02avw3ih4h96dbamqc3xw/nix-2.13.3-armv7l-linux.tar.xz
        system=armv7l-linux
        ;;
    Darwin.x86_64)
        hash=2b036faeccdce4b6ba9ca4aa7971b49d13a1ee4ac13ea834fb24c4d0d7e8536a
        path=s8yj3dnxmkk8ggzzymxwm1qb3s5hk133/nix-2.13.3-x86_64-darwin.tar.xz
        system=x86_64-darwin
        ;;
    Darwin.arm64|Darwin.aarch64)
        hash=200e777b541ddbadafdbe3817dba504754fa8fb0c9d1dd3f33dee6578045ab39
        path=nh14dgrnxgd9dgg14bvbmi33dd5ipnml/nix-2.13.3-aarch64-darwin.tar.xz
        system=aarch64-darwin
        ;;
    *) oops "sorry, there is no binary distribution of Nix for your platform";;
esac

# Use this command-line option to fetch the tarballs using nar-serve or Cachix
if [ "${1:-}" = "--tarball-url-prefix" ]; then
    if [ -z "${2:-}" ]; then
        oops "missing argument for --tarball-url-prefix"
    fi
    url=${2}/${path}
    shift 2
else
    url=https://releases.nixos.org/nix/nix-2.13.3/nix-2.13.3-$system.tar.xz
fi

tarball=$tmpDir/nix-2.13.3-$system.tar.xz

require_util tar "unpack the binary tarball"
if [ "$(uname -s)" != "Darwin" ]; then
    require_util xz "unpack the binary tarball"
fi

if command -v curl > /dev/null 2>&1; then
    fetch() { curl --fail -L "$1" -o "$2"; }
elif command -v wget > /dev/null 2>&1; then
    fetch() { wget "$1" -O "$2"; }
else
    oops "you don't have wget or curl installed, which I need to download the binary tarball"
fi

echo "downloading Nix 2.13.3 binary tarball for $system from '$url' to '$tmpDir'..."
fetch "$url" "$tarball" || oops "failed to download '$url'"

if command -v sha256sum > /dev/null 2>&1; then
    hash2="$(sha256sum -b "$tarball" | cut -c1-64)"
elif command -v shasum > /dev/null 2>&1; then
    hash2="$(shasum -a 256 -b "$tarball" | cut -c1-64)"
elif command -v openssl > /dev/null 2>&1; then
    hash2="$(openssl dgst -r -sha256 "$tarball" | cut -c1-64)"
else
    oops "cannot verify the SHA-256 hash of '$url'; you need one of 'shasum', 'sha256sum', or 'openssl'"
fi

if [ "$hash" != "$hash2" ]; then
    oops "SHA-256 hash mismatch in '$url'; expected $hash, got $hash2"
fi

unpack=$tmpDir/unpack
mkdir -p "$unpack"
tar -xJf "$tarball" -C "$unpack" || oops "failed to unpack '$url'"

script=$(echo "$unpack"/*/install)

[ -e "$script" ] || oops "installation script is missing from the binary tarball!"
export INVOKED_FROM_INSTALL_IN=1
"$script" "$@"

} # End of wrapping
