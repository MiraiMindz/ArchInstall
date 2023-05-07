# Arch Linux Install Script Written in Go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/MiraiMindz/ArchInstall?style=flat-square)

Welcome to my project of creating a universal Arch Linux install script in Go \[CURRENTLY IN DEVELOPMENT\]

This script aims to fully automate the hard-labour of typing a lot of commands while maintaining the DIY approach of Arch Linux, in other words this script is a easy-to-use/idiot-proof installer; and it ensures some UNIX standards aswell (like hostname and username conventions and rules).

*Why I called it 'universal' you might ask?* This is because it aims to install Arch on the most common ways (and some uncommon ones), I've scrapped the wiki for "Installation Methods" to bundle them here, currently it has only 3 methods:
- PC (Your usual installation)
- Server (Is self-explanatory)
- Removable Medium (Install Arch on a removable medium like a USB stick and make the proper configurations)

<!-- ## Notes

As of easy-of-use, the current script uses a JSON file to stores any temporary variable, and this file is shredded and deleted at the final of the installation, that being said, any information is stored in *plain text*.

I thought about hashing those infos (like passwords), but it will require the user to re-type the information everytime the script would use them, so be concerned about someone accessing your device disk (or RAM) during installation (tbh, if someone could do this you would have more problems than a leaked info).

## .dotfiles Folder Structure

```plaintext
.dotfiles/
├─┬── Catppuccin/
│ ├──┬── Terminals/
│ │  ├────── Terminator/
│ │  │       ├────── config
│ │  │       ├────── additional_packages.json
│ │  │       └────── install.json
│ │  └────── Alacritty/
│ │          ├────── config
│ │          ├────── additional_packages.json
│ │          └────── install.json
│ ├──┬── CodeEditors/
│ │  ├────── VisualStudioCode/
│ │  │       ├────── extensions.json
│ │  │       ├────── profiles.json
│ │  │       ├────── settings.json
│ │  │       └────── install.json
│ │  └────── NeoVim/
│ │          ├────── configs/
│ │          ├────── extensions.json
│ │          ├────── settings.json
│ │          └────── install.json
│ ├──┬── Shells/
│ │  ├────── BASH/
│ │  │       ├────── Aliases.json
│ │  │       ├────── Functions.json
│ │  │       ├────── Prompts.json
│ │  │       └────── install.json
│ │  └────── ZSH/
│ │          ├────── Aliases.json
│ │          ├────── Functions.json
│ │          ├────── Prompts.json
│ │          └────── install.json
│ ├──┬── CliTools/
│ │  ├────── Bat/
│ │  │       ├────── config
│ │  │       └────── install.json
│ │  ├────── CAVA/
│ │  │       ├────── config
│ │  │       └────── install.json
│ │  └────── ls/
│ │          ├────── config
│ │          └────── install.json
│ ├──┬── SystemPackages/
│ │  ├────── DesktopEnvironment/
│ │  │       ├────── KDE/
│ │  │       │       ├────── configs/
│ │  │       │       └────── install.json
│ │  │       └────── GNOME/
│ │  │               ├────── configs/
│ │  │               └────── install.json
│ │  ├────── WindowManager/
│ │  │       ├────── i3/
│ │  │       │       ├────── config
│ │  │       │       └────── install.json
│ │  │       └────── BSPWM/
│ │  │               ├────── config
│ │  │               └────── install.json
│ │  ├────── DisplayManager/
│ │  │       ├────── SDDM/
│ │  │       │       ├────── configs/
│ │  │       │       └────── install.json
│ │  │       └────── LXDM/
│ │  │               ├────── configs/
│ │  │               └────── install.json
│ │  └────── Compositor/
│ │          ├────── Picom/
│ │          │       ├────── config
│ │          │       └────── install.json
│ │          └────── XCOMP/
│ │                  ├────── config
│ │                  └────── install.json
│ ├──┬── UserApps/
│ │  └────── Discord/
│ │          ├────── CustomStyles.css
│ │          └────── install.json
│ ├──┬── Assets/
│ │  ├────── Wallpapers/
│ │  │       ├────── LockScreen/
│ │  │       │       ├────── Wallpaper01.png
│ │  │       │       └────── Wallpaper02.png
│ │  │       ├────── WindowManager/
│ │  │       │       ├────── Wallpaper01.png
│ │  │       │       └────── Wallpaper02.png
│ │  │       └────── DesktopManager/
│ │  │               ├────── Wallpaper01.png
│ │  │               └────── Wallpaper02.png
│ │  └────── ProfilePictures/
│ │          ├────── ProfilePicture01.png
│ │          └────── ProfilePicture02.png
│ └───── customization.json
│
└──── Nord/ ...
```

## `install.json` Example Structure

```json
{
    "Instructions": [
        {
            "Command": "MOVE",
            "from": "sorce",
            "to": "source"
        },
        {
            "Command": "COPY",
            "from": "sorce",
            "to": "source"
        },
        {
            "Command": "RENAME",
            "Source": "Name",
            "NewName": "NewName"
        },
        {
            "Command": "EXECUTE",
            "UseSudo": false,
            "Source": "CommandName",
            "Args": [
                "1",
                "2",
                "3"
            ]
        }
    ]
}
```

### List of Available Commands:

Notes:

If any of those paths starts with a `&` (Ampersand) it will replace the symbol with the current path (The path of the `install.json` file), otherwise it needs to be an absolute path.

If it starts with a `%` (percentage-sign) it will use the Theme root location (the one specified in the `Location` attribute inside of the `customization.json`)

The `#` means `parent-directory`, It needs to be at the beginning, you can have any `#` as you want, all of then will be treated at a separated parent.

so using `##/file.ext` will mean `Go up 2 parent directories to access file.ext`, or the Shell equivalent of `../../file.ext`

The `PathVars` attribute is a list of variables that will be expanded in any path.

The `UseAssets` is a list of indexes referents to the `Assets` element in the `customization.json`, it is already filled with the Theme absolute path.

The indexes starts at `0`, so if you want to use the second and the fourth item, you need to put `[1, 3]` in the array.

To use the selected assets, you need to specify a `$` (dollar sign) with the index of the element in the `UseAssets` array (it starts at `0`) too.

So in the following example:

```json
"CurrentTheme": "Catppuccin",
    "Themes": [
        {
            "Name": "Catppuccin",
            "Location": "/Catppuccin",
            "Assets": {
                "Wallpaper": "/Assets/Wallpapers/WindowManager/Wallpaper01.png",
                "LockscreenWallpaper": "/Assets/Wallpapers/LockScreen/Wallpaper01.png",
                "UserProfilePicture": "/Assets/ProfilePictures/ProfilePicture01.png",
                "Photo01": "/Assets/Photos/001.jpg",
                "SystemFont": null
            }
        }
    ]
```

```json
{
    "Command": "MOVE",
    "UseAssets": [1, 3],
    "From": "$1",
    "To": "%/OtherDirectory/"
}
```

It will use the second and the fourth Assets of the given theme, and will move the fourth element into a new location.

*This is because, the `$1`, is referent to the SECOND item in `UseAssets`, if we wanted the first item in `UseAssets`, we simply put `$0` instead of `$1`.*

*And the second item in `UserAssets` is the fourth item in the `Assets` map of the theme, and the first item is the second element in that map.*

Assuming that the Theme is on my $HOME directory, the above example will result in the following command:

```sh
mv /home/mirai/Themes/Catppuccin/Assets/Photos/001.jpg /home/mirai/Themes/Catppuccin/OtherDirectory/
```

- `MOVE`:
    - Moves a file from `From` to `To`.
    - ARGS:
        - `From`: The source to move from.
        - `To`: The destination.
    - EXAMPLES:
        - ```json
            {
                "Command": "MOVE",
                "UseAssets": [1, 3],
                "From": "$1",
                "To": "%/OtherDirectory/"
            }
            ```
        - ```json
            {
                "Command": "MOVE",
                "From": "/home/mirai/ConfigDir/config.conf",
                "To": "/home/mirai/config.conf"
            }
            ```
- `COPYFILE`:
    - Copy a file from `From` to `To`, if any of those starts with a `&` (Ampersand) it will replace the symbol with the current path (The path of the `install.json` file), otherwise all paths must be absolute.
    - ARGS:
        - `From`: The source to copy from.
        - `To`: The destination.
    - EXAMPLES:
        - ```json
            {
                "Command": "MOVE",
                "From": "%/ConfigDir/config.conf",
                "To": "%/config.conf"
            }
            ```
        - ```json
            {
                "Command": "MOVE",
                "From": "/home/mirai/ConfigDir/config.conf",
                "To": "/home/mirai/config.conf"
            }
            ```
- `COPYDIR`:
    - Copies a directory and it's contents from `From` to `To`, if any of those starts with a `&` (Ampersand) it will replace the symbol with the current path (The path of the `install.json` file), (You can specify only the parent dir if you want to.), otherwise all paths must be absolute.
    - ARGS:
        - `From`: The source to copy from.
        - `To`: The destination.
    - EXAMPLES:
        - ```json
            {
                "Command": "MOVE",
                "From": "%/ConfigDir/config.conf",
                "To": "%/config.conf"
            }
            ```
        - ```json
            {
                "Command": "MOVE",
                "From": "/home/mirai/ConfigDir/config.conf",
                "To": "/home/mirai/config.conf"
            }
            ```
- `RENAME`:
    - ARGS:
        - `Source`: The source name.
        - `NewName`: The source new name.
    - EXAMPLES:
        - ```json
            {
                "Command": "MOVE",
                "From": "%/ConfigDir/config.conf",
                "To": "%/config.conf"
            }
            ```
        - ```json
            {
                "Command": "MOVE",
                "From": "/home/mirai/ConfigDir/config.conf",
                "To": "/home/mirai/config.conf"
            }
            ```
- `EXECUTE`:
    - ARGS:
        - `UseSudo`: Tell the parser to use Root privileges.
        - `CommandDirectory`: The directory to run the command (Defaults to the `install.json` directory).
        - `Source`: The Command itself.
        - `Args`: The list of argumets for the command.
    - EXAMPLES:
        - ```json
            {
                "Command": "MOVE",
                "From": "%/ConfigDir/config.conf",
                "To": "%/config.conf"
            }
            ```
        - ```json
            {
                "Command": "MOVE",
                "From": "/home/mirai/ConfigDir/config.conf",
                "To": "/home/mirai/config.conf"
            }
            ```
- `REPLACELINE`:
    - ARGS:
        - `FileReplace`: The file to replace the line.
        - `SourceLine`: The line to replace.
        - `EditedLine`: The new content.
    - EXAMPLES:
        - ```json
            {
                "Command": "MOVE",
                "From": "%/ConfigDir/config.conf",
                "To": "%/config.conf"
            }
            ```
        - ```json
            {
                "Command": "MOVE",
                "From": "/home/mirai/ConfigDir/config.conf",
                "To": "/home/mirai/config.conf"
            }
            ```
- `GITCLONE`:
    - ARGS:
        - `GitURL`: URL of the Git repository.
        - `GitDestination`: The destination of the Git repository.
    - EXAMPLES:
        - ```json
            {
                "Command": "MOVE",
                "From": "%/ConfigDir/config.conf",
                "To": "%/config.conf"
            }
            ```
        - ```json
            {
                "Command": "MOVE",
                "From": "/home/mirai/ConfigDir/config.conf",
                "To": "/home/mirai/config.conf"
            }
            ```

## `customization.json` Example Structure

```json
{
    "CurrentTheme": "Catppuccin",
    "Themes": [
        {
            "Name": "Catppuccin",
            "Location": "/Catppuccin",
            "PackagesKeys": [
                "Terminals",
                "Shells",
                "CodeEditors",
                "CliTools",
                "SystemPackages",
                "UserApps"
            ], 
            "Packages": {
                "Terminals": "/Terminals/Terminator",
                "Shells": "/Shells/ZSH",
                "CodeEditors": {
                    "neovim": "/CodeEditors/NeoVim",
                    "vscode": "null"
                },
                "CliTools": {
                    "bat": "/CliTools/Bat"
                },
                "SystemPackages": {
                    "DesktopEnvironment": null,
                    "WindowManager": "/DesktopEnvironment/WindowManager/i3",
                    "DisplayManager": "/DisplayManager/SDDM",
                    "Compositor": null
                },
                "UserApps": {
                    "Discord": "/UserApps/Discord"
                }
            },
            "Assets": {
                "Wallpaper": "/Assets/Wallpapers/WindowManager/Wallpaper01.png",
                "LockscreenWallpaper": "/Assets/Wallpapers/LockScreen/Wallpaper01.png",
                "UserProfilePicture": "/Assets/ProfilePictures/ProfilePicture01.png",
                "Photo01": "/Assets/Photos/001.jpg",
                "SystemFont": null
            }
        },
        {
            "Name": "Nord",
            "Location": "/Nord",
            "PackagesKeys": [
                "Terminals",
                "Shells",
                "CodeEditors",
                "CliTools",
                "SystemPackages",
                "UserApps"
            ], 
            "Packages": {
                "Terminals": "/Terminals/Terminator",
                "Shells": "/Shells/ZSH",
                "CodeEditors": {
                    "neovim": "/CodeEditors/NeoVim",
                    "vscode": "null"
                },
                "CliTools": {
                    "bat": "/CliTools/Bat"
                },
                "SystemPackages": {
                    "DesktopEnvironment": null,
                    "WindowManager": "/DesktopEnvironment/WindowManager/i3",
                    "DisplayManager": "/DisplayManager/SDDM",
                    "Compositor": null
                },
                "UserApps": {
                    "Discord": "/UserApps/Discord"
                }
            },
            "Assets": {
                "Wallpaper": "/Assets/Wallpapers/WindowManager/Wallpaper01.png",
                "LockscreenWallpaper": "/Assets/Wallpapers/LockScreen/Wallpaper01.png",
                "UserProfilePicture": "/Assets/ProfilePictures/ProfilePicture01.png",
                "SystemFont": null
            }
        }
    ]
}
``` -->
