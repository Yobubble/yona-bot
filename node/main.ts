import express from "express";
import { DiscordCommand } from "./src/utils/discord_command";
import "dotenv/config";
import { allCommands } from "./src/utils/constants/commands";
import { verifyKeyMiddleware } from "discord-interactions";
import { discordRoute } from "./src/discord_route";

const app = express();
const port = 3000;

const appId = process.env.APP_ID as string;
const botToken = process.env.BOT_TOKEN as string;
const publicKey = process.env.PUBLIC_KEY as string;

new DiscordCommand(appId, botToken).initGlobalCommands(allCommands);

app.post("/interactions", verifyKeyMiddleware(publicKey), discordRoute);

app.listen(port, () => {
  console.log(`Server is running at http://localhost:${port}`);
});
