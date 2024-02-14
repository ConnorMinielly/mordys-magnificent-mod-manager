# Mordenkainen’s Magnificent Mod Manager

TODO: Rewrite this intro in the voice of Mordenkainen, because it's fun.

Okay here's the deal: Mod management for Baldur's Gate 3 on non-windows systems is a mess. The widely used/preferred mod manager is great, but because of fundamental limitations of the underlying dependencies (namely Norbyte's LSLib which is used for parsing and writing `.pak` files in the BG3 source, which is dependent on the archaic dotnet 4) it isn't cross platform and the path to getting it there is extremely muddy. We're looking at a real "foundations of sand" situation here. So the proposal of this project is "Do less, for more". This project is going to laser focus BG3, with no effort to support any of Larian's older games or engines, and we're going to dump the UI entirely. This is going to be a stripped down, elegant mod management tool for the release version of BG3 on linux/MacOS.

# Goals

**Modding Knowledge Base (Mordenkainen's Omnipotent Tome):**

- [ ] Document the functionality and purpose of a BG3 mod manager, and clarify what exact must be done for existing mods to work with the game.

**Mod Management CLI (Mordenkainen's Magnificent Mod Manager):**

- [ ] Implement config wizard for first time setup/updating settings. This will all get saved to a `json` config (possibly at a path like `~/.m4/config.json`):
  - [ ] Allow user to define path to BG3 install locations (for default try to auto discover BG3 install locations)
  - [ ] Allow user to provide a mod source folder (default at `~/m4-mods`?).
  - [ ] Allow user to set location to export mod load order (default at `~/.m4/load-order.json`)
- [ ] Initially, for proof of concept, mods will be added by manually updating the `load-order.json` by providing entries to the mods array (see `load-order.json` spec below)
- [ ] Provide commit/save function that will export mods to the game install as per the `load-order.json` and update the modsettings.lsx`
- [ ] Provide a scrub mods function, that will remove mods from install and reset `modsettings.lsx`.
- [ ] Provide manual export function to dump mods and `modsettings.lsx` to a specified directory instead of directly to the game install.
- [ ] Support exporting a portable cross platform executable binary.

**Mod Creation Toolchain (Mordenkainen's Magnificent Mod Maker):**

- [ ] Support reading `.pak` files by porting LSLib's PackageReader functions for BG3 to go.

# Initial `load-order.json` Spec

Load order is determined by the index of mod entry in the `mods` array (in descending order, so 0 index is first in the load order).

```json
{
  "modFixer": true, // We should package the mod fixer and allow users to activate it when needed with this simple flag.
  "mods": [
    {
      "name": "Aether's Black Dyes",
      "path": "M4_MOD_FOLDER/black-dye/aethers-black-dye.pak",
      "author": "Aetherpoint" // optional
    }
  ]
}
```

# FAQ With Mordenkainen 🧙🏼‍♂️

## Q: Wait, why not just use the existing tools again?

To put it simply, there are three major reasons why the current tooling isn't sufficient for a thriving BG3 modding community:

1. The core 3 tools for mod management and creation are built primarily on `.net` 4.8, a Microsoft C# and C++ framework that only functions on windows platforms. Because of this even though BG3 is available on linux (through steam's interop tools), MacOS, and Windows, the current modding suite supports Windows only leaving out a sizable chunk of players and creators who don't or can't use windows.
2. C# is a great language for developing enterprise applications, but it creates a ton of overhead when used for low level, highly focused applications like a mod tool. Like wise, C++ is fantastic for highly technical and intricate systems requiring high performance but is (in the opinion of the Lord Mage of Greyhawk) the wrong choice developing an open source project that intends to emphasize simplicity and accessibility. What takes a hundred lines and half a dozen files of code in C# can be expressed in just a few dozen lines in Go, and yet it remains more readable and accessible than most C++ code.
3. Norbyte and LaughngLeader and dozens of other developers have done truly phenomenal work in decoding the Larian Studios Divinity game engine and providing modding tools over the years. This cannot be overstated: this project doesn't exist without the mages who came before and deciphered the most arcane and obtuse of binary scrolls. However, this work has been ongoing since the original Divinity Original Sin, and as such the existing modding tools for BG3 have a lot of baggage from the previous games and versions of the Divinity game engine. BG3 has an exponentially larger popular interest than the DOS games and will likely have a long lasting fandom among gamers and modders due to the D&D 5E IP and BG3's unprecedented quality as an RPG. BG3 deserves a bespoke tool set, with built in awareness of it's particular quirks and content. For example, this mod tool set could eventually provide outputs for item tables and various UUIDs for everything from NPCs to in-world objects. If BG3 is the sole scope, It's actually possible to dream bigger and provide a more beautifully catered mod managing and creating experience.

## Q: Okay, so why not work on the existing open source tools and fix all these points? Or fork them if you want BG3 specific versions?

You can likely intuit the answer for this by reading the above answer, but to make it crystal clear: There are some fundamental changes that this tool chain endeavors to make to the structure of the modding ecosystem. I, the great wizard Mordenkainen, intend to bring the entire modding process under my domain: from unpacking original source files, to mod bundling, to end user mod management. I prefer to use the right wand for the job, and believe that using Go to build a non-graphical, CLI-first suite of tools will result in a more focused and accessible code base and provide enough other advantages to make re-enchanting the wheel worthwhile. Also, again, the existing tools mostly use .net 4.8 which is not sufficiently cross platform. Even if the work of upgrading to a more recent version of the .net framework is completed, the other philosophical differences in the structure and purpose of the projects would still stand.

## Q: Why is a CLI driven, non-graphical approach your priority?

I've thought carefully about what these tools need to be able to do. Every function is highly possible and ergonomic without the use of a full graphical UI. With that in mind, why engage with the massive additional overhead of building an interface? This project will focus on building a CLI with Go first and foremost, and will possibly expand into an interactive terminal UI to meet less technical users in the middle. If predictable `json` formatted outputs are provided for every operation, this toolset could be used as a backend by other wizards who want to build something flashy on top. That's up to them though, the mighty wizard Mordenkainen shall focus on functionality and portability, leave it to Elminster, Tasha, or some other arcane charlatan to implement the frivolities.

## Q: Hey Mordenkainen, aren't you from the Greyhawk setting? And isn't BG3 set in The Forgotten Realms and the land of Faerûn?

Well well well, think you clever eh? If you know that much, then you'd also know that I've traveled throughout the many realms and been to countless worlds. In fact, I've even been to your Earth! I'm a good friend of Elminster Aumer who got his grey beard all twisted up in the story of this game, and I was even I Avernus monitoring the state of the Blood War during the events that immediately preceded the story of BG3. For more on that, you can reference the earthling tome "Descent Into Avernus". All that to say: Nice try, but shut up Nerd!

# Prior Art and References

- [LSLib](https://github.com/Norbyte/lslib) by Norbyte
- [bg3se (Baldur's Gate 3 script extender)](https://github.com/Norbyte/bg3se) by Norbyte
- [BG3 Mod Manager](https://github.com/LaughingLeader/BG3ModManager) by LaughingLeader
- [BG3 Modding Tools](https://github.com/LaughingLeader/BG3ModdingTools/tree/master) by LaughingLeader
- [baldurs gate 3 mod manager](https://github.com/mkinfrared/baldurs-gate3-mod-manager) by mkinfrared (attempts to support MacOS)
- [BG3-Modders-Multitool](https://github.com/ShinyHobo/BG3-Modders-Multitool) by ShinyHobo
- [Awesome BG3](https://github.com/bg3mods/awesome-bg3)
- [BG3 Community Library](https://github.com/BG3-Community-Library-Team/BG3-Community-Library)
- [Norbyte's Resource Search Engine](https://bg3.norbyte.dev/search) by Norbyte

# Where Mod Files Go In Linux

PAK files - C:\Users\ronno\AppData\Local\Larian Studios\Baldur's Gate 3\Mods\
Mod Files - C:\Users\ronno\AppData\Local\Larian Studios\Baldur's Gate 3\PlayerProfiles\Public\modsettings.lsx

File path linux reference - /home/deck/.steam/steam/steamapps/compatdata/1086940/pfx/drive_c/users/steamuser/AppData/Local/Larian Studios/Baldur's Gate 3/

![mods-folder](https://github.com/ConnorMinielly/mordys-magnificent-mod-manager/assets/25215145/1daddc4b-7e54-453e-8302-d73be8f3ff4f)
