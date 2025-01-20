"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.LoggerService = void 0;
const winston_1 = __importDefault(require("winston"));
class LoggerService {
    constructor() {
        this.logger = winston_1.default.createLogger({
            format: winston_1.default.format.combine(winston_1.default.format.colorize(), winston_1.default.format.timestamp(), winston_1.default.format.simple()),
            transports: [new winston_1.default.transports.Console()],
        });
    }
    info(message, name) {
        this.logger.info(message, { name });
    }
    error(message, name) {
        this.logger.error(message, { name });
    }
}
exports.LoggerService = LoggerService;
