# yona-bot <img src="/assets/images/bot_icon.jpg" width="50">

---

A self-hosted discord bot for real-time conversation, utilizing various external libraries which categorized into

- **Language Model**
  - OpenAI GPT4o
  - more in the future!...
- **Speech To Text**
  - OpenAI Whisper
- **Text To Speech**
  - VOICEVOX ENGINE (Japanese voice engine)
  - more in the future!...

### Features

- Real-time conversation on discord voice channel.
- Conversation histories available for each discord server.
- Handle multiple questions at once.
- AWS S3 integration for cloud storage.
- .env configuration script.
- Leverage docker compose for simple deployment process.

> 📔 NOTES
>
> - Currently only support Japanese and GPT4o for language model.
> - you can just terminate the bot directly in the terminal to reveal the recorded files or use `/disconnect` to remove the recorded files

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
4. Done!

or

1. clone the repository
2. run `./scripts/setup_env` and complete the configuration process.e
3. run `go run main.go` (make sure you have installed those libraries)
4. Done!

---

#### Issues encountered and possible solution I can think of

- Conversation delay (might be able to solve by controlling length of answer and improving code quality)

#### Contributions ? I don't know man

- If anyone interest in contributing or would like to scold about my shit code please contact me via discord "yobu\_" 💀

#### Goal ? I don't know man
