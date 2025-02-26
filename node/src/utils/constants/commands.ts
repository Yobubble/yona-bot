import { CommandType } from "../types/command_type";

// Application Command Types: https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
// Interaction Context Types: https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-context-types
export const test: CommandType = {
  name: "test",
  description: "A test command",
  type: 1,
  integration_types: [0, 1],
  contexts: [0, 1, 2],
};

export const animeSearch: CommandType = {
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
    // TODO: more options
  ],
  integration_types: [0, 1],
  contexts: [0, 1, 2],
};

export const seasonalAnimeAiringSchedule: CommandType = {
  name: "seasonal_anime_airing_schedule",
  description: "Get the airing schedule of seasonal anime",
  type: 1,
};

export const voicevoxChat: CommandType = {
  name: "voicevox_chat",
  description: "A chatting simulation backed by voicevox and chatGPT",
  type: 1,
  contexts: [0]
}

export const allCommands: CommandType[] = [
  test,
  animeSearch,
  seasonalAnimeAiringSchedule,
  voicevoxChat
];
