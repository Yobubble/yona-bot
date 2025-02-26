import { DiscordVoiceState } from "../../llm/entities/discord_voice_state";
import { URL } from "../constants/url";

// Get User Voice State: https://discord.com/developers/docs/resources/voice#voice-state-object
export class DiscordHelper {
    static async getUserVoiceState(guildId: string, userId: string): Promise<[DiscordVoiceState | null, Error | null]> {
        const response = await fetch(`${URL.baseDiscordUrl}/guilds/${guildId}/voice-states/${userId}`)
        if (!response.ok) {
           return [null, new Error(response.statusText)] 
        }

        const data = await response.json();
        return [new DiscordVoiceState(data.channel_id), null];
    }
}