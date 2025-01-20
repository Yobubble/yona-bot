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
const services_1 = require("../enums/services");
class CommandService {
    constructor(appId, botToken, logger) {
        this.appId = appId;
        this.botToken = botToken;
        this.logger = logger;
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
                const data = yield res.json();
                this.logger.error(`Failed to initialize global commands: ${res.statusText} - ${data.message}`, services_1.Services.COMMAND_SERVICE);
                return;
            }
            this.logger.info("Global Commands initialized successfully", services_1.Services.COMMAND_SERVICE);
        });
    }
}
exports.CommandService = CommandService;
