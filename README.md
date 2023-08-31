<div align="center">
  <a href="https://discord.shaybox.com">
    <img alt="Discord" src="https://img.shields.io/discord/824865729445888041?color=404eed&label=Discord&logo=Discord&logoColor=FFFFFF">
  </a>
  <a href="https://github.com/shaybox/launcher-curseforge/releases/latest">
    <img alt="Downloads" src="https://img.shields.io/github/downloads/shaybox/launcher-curseforge/total?color=3fb950&label=Downloads&logo=github&logoColor=FFFFFF">
  </a>
</div>

# Launcher-Curseforge

Integrates the [Curseforge] [Minecraft] Modpack installation button to any [MultiMC] based [Minecraft] launcher.  
Handles the `curseforge://` custom protocol and executes the launcher with the `--import` argument.

## Supported Launchers:
- [PrismLauncher] [^1]
- [PolyMC]
- [MultiMC]

## Installation:

### Windows:
- [Download] and Extract the latest release
- Move the `.exe` into the same directory as the launcher
- Run as **Administrator** once to update registry values

### macOS:
- [Download] the latest release
- Move the `.app` into the `Applications` directory

### Linux:

#### Archlinux: [AUR]

#### Debian/Ubuntu:
- [Download] and Extract the latest release
- Run `sudo dpkg -i launcher-curseforge_X.X.X_amd64.deb`

#### Other:
- [Download] and Extract the latest release
- Extract the `.deb` package and the `data.tar.gz` inside
- Manually move the files to `~/.local`
- Run `xdg-mime default launcher-curseforge.desktop x-scheme-handler/curseforge`

[^1]: The last supported version of PrismLauncher is version 6, version 7 broke the `--import` flag, and version 8 integrated Launcher-Curseforge functionality

[PrismLauncher]: https://prismlauncher.org
[PolyMC]: https://polymc.org
[MultiMC]: https://multimc.org

[Curseforge]: https://curseforge.com
[Minecraft]: https://minecraft.net
[MultiMC]: https://multimc.org
[Download]: https://github.com/ShayBox/Launcher-Curseforge/releases/latest
[AUR]: https://aur.archlinux.org/packages/launcher-curseforge-bin
