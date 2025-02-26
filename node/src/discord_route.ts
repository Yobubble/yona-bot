import { Request, Response } from "express";
import { logger } from "./utils/logger";
import { AnimeSearch } from "./anime/controllers/anime_search";
import { InteractionResponseType, InteractionType } from "discord-interactions";
import { animeSearch, test, voicevoxChat } from "./utils/constants/commands";
import { ComponentType } from "./utils/enums/component_types";
import { FeatureTags } from "./utils/enums/feature_tags";
import { VoicevoxChat } from "./llm/controllers/voicevox_chat";

const animeSearchController = new AnimeSearch();
const voicevoxChatController = new VoicevoxChat();

export async function discordRoute(req: Request, res: Response) {
  // Interaction Object: https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
  const { type, data, guild_id, user } = req.body;

  if (type === InteractionType.APPLICATION_COMMAND) {
    const { name } = data;
    logger.trace(`Received command /${name}`);

    switch (name) {
      case test.name:
        res.send({
          type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
          data: {
            content: "Hello, world!",
          },
        });
        break;

      case animeSearch.name:
        let animeName = data.options[0].value;
        await animeSearchController.getFirstPage(animeName, res);
        break;

      case voicevoxChat.name:
        // TODO: join target voice channel
        await voicevoxChatController.establishVoiceConnection(res, guild_id, user.id);
        break;
      default:
        logger.error(`Unknown command name:`, name);
        res.status(400).json({ error: "Unknown command name" });
        break;
    }
    return;
  }

  if (type === InteractionType.MESSAGE_COMPONENT) {
    const { custom_id, component_type } = data;
    logger.data("Received interaction with:", { custom_id, component_type });

    // NOTE: Index definition
    // 0: FeatureTag
    // For Anime Search: 1: context / anime_id , 2: number of current page, 3: has next page or not, 4: anime name

    const inputs = custom_id.split(" ").filter((input: string) => input !== "");

    const featureTag = inputs[0];

    switch (component_type) {
      case ComponentType.BUTTON:
        if (featureTag === FeatureTags.ANIME_SEARCH) {
          const ctxOrAnimeId = inputs[1];
          const animeName = inputs[4];
          const currentPage = inputs[2];
          const hasNextPage = inputs[3];

          if (inputs.length === 2) {
            await animeSearchController.getAnimePage(ctxOrAnimeId, res);
          } else if (inputs.length === 5) {
            await animeSearchController.pagination(
              currentPage,
              hasNextPage,
              animeName,
              ctxOrAnimeId,
              res
            );
          }
        }
        break;

      default:
        logger.error("Unknown component type", component_type);
        res.status(400).json({ error: "Unknown component type" });
        break;
    }
    return;
  }

  logger.error("Unknown interaction type", type);
  res.status(400).send({
    type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
    data: {
      content: "Unknown interaction type",
    },
  });
  return;
}
