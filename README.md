# MultiMC-Curseforge

A simple Go program that handles Curseforge's custom protocol and ccip files
1. Reads the [CurseForge] `.ccip` file or `curseforge://` protocol  
2. Requests the [CurseForge] to get the zip url  
3. Launches [MultiMC] with the `--import` flag, with the url  

Instructions:
  - macOS - Move `MultiMC-Curseforge.app` into `Applications`
  - Linux - [AUR] or Manually install files into system
  - Windows - Move `MultiMC-Curseforge.exe` into `MultiMC` folder and execute

Note: Having the Curse app installed may break this.

[Download](https://github.com/ShayBox/MultiMC-Curseforge/releases)

[CurseForge]: https://www.curseforge.com/
[MultiMC]: https://multimc.org/
[AUR]: https://aur.archlinux.org/packages/multimc-curseforge/