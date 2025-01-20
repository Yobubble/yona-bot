import express from "express";
import { LoggerService } from "./src/utils/services/logger_service";
import { CommandService } from "./src/utils/services/command_service";
import "dotenv/config";
import { allCommands, testCommand } from "./src/utils/constants/commands";
import {
  InteractionType,
  InteractionResponseType,
  verifyKeyMiddleware,
} from "discord-interactions";
import { Services } from "./src/utils/enums/services";
import { Features } from "./src/utils/enums/features";

const app = express();
const port = 3000;

const logger = new LoggerService();

const appId = process.env.APP_ID as string;
const botToken = process.env.BOT_TOKEN as string;
const publicKey = process.env.PUBLIC_KEY as string;

const command = new CommandService(appId, botToken, logger);
command.initGlobalCommands(allCommands);

app.post(
  "/interactions",
  verifyKeyMiddleware(publicKey),
  async function (req, res) {
    const { type, data } = req.body;

    if (type === InteractionType.PING) {
      res.send({
        type: InteractionResponseType.PONG,
      });
      return;
    }

    if (type === InteractionType.APPLICATION_COMMAND) {
      const { name } = data;
      logger.info(`Received command /${name}`, Services.DISCORD_SERVICE);

      if (name === testCommand.name) {
        res.send({
          type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
          data: {
            content: "Hello, world!",
          },
        });
        return;
      }

      logger.error(`Unknown command name ${name}`, Services.DISCORD_SERVICE);
      res.status(400).json({ error: "Unknown command name" });
      return;
    }

    logger.error("Unknown interaction type", Services.DISCORD_SERVICE);
    res.status(400).json({ error: "Unknown interaction type" });
    return;
  }
);

app.listen(port, () => {
  console.log(`Server is running at http://localhost:${port}`);
});
