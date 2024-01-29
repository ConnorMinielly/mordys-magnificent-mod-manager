# Mordy's Magnificent Mod Manager

Okay here's the deal: Mod management for Baldur's Gate 3 on non-windows systems is a mess. The widely used/preferred mod manager is great, but because of fundamental limitations of the underlying dependencies (namely Norbyte's LSLib which is used for parsing and writing `.pak` files in the BG3 source, which is dependent on the archaic dotnet 4) it isn't cross platform and the path to getting it there is extremely muddy. We're looking at a real "foundations of sand" situation here. So the proposal of this project is "Do less, for more". This project is going to laser focus BG3, with no effort to support any of Larian's older games or engines, and we're going to dump the UI entirely. This is going to be a stripped down, elegant mod management tool for the release version of BG3 on linux/MacOS.

# Goals

Core:

- [ ] Support reading `.pak` mode files by porting LSLib's PackageReader functions for BG3 to go.
- [ ] Support exporting a portable cross platform binary than could be added to $PATH to make your life easy.

Mod Management:

- [ ] Auto discover BG3 install locations.
- [ ] Allow users to select campaign to apply mods against.
- [ ] Let users designate a mod source folder on first use (and change on later uses if desired)
- [ ] Allow users to "activate" downloaded mods in source folder, set mod load order with a terminal based UI.
- [ ] Export a `modsettings.lsx` file and mod source to game folder.
- [ ] Save and load mods and load order to json files (stretch goal: auto download missing mod files when importing from json)

# Approach

I'm going to try approaching this with some modern `go` because its highly portable, fairly low level while being easy to read and write, and most importantly it has the awesome Charm CLI tool set.

First we're going to rebuild the LSLib PakReader functionality for BG3 mod files (and ONLY BG3 mod files), and if this proves to be possible then and only then while we approach building the actual mod manager CLI.

# Prior Art

- [LSLib](https://github.com/Norbyte/lslib) by Norbyte
- [bg3se (Baldur's Gate 3 script extender)](https://github.com/Norbyte/bg3se) by Norbyte
- [BG3 Mod Manager](https://github.com/LaughingLeader/BG3ModManager) by LaughingLeader
- [baldurs gate 3 mod manager](https://github.com/mkinfrared/baldurs-gate3-mod-manager) by mkinfrared (attempts to support MacOS)
- [BG3-Modders-Multitool](https://github.com/ShinyHobo/BG3-Modders-Multitool) by ShinyHobo

# Where Mod Files Go

PAK files - C:\Users\ronno\AppData\Local\Larian Studios\Baldur's Gate 3\Mods\
Mod Files - C:\Users\ronno\AppData\Local\Larian Studios\Baldur's Gate 3\PlayerProfiles\Public\modsettings.lsx

File path linux reference - /home/deck/.steam/steam/steamapps/compatdata/1086940/pfx/drive_c/users/steamuser/AppData/Local/Larian Studios/Baldur's Gate 3/ 

![mods-folder](https://github.com/ConnorMinielly/mordys-magnificent-mod-manager/assets/25215145/1daddc4b-7e54-453e-8302-d73be8f3ff4f)
