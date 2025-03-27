# Branded

A Fire Emblem-inspired tactical RPG with D&D-style combat.
_Built with passion, caffeine, and just a little bit of frustration._

### What is This?
Ever wanted to play Fire Emblem but with dice rolls? Welcome to Branded! Here, your attacks aren't just predetermined numbers—they're at the mercy of the RNG gods. Every swing of your sword, every dodge, every spell... it all depends on a roll of the dice.

**Features**:

- 🎲 **D&D-style combat** – To-hit rolls, AC, and damage dice!

- 🏹 **No weapon restrictions** – Want to wield a battle axe as a mage? Go ahead.

- ☠️ **Permadeath**? Maybe. – You’ll find out.

- 🍀 **An evil clover exists** – Don’t ask. Just fear it.

### How to Play?

First thing first, take a quick look at the [basic control](doc/CONTROL.md).

Then you can read more about general gameplay in [the Handbook](doc/HANDBOOK.md).

### Building from Source

1. Clone the repo:

```sh
git clone https://github.com/baolhq/branded.git  
cd branded  
```

2. Build (assuming you have Go installed):

```sh
go build -o branded
```

3. Run the game:

```sh
./branded
```

### Project Structure

```sh
/Branded
├── main.go                 # Entry point
├── go.mod                  # Go module file
├── go.sum                  # Dependencies lock file
├── /lib                    # Game source code
│   ├── ai/                 # Monster logic, pathfinding, etc.
│   ├── core/               # Core game logic (map, entities, combat, etc.)
│   ├── data/               # Loading game data (items, monsters, etc.)
│   ├── env/                # World generation, procedural content
│   ├── meta/               # Metadata, constants, etc.
│   ├── ui/                 # Rendering, input handling, TUI
│   ├── util/               # Helper functions
│   └── world/              # Environment variables
├── /dist                   # Build output
├── /doc                    # Documentation
│   ├── CHANGELOG.md        # Rendering, input handling, TUI
│   ├── TRANSLATION.md      # Translation guide
│   ├── CONTROL.md          # Game control guide
│   ├── HANDBOOK.md         # The player handbook
│   └── STORY.md            # Game story
├── /lang                   # Translations
├── /res                    # Game resources
│   ├── audio/              # Sound effects
│   ├── font/               # Recommended fonts
│   ├── img/                # Images and banners
├── /script                 # Build/automation scripts
│   ├── build-windows.ps1   # Automated build script for Windows
│   ├── build-linux.sh      # Automated build script for Linux
├── /test                   # Unit and integration tests
│   ├── game_test.go        # Test for core game logic
│   ├── ui_test.go          # Test for UI handling
│   └── ...                 # More test files
```

### License

This project is issued under the [GPLv3](LICENSE) License.

### Contributing

Wanna help? Great! Fork the repo, make your changes, and send a pull request.

Just don’t break the game more than it already is.

For more information and tips, please read the [CONTRIBUTING](CONTRIBUTING.md).

### Acknowledgments

Thanks to D&D and Fire Emblem for existing.

Special thanks to random number generators for ensuring players suffer.

And, of course, you, for reading this far.