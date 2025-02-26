import { CommandType } from "./types/command_type";
import { logger } from "./logger";

export class DiscordCommand {
  private appId: string;
  private botToken: string;

  constructor(appId: string, botToken: string) {
    this.appId = appId;
    this.botToken = botToken;
  }

  public async initGlobalCommands(commands: CommandType[]): Promise<void> {
    const url = `https://discord.com/api/v10/applications/${this.appId}/commands`;
    const headers = {
      Authorization: `Bot ${this.botToken}`,
      "Content-Type": "application/json",
    };
    const res = await fetch(url, {
      method: "PUT",
      headers: headers,
      body: JSON.stringify(commands),
    });

    if (!res.ok) {
      const error = await res.json();
      logger.error(
        "Failed to initialize global commands:",
        new Error(`${res.statusText} ${error.message}`)
      );
      return;
    }

    logger.info("Global Commands initialized successfully");
  }
}
