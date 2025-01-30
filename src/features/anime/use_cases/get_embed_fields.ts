import { MessageComponentTypes } from "discord-interactions";
import { FeatureTags } from "../../../utils/enums/feature_tags";
import { AnilistSmallMedia } from "../entities/anilist_small_media";
import { AnimeUseCases } from "./anime_use_cases";

export class GetEmbedFields {
  execute(animes: AnilistSmallMedia[]) {
    const fields = animes.map((anime: AnilistSmallMedia) => {
      return {
        name: "▸ " + anime.title.romaji,
        value: anime.genres.join(" | "),
        inline: false,
      };
    });

    fields.push({
      name: "\u200B",
      value: `💫  Anime with the highest score in the list: **${
        AnimeUseCases.getBestRatingAnimeUseCase().execute(animes).title.romaji
      }**`,
      inline: false,
    });

    const fieldsBtn = animes!.map((anime, idx) => {
      return {
        type: MessageComponentTypes.BUTTON,
        label: idx + 1,
        style: 2,
        custom_id: `${FeatureTags.ANIME_SEARCH} ${anime.id}`,
      };
    });

    return [fields, fieldsBtn];
  }
}
