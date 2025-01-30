import { logger } from "../../../utils/services/logger_service";
import { AnilistPageInfo } from "../entities/anilist_page_info";
import { AnilistSmallMedia } from "../entities/anilist_small_media";
import { IRepository } from "../repositories/irepository";

export class GetAnimeByName {
  private repo: IRepository;

  constructor(repo: IRepository) {
    this.repo = repo;
  }

  async execute(
    name: string
  ): Promise<
    [AnilistSmallMedia[] | null, AnilistPageInfo | null, Error | null]
  > {
    logger.trace("Call GetAnimeByName use case");

    const [animes, page, error] = await this.repo.fetchAnimesByName(name);
    if (error) {
      return [null, null, error];
    }
    return [animes, page, null];
  }
}
