// Voice State Object: https://discord.com/developers/docs/resources/voice#voice-state-object
export class DiscordVoiceState {
    constructor(
        // public sessionId: string,
        // public deaf: boolean,
        // public mute: boolean,
        // public selfDeaf: boolean,
        // public selfMute: boolean,
        // public selfStream: boolean,
        // public selfVideo: boolean,
        // public suppress: boolean,
        // public guildId?: string | null, // snowflake
        public channelId?: string | null, // snowflake
        // public userId?: string | null, // snowflake
        // public requestToSpeakTimestamp?: string | null, // ISO8601 timestamp
        // public member?: GuildMember
    ) {}
}
