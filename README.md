# yona-bot <img src="/assets/images/bot_icon.jpg" width="50">

---

A self-hosted discord bot for turn-based conversation, utilizing various external libraries which categorized into

- **Language Model**
  - OpenAI GPT4o
  - more in the future!...
- **Speech To Text**
  - OpenAI Whisper
- **Text To Speech**
  - VOICEVOX ENGINE (Japanese voice engine)
  - more in the future!...

### Features

- Turn-based conversation with AI on discord voice channel.

---

### Setup instruction

#### Install these in your system first

- [Docker](https://www.docker.com/get-started/)
- [ffmpeg](https://www.ffmpeg.org/download.html) (don't need if use docker)
- [dca](https://github.com/bwmarrin/dca/tree/master/cmd/dca) (don't need if use docker)
- [go](https://go.dev/doc/install) (don't need if use docker)

#### Deployment (local and cloud)

First and foremost, prepare your discord bot token from your discord bot

1. clone the repository
2. run `./scripts/setup_env` and complete the configuration process.
3. run `docker compose up -d`
4. done!

or

1. clone the repository
2. run `./scripts/setup_env` and complete the configuration process.e
3. run `go run main.go` (make sure you have installed those libraries)
4. done!

#### How to use it

1. make sure your bot is online and already in the discord server
2. use `/join` to join the voice channel
3. use `/listen` to make the bot listen for questions at once
4. wait for the answers!
5. go back to #3

---

#### Issues encountered and possible solution I can think of

- Conversation delay (might be able to solve by controlling length of answer and improving code quality)

#### Contributions ? I don't know man

- If anyone interest in contributing or would like to scold about my shit code please contact me via discord "yobu\_" 💀

#### Goal ? I don't know man
