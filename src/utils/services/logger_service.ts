import winston from "winston";

export class LoggerService {
  private logger: winston.Logger;

  constructor() {
    this.logger = winston.createLogger({
      format: winston.format.combine(
        winston.format.colorize(),
        winston.format.timestamp(),
        winston.format.simple()
      ),
      transports: [new winston.transports.Console()],
    });
  }

  public info(message: string, name: string): void {
    this.logger.info(message, { name });
  }

  public error(message: string, name: string): void {
    this.logger.error(message, { name });
  }
}
