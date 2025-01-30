import { CommandType } from "../types/command_type";

export const testCommand: CommandType = {
  name: "test",
  description: "A test command",
  type: 1,
  integration_types: [0, 1],
  contexts: [0, 1, 2],
};

export const animeSearchCommand: CommandType = {
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

export const seasonalAnimeAiringScheduleCommand: CommandType = {
  name: "seasonal_anime_airing_schedule",
  description: "Get the airing schedule of seasonal anime",
  type: 1,
};

export const allCommands: CommandType[] = [
  testCommand,
  animeSearchCommand,
  seasonalAnimeAiringScheduleCommand,
];
