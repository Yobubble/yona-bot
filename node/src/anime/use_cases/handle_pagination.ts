import { logger } from "../../utils/logger";
import { AnilistPageInfo } from "../entities/anilist_page_info";
import { AnilistSmallMedia } from "../entities/anilist_small_media";
import { IRepository } from "../repositories/irepository";

export class HandlePagination {
  private repo: IRepository;

  constructor(repo: IRepository) {
    this.repo = repo;
  }

  async execute(
    ctx: string,
    name: string,
    pageInfo: AnilistPageInfo
  ): Promise<
    [AnilistSmallMedia[] | null, AnilistPageInfo | null, Error | null]
  > {
    logger.trace("Call HandlePagination use case");
    let pageNum: number;

    if (ctx === "next" && pageInfo.hasNextPage) {
      pageNum = pageInfo.currentPage + 1;
    } else if (ctx === "prev" && pageInfo.currentPage > 1) {
      pageNum = pageInfo.currentPage - 1;
    } else {
      return [null, null, new Error("next or prev page not available")];
    }

    const [animes, page, error] = await this.repo.fetchAnimesByNameAndPage(
      name,
      pageNum
    );
    if (error) {
      return [null, null, error];
    }

    return [animes, page, null];
  }
}
