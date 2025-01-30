import { logger } from "../../../utils/services/logger_service";
import { AnilistSmallMedia } from "../entities/anilist_small_media";

export class GetBestRatingAnime {
  execute(animes: AnilistSmallMedia[]): AnilistSmallMedia {
    logger.trace("Call GetBestRatingAnime use case");

    const bestRatingAnime = animes.reduce((prev, current) => {
      return prev.averageScore > current.averageScore ? prev : current;
    });

    return bestRatingAnime;
  }
}
