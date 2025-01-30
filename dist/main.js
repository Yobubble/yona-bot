"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const command_service_1 = require("./src/utils/services/command_service");
require("dotenv/config");
const commands_1 = require("./src/utils/constants/commands");
const discord_interactions_1 = require("discord-interactions");
const discord_route_1 = require("./src/discord_route");
const app = (0, express_1.default)();
const port = 3000;
const appId = process.env.APP_ID;
const botToken = process.env.BOT_TOKEN;
const publicKey = process.env.PUBLIC_KEY;
new command_service_1.CommandService(appId, botToken).initGlobalCommands(commands_1.allCommands);
app.post("/interactions", (0, discord_interactions_1.verifyKeyMiddleware)(publicKey), discord_route_1.discordRoute);
app.listen(port, () => {
    console.log(`Server is running at http://localhost:${port}`);
});
