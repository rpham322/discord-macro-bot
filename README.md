# Discord Macro Bot

A Discord bot written in Go that provides nutrition information using the Nutritionix API. Get instant macro and nutrition facts for any food item directly in your Discord server.

## Features

- 🍎 **Nutrition Lookup**: Query nutrition information for any food item using the `!macro` command
- 📊 **Detailed Macros**: Displays calories, protein, carbs, fat, fiber, and sugars
- 🎨 **Rich Embeds**: Beautiful Discord embed formatting for easy-to-read nutrition facts
- ⚡ **Fast Responses**: Quick API integration with error handling
- 🔍 **Natural Language**: Supports natural language food queries

## Prerequisites

- Go 1.25.0 or higher
- A Discord Bot Token ([Create a Discord Application](https://discord.com/developers/applications))
- Nutritionix API credentials ([Get Nutritionix API Key](https://www.nutritionix.com/business/api))

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/discord-macro-bot.git
cd discord-macro-bot
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory:
```env
DISCORD_BOT_TOKEN=your_discord_bot_token_here
NUTRITIONIX_APP_ID=your_nutritionix_app_id_here
NUTRITIONIX_TOKEN=your_nutritionix_token_here
```

4. Build the bot:
```bash
go build -o discord-macro-bot
```

5. Run the bot:
```bash
./discord-macro-bot
```

Or run directly with Go:
```bash
go run main.go
```

## Configuration

The bot requires the following environment variables:

| Variable | Description | Required |
|----------|-------------|----------|
| `DISCORD_BOT_TOKEN` | Your Discord bot token | Yes |
| `NUTRITIONIX_APP_ID` | Your Nutritionix Application ID | Yes |
| `NUTRITIONIX_TOKEN` | Your Nutritionix API Token | Yes |

These can be set in a `.env` file (recommended) or as system environment variables.

## Usage

### Commands

- **`!macro <food>`** - Get nutrition information for a food item
  - Example: `!macro 1 cup of rice`
  - Example: `!macro chicken breast`
  - Example: `!macro 2 slices of pizza`

### Other Triggers

- **`nutrition`** - Bot responds with a help message
- **`bot`** - Bot confirms its presence

## Query Best Practices & Limitations

### What Works Well

The Nutritionix API works best with:

- **Common Western foods**: Pizza, burgers, chicken breast, rice, pasta, etc.
- **Branded products**: "Coca Cola", "McDonald's Big Mac", "Oreo cookies"
- **Standard measurements**: Include serving sizes for better accuracy
  - `!macro 1 cup of rice`
  - `!macro 200g chicken breast`
  - `!macro 2 slices of bread`
- **Restaurant items**: Many chain restaurant items are in the database
  - `!macro Chipotle burrito bowl`
  - `!macro Starbucks grande latte`

### Known Limitations

⚠️ **Asian and International Foods**: The Nutritionix database has limited coverage for:
- Traditional Asian dishes (kimchi, pho, dim sum, etc.)
- Regional or ethnic specialty foods
- Homemade or custom recipes
- Less common international cuisines

### Tips for Better Results

1. **Be Specific**: Include serving sizes and measurements
   - ✅ `!macro 1 cup cooked white rice`
   - ❌ `!macro rice`

2. **Use Common Names**: Try English/common names instead of native names
   - ✅ `!macro steamed dumplings`
   - ❌ `!macro xiaolongbao`

3. **Break Down Complex Dishes**: Query individual components
   - Instead of: `!macro pad thai`
   - Try: `!macro rice noodles` and `!macro shrimp` separately

4. **Try Generic Descriptions**: Use descriptive terms
   - ✅ `!macro stir fried vegetables`
   - ✅ `!macro grilled chicken`
   - ✅ `!macro steamed rice`

5. **Include Brand Names**: If you know the brand, include it
   - ✅ `!macro Trader Joe's kimchi`
   - ❌ `!macro kimchi`

### Workarounds for Missing Foods

If a food isn't found in the database, consider:
- Searching for similar/common alternatives
- Breaking down the dish into individual ingredients
- Using generic descriptions (e.g., "stir fried chicken" instead of "kung pao chicken")

## Example Output

When you use `!macro chicken breast`, the bot will respond with an embed showing:
- Food name and brand (if available)
- Serving size
- Calories
- Protein, Carbs, Fat
- Fiber and Sugars

## Project Structure

```
discord-macro-bot/
├── bot/
│   ├── bot.go              # Main bot logic and message handlers
│   └── command-macro.go    # Nutrition API integration
├── main.go                 # Entry point and configuration
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
└── README.md              # This file
```

## Dependencies

- [discordgo](https://github.com/bwmarrin/discordgo) - Discord API wrapper for Go
- [godotenv](https://github.com/joho/godotenv) - Environment variable management

## Development

To contribute to this project:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source and available under the MIT License.

## Support

If you encounter any issues or have questions, please open an issue on the GitHub repository.

