# Translation Guide for Branded

Welcome, language wizards! This guide will help you translate **Branded** into different languages, ensuring players worldwide can enjoy the game in their native tongue. Let's make every word count!

### 1. How to Translate

##### File Structure
All translation files live in the `/lang` directory. Each language has its own cozy folder named after its [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code.

Example:
```
/lang
  ├── en/      # English (default)
  ├── vi/      # Vietnamese
  ├── fr/      # French
  └── es/      # Spanish
```
Each folder contains JSON files for different parts of the game (e.g., `ui.json`, `items.json`, `dialogue.json`).

##### Editing Translation Files
Each translation file follows this simple structure:
```json
{
  "key": "Translated text here",
  "example_key": "Example translation"
}
```
- **Do NOT change the keys**. They’re the backbone of the game’s text system.
- Keep translations grammatically correct and natural—no robotic translations, please! 🤖🚫

##### Adding a New Language
1. Create a new folder inside `/lang` using the correct language code.
2. Copy the English (`en/`) files into the new folder.
3. Translate everything while keeping the JSON structure intact.
4. Test your translation in-game to ensure it looks good (and doesn't break things). 

### 2. Best Practices

- **Keep It Consistent**: Stick to the same terminology, especially for key game mechanics and item names.
- **Context Matters**: Some words mean different things in different situations. Read carefully before translating!
- **Respect Formatting**:
  - Keep placeholders (`{player_name}`, `{hp}`, etc.) unchanged.
  - Make sure text fits inside UI elements—nobody likes text cut in half.
- **Make It Sound Natural**: If it feels weird to say out loud, it probably needs tweaking.
- **Avoid Machine Translation**: Google Translate is great for quick checks, but it won’t capture nuance. Humans do it better! 🧠

### 3. Pro Tips & Tricks

- **Check for Updates**: We may add new strings over time, so stay tuned!
- **Proofread & Test**: Always review and test translations in-game to avoid awkward phrasing.
- **Teamwork Wins**: If you're working with others, use Git or another version control system to track changes smoothly.
- **Ask When in Doubt**: If an English string seems unclear, don’t guess—ask for clarification!

Thanks for helping make **Branded** multilingual! Your work ensures that players everywhere can enjoy the game in their language. Now go forth and translate like a legend!