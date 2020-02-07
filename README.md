# MultiMC-Twitch

A simple Go program that handles Twitch's custom protocol and ccip files
1. Reads the [CurseForge] `.ccip` file or [Twitch] `twitch://` protocol  
2. Requests the [TwitchAPI] to get the zip url  
3. Launches [MultiMC] with the `--import` flag, with the url  

Instructions:
  - macOS - Move `MultiMC-Twitch.app` into `Applications`
  - Linux - [AUR] or Manually install files into system
  - Windows - Move `MultiMC-Twitch.exe` into `MultiMC` folder and execute

Note: Having the Twitch app installed may break this.

[Download](https://github.com/ShayBox/MultiMC-Twitch/releases)

[CurseForge]: https://www.curseforge.com/
[Twitch]: https://twitch.tv/
[TwitchAPI]: https://twitchappapi.docs.apiary.io/
[MultiMC]: https://multimc.org/
[AUR]: https://aur.archlinux.org/packages/multimc-twitch/