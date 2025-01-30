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
Object.defineProperty(exports, "__esModule", { value: true });
exports.CommandService = void 0;
const logger_service_1 = require("./logger_service");
class CommandService {
    constructor(appId, botToken) {
        this.appId = appId;
        this.botToken = botToken;
    }
    initGlobalCommands(commands) {
        return __awaiter(this, void 0, void 0, function* () {
            const url = `https://discord.com/api/v10/applications/${this.appId}/commands`;
            const headers = {
                Authorization: `Bot ${this.botToken}`,
                "Content-Type": "application/json",
            };
            const res = yield fetch(url, {
                method: "PUT",
                headers: headers,
                body: JSON.stringify(commands),
            });
            if (!res.ok) {
                const error = yield res.json();
                logger_service_1.logger.error("Failed to initialize global commands:", new Error(`${res.statusText} ${error.message}`));
                return;
            }
            logger_service_1.logger.info("Global Commands initialized successfully");
        });
    }
}
exports.CommandService = CommandService;
