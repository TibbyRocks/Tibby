## Tibby
 
A Discord bot written in Go utilizing the DiscordGo library


### Documentation

As a user you can read the [command documentation](https://tibby.rocks/docs/) here.

If you are a developer or want to use Tibby yourself the rest of this README is for you.

### Requirements

 - Ideally a AMD64 or ARM64 Linux/Unix based system but we build a Windows binary (AMD64 only) as well
 - A CDN (content delivery network) for images
 - An API key for Microsoft Translations API (Cognitive Services)
 - Or an API key for Google Translation API v2

### Setup
If you want to run Tibby you can get the latest binary from the [releases page](https://github.com/TibbyRocks/Tibby/releases)

After which you'll need to setup 2 things:

Customization options in customizations/botproperties.json:
```json
{
    "BotName": "Bot Name, for interaction titles and logging",
    "DocsURL": "Documentation URL",
    "CDN": {
        "BaseURL": "CDN Base-URL",
        "Files": {
            "8ball-icon": "/bot/icons/8-ball.png"
        }
    }
}
```

Credentials in .env (can be overridden with environment variables)

```ini 
WB_DC_TOKEN="your-discord-bot-token"
WB_MS_TRANSLATE_KEY="microsoft-translation-api-key"
WB_MS_TRANSLATE_REGION="microsoft-translations-api-region"
GOOGLE_API_KEY="google-translations-api-key"
```


### License

    Tibby, a Discord bot.
    Copyright (C) 2024  Dylan Maassen van den Brink and Tibby contributors.

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>