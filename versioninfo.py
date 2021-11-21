import sys
import json
import semver

from datetime import datetime

if __name__ == '__main__':
    now = datetime.now()

    if len(sys.argv) < 1:
        print("You must provide the semantic version as the first argument!")
        raise SystemExit(1)

    version_str = sys.argv[1]
    version_str = version_str.lstrip("v")
    print("Received Semantic Version:", version_str)
    try:
        ver = semver.VersionInfo.parse(version_str)
    except ValueError as e:
        print("ERROR! -", e)
        print("WARNING! - Falling back to a generated develop release version")
        ver = semver.VersionInfo.parse(f"0.0.0-{version_str}+{int(now.timestamp())}")

    version_info = {
        "FixedFileInfo": {
            "FileVersion": {
                "Major": ver.major,
                "Minor": ver.minor,
                "Patch": ver.patch or 0,
                "Build": int(ver.build or 0)
            },
            "ProductVersion": {
                "Major": ver.major,
                "Minor": ver.minor,
                "Patch": ver.patch or 0,
                "Build": int(ver.build or 0)
            },
            "FileFlagsMask": "3f",
            "FileFlags ": "00",
            "FileOS": "040004",
            "FileType": "01",
            "FileSubType": "00"
        },
        "StringFileInfo": {
            "Comments": "go-socks5-cli - A CLI wrapper around `theriverman/go-socks5`",
            "CompanyName": "theriverman",
            "FileDescription": "go-socks5-cli",
            "FileVersion": str(ver),
            "InternalName": "gos5cli",
            "LegalCopyright": f"github.com/theriverman - All Rights Reserved. 2021 - {now.year}",
            "LegalTrademarks": "MIT License",
            "OriginalFilename": "socks5-cli.exe",
            "PrivateBuild": str(ver),
            "ProductName": "Go SOCKS5 CLI",
            "ProductVersion": str(ver),
            "SpecialBuild": ver.prerelease or "",
        },
        "VarFileInfo": {
            "Translation": {
                "LangID": "0409",
                "CharsetID": "04B0"
            }
        },
        "ManifestPath": ""
    }

    with open('versioninfo.json', 'w') as f:
        json.dump(version_info, f, indent=4)
