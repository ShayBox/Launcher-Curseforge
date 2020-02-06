# MultiMC-Twitch

A simple Go program that handles Twitch's custom protocol and ccip files
1. Reads the [CurseForge] `.ccip` file or [Twtch] `twitch://` protocol  
2. Requests the [TwitchAPI] to get the zip url  
3. Launches [MultiMC] with the `--import` flag, with the url  

Examples:
  - `MultiMC-Twitch example.ccip`
  - `MultiMC-Twitch twitch://www.curseforge.com/minecraft/modpacks/aesthetic-construction/download-client/2246179`

Includes:
  - Pre-configured Info.plist for handling Twitch protocol
  - Pre-configured multimc-twitch.desktop for handling Twitch protocol
  - Twitch.reg for handling Twitch protocol (Must be configured)

Configure Twitch.reg for Windows:
1. Edit `Twitch.reg` with any text editor  
2. Update the two paths to `MultiMC-Twitch.exe`   
  (Defaults to C:\MultiMC\MultiMC-Twitch.exe)
3. Run `Twitch.reg`  

[Download](https://github.com/ShayBox/MultiMC-CCIP/releases)

[CurseForge]: https://www.curseforge.com/
[Twitch]: https://twitch.tv/
[TwitchAPI]: https://twitchappapi.docs.apiary.io/
[MultiMC]: https://multimc.org/