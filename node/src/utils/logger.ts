import winston from "winston";

export class CustomLogger {
  private logger: winston.Logger;

  constructor() {
    this.logger = winston.createLogger({
      level: "debug",
      format: winston.format.combine(
        winston.format.colorize(),
        winston.format.timestamp({ format: "YYYY-MM-DD HH:mm:ss" }),
        winston.format.printf(({ timestamp, level, message }) => {
          return `${level} [${timestamp}]: ${message}`;
        })
      ),
      transports: [new winston.transports.Console()],
    });
  }

  public info(message: string): void {
    this.logger.info(message);
  }

  public error(message: string, error: Error): void {
    this.logger.error(`${message} - ${error.name}`);
  }

  public data(message: string, obj: object | string | number): void {
    this.logger.info(`${message} - ${JSON.stringify(obj)}`);
  }

  public trace(message: string): void {
    this.logger.debug(message);
  }
}

export const logger = new CustomLogger();
