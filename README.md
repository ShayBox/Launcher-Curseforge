# PolyMC-Curseforge

A simple Go program that handles Curseforge's custom protocol and ccip files
1. Reads the [CurseForge] `.ccip` file or `curseforge://` protocol  
2. Requests the [CurseForge] to get the zip url  
3. Launches [PolyMC] with the `--import` flag, with the url  
4. Downloads the modpack icon into the icons folder

Instructions:
  - macOS - Move `PolyMC-Curseforge.app` into `Applications`
  - Linux - [AUR] or Manually install files into system
  - Windows - Move `PolyMC-Curseforge.exe` into `PolyMC` folder and execute

Note: Having the Curse app installed may break this.

[Download](https://github.com/ShayBox/PolyMC-Curseforge/releases)

I do not support MultiMC and its developers anymore  
If you would like to know more, check out [PolyMC]

[CurseForge]: https://www.curseforge.com/
[PolyMC]: https://polymc.org/
[AUR]: https://aur.archlinux.org/packages/Polymc-curseforge/