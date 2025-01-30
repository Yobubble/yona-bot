import { logger } from "../../../utils/services/logger_service";
import { AnilistBigMedia } from "../entities/anilist_big_media";
import { IRepository } from "../repositories/irepository";

export class GetAnimeById {
  private repo: IRepository;

  constructor(repo: IRepository) {
    this.repo = repo;
  }

  async execute(
    animeId: number
  ): Promise<[AnilistBigMedia | null, Error | null]> {
    logger.trace("Call GetAnimeById use case");

    const [animes, error] = await this.repo.fetchAnimeById(animeId);
    if (error) {
      return [null, error];
    }

    return [animes, null];
  }
}
