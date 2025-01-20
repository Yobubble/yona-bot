import { Services } from "../enums/services";
import { CommandType } from "../types/command_type";
import { LoggerService } from "./logger_service";

export class CommandService {
  private appId: string;
  private botToken: string;
  private logger: LoggerService;

  constructor(appId: string, botToken: string, logger: LoggerService) {
    this.appId = appId;
    this.botToken = botToken;
    this.logger = logger;
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
      const data = await res.json();
      this.logger.error(
        `Failed to initialize global commands: ${res.statusText} - ${data.message}`,
        Services.COMMAND_SERVICE
      );
      return;
    }

    this.logger.info(
      "Global Commands initialized successfully",
      Services.COMMAND_SERVICE
    );
  }
}
