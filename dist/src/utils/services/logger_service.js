"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.logger = exports.LoggerService = void 0;
const winston_1 = __importDefault(require("winston"));
class LoggerService {
    constructor() {
        this.logger = winston_1.default.createLogger({
            level: "debug",
            format: winston_1.default.format.combine(winston_1.default.format.colorize(), winston_1.default.format.timestamp({ format: "YYYY-MM-DD HH:mm:ss" }), winston_1.default.format.printf(({ timestamp, level, message }) => {
                return `${level} [${timestamp}]: ${message}`;
            })),
            transports: [new winston_1.default.transports.Console()],
        });
    }
    info(message) {
        this.logger.info(message);
    }
    error(message, error) {
        this.logger.error(`${message} - ${error.name}`);
    }
    data(message, obj) {
        this.logger.info(`${message} - ${JSON.stringify(obj)}`);
    }
    trace(message) {
        this.logger.debug(message);
    }
}
exports.LoggerService = LoggerService;
exports.logger = new LoggerService();
