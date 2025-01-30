"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.allCommands = exports.seasonalAnimeAiringScheduleCommand = exports.animeSearchCommand = exports.testCommand = void 0;
exports.testCommand = {
    name: "test",
    description: "A test command",
    type: 1,
    integration_types: [0, 1],
    contexts: [0, 1, 2],
};
exports.animeSearchCommand = {
    name: "search_anime",
    description: "Get information about an anime",
    type: 1,
    options: [
        {
            name: "by_name",
            description: "Find an anime from its name",
            type: 3,
            required: true,
        },
        // NOTE: may add more search keyword in the future
    ],
    integration_types: [0, 1],
    contexts: [0, 1, 2],
};
exports.seasonalAnimeAiringScheduleCommand = {
    name: "seasonal_anime_airing_schedule",
    description: "Get the airing schedule of seasonal anime",
    type: 1,
};
exports.allCommands = [
    exports.testCommand,
    exports.animeSearchCommand,
    exports.seasonalAnimeAiringScheduleCommand,
];
