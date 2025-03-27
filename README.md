# Branded

A turn-based strategy game inspired by classic Fire Emblem titles and D&D-style combat.

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

### Project Structure

```sh
/Branded
├── main.go                 # Entry point
├── go.mod                  # Go module file
├── go.sum                  # Dependencies lock file
├── /lib                    # Game source code
│   ├── bot/                # Enemy logics, pathfinding, etc.
│   ├── core/               # Core game logic (map, entities, combat, etc.)
│   ├── data/               # Loading game data (items, monsters, etc.)
│   ├── env/                # Environment variables
│   ├── gen/                # World generation, procedural content
│   ├── meta/               # Metadata, constants, etc.
│   ├── ui/                 # Rendering, input handling, TUI
│   └── util/               # Helper functions
├── /res                    # Static resources
│   ├── audio/              # Sound effects
│   ├── font/               # Recommended fonts
│   ├── img/                # Images and banners
├── /dist                   # Build output
├── /doc                    # Documentations, handbooks, etc.
├── /lang                   # Translations
├── /script                 # Build/automation scripts
├── /test                   # Unit and integration tests
│   ├── game_test.go        # Test for core game logic
│   └── ...                 # More test files
```

### License

This project is issued under the [GPLv3](LICENSE) License.

### Contributing

Wanna help? Great! Fork the repo, make your changes, and send a pull request.

Just don’t break the game more than it already is.

For more information and tips, please follow the [Code of Conduct](CODE_OF_CONDUCT.md) as well as the [Contributing guide](doc/CONTRIBUTING.md).

### Acknowledgments

Thanks to D&D and Fire Emblem for existing.

Special thanks to random number generators for ensuring players suffer.

And, of course, you, for reading this far.
