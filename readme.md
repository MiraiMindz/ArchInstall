# Arch Linux Install Script Written in Go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/MiraiMindz/ArchInstall?style=flat-square)

Welcome to my project of creating a universal Arch Linux install script in Go \[CURRENTLY IN DEVELOPMENT\]

This script aims to fully automate the hard-labour of typing a lot of commands while maintaining the DIY approach of Arch Linux, in other words this script is a easy-to-use/idiot-proof installer; and it ensures some UNIX standards aswell (like hostname and username conventions and rules).

*Why I called it 'universal' you might ask?* This is because it aims to install Arch on the most common ways (and some uncommon ones), I've scrapped the wiki for "Installation Methods" to bundle them here, currently it has only 3 methods:
- PC (Your usual installation)
- Server (Is self-explanatory)
- Removable Medium (Install Arch on a removable medium like a USB stick and make the proper configurations)

## Notes

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

- `MOVE`:
    - ARGS:
        - `From`: The source to move from.
        - `To`: The destination.
- `COPYFILE`:
    - ARGS:
        - `From`: The source to copy from.
        - `To`: The destination.
- `COPYDIR`:
    - ARGS:
        - `From`: The source to copy from.
        - `To`: The destination.
- `RENAME`:
    - ARGS:
        - `Source`: The source name.
        - `NewName`: The source new name.
- `EXECUTE`:
    - ARGS:
        - `UseSudo`: Tell the parser to use Root privileges.
        - `CommandDirectory`: The directory to run the command.
        - `Source`: The Command itself.
        - `Args`: The list of argumets for the command.
- `REPLACELINE`:
    - ARGS:
        - `FileReplace`: The file to replace the line.
        - `SourceLine`: The line to replace.
        - `EditedLine`: The new content.
- `GITCLONE`:
    - ARGS:
        - `GitURL`: URL of the Git repository.
        - `GitDestination`: The destination of the Git repository.

## `customization.json` Example Structure

```json
{
    "Themes": [
        {
            "Name": "Catppuccin",
            "Location": "/Catppuccin",
            "Packages": {
                "SystemPackages": {
                    "Terminal": "/Terminals/Terminator",
                    "Shell": "/Shells/Terminator",
                    "DesktopEnvironment": null,
                    "WindowManager": "DesktopEnvironment/WindowManager/i3",
                    "DisplayManager": "DisplayManager/SDDM",
                    "Compositor": null
                },
                "CodeEditors": {
                    "neovim": "/CodeEditors/NeoVim",
                    "vscode": "null"
                },
                "CliTools": {
                    "bat": "/CliTools/Bat"
                },
                "UserApps": {
                    "Discord": "/UserApps/Discord"
                },
                "OtherApps": {
                    "": ""
                }
            },
            "Assets": {
                "Wallpaper": "Assets/Wallpapers/WindowManager/Wallpaper01.png",
                "LockscreenWallpaper": "Assets/Wallpapers/LockScreen/Wallpaper01.png",
                "UserProfilePicture": "Assets/ProfilePictures/ProfilePicture01.png",
                "SystemFont": null
            }
        },
        {
            "Name": "Nord",
            "Location": "/Nord",
            "Packages": {
                "SystemPackages": {
                    "Terminal": "/Terminals/Terminator",
                    "Shell": "/Shells/Terminator",
                    "DesktopEnvironment": null,
                    "WindowManager": "DesktopEnvironment/WindowManager/i3",
                    "DisplayManager": "DisplayManager/SDDM",
                    "Compositor": null
                },
                "CodeEditors": {
                    "neovim": "/CodeEditors/NeoVim",
                    "vscode": "null"
                },
                "CliTools": {
                    "bat": "/CliTools/Bat"
                },
                "UserApps": {
                    "Discord": "/UserApps/Discord"
                },
                "OtherApps": {
                    "": ""
                }
            },
            "Assets": {
                "Wallpaper": "Assets/Wallpapers/WindowManager/Wallpaper01.png",
                "LockscreenWallpaper": "Assets/Wallpapers/LockScreen/Wallpaper01.png",
                "UserProfilePicture": "Assets/ProfilePictures/ProfilePicture01.png",
                "SystemFont": null
            }
        }
    ]
}
```
