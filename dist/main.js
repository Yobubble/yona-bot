"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const logger_service_1 = require("./src/utils/services/logger_service");
const command_service_1 = require("./src/utils/services/command_service");
require("dotenv/config");
const commands_1 = require("./src/utils/constants/commands");
const discord_interactions_1 = require("discord-interactions");
const services_1 = require("./src/utils/enums/services");
const app = (0, express_1.default)();
const port = 3000;
const logger = new logger_service_1.LoggerService();
const appId = process.env.APP_ID;
const botToken = process.env.BOT_TOKEN;
const publicKey = process.env.PUBLIC_KEY;
const command = new command_service_1.CommandService(appId, botToken, logger);
command.initGlobalCommands(commands_1.allCommands);
app.post("/interactions", (0, discord_interactions_1.verifyKeyMiddleware)(publicKey), function (req, res) {
    return __awaiter(this, void 0, void 0, function* () {
        const { type, data } = req.body;
        if (type === discord_interactions_1.InteractionType.PING) {
            res.send({
                type: discord_interactions_1.InteractionResponseType.PONG,
            });
            return;
        }
        if (type === discord_interactions_1.InteractionType.APPLICATION_COMMAND) {
            const { name } = data;
            logger.info(`Received command /${name}`, services_1.Services.DISCORD_SERVICE);
            if (name === commands_1.testCommand.name) {
                res.send({
                    type: discord_interactions_1.InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
                    data: {
                        content: "Hello, world!",
                    },
                });
                return;
            }
            logger.error(`Unknown command name ${name}`, services_1.Services.DISCORD_SERVICE);
            res.status(400).json({ error: "Unknown command name" });
            return;
        }
        logger.error("Unknown interaction type", services_1.Services.DISCORD_SERVICE);
        res.status(400).json({ error: "Unknown interaction type" });
        return;
    });
});
app.listen(port, () => {
    console.log(`Server is running at http://localhost:${port}`);
});
