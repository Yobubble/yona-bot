import { Response } from "express";
import { DiscordHelper } from "../../utils/helpers/discord_helper";
import { InteractionResponseType } from "discord-interactions";
import { logger } from "../../utils/logger";

export class VoicevoxChat {
  async establishVoiceConnection(res: Response, guildId: string, userId: string): Promise<void> {
    const [vs, err] = await DiscordHelper.getUserVoiceState(guildId, userId)
    //TODO: handle error and when vs voice channel id is null
    if (err != null) {
      logger.error("Failed to fetch user's voice state", err)
      res.sendStatus(500).send({
        type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
        data: {
          content: "internal server error",
        },
      })
    } 

    if (vs === null) {
      logger.error("Failed to fetch user's voice state", new Error("User is currently not in a voice channel"))
      res.sendStatus(400).send({
        type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
        data: {
          content: "internal server error",
        },
      }) 
    }


    return  
  }
}
