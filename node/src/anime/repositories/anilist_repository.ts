import { AnilistQueries } from "../../utils/constants/anilist_queries";
import { logger } from "../../utils/logger";
import { AnilistAiringSchedule } from "../entities/anilist_airing_schedule";
import { AnilistBigMedia } from "../entities/anilist_big_media";
import { AnilistPageInfo } from "../entities/anilist_page_info";
import { AnilistSmallMedia } from "../entities/anilist_small_media";
import { IRepository } from "./irepository";

export class AnilistRepository implements IRepository {
  fetchSeasonalAnimes(
    season: string
  ): Promise<[AnilistSmallMedia[] | null, Error | null]> {
    // TODO
    throw new Error("Method not implemented.");
  }
  fetchAiringSchedule(
    animeId: number
  ): Promise<[AnilistAiringSchedule | null, Error | null]> {
    // TODO
    throw new Error("Method not implemented.");
  }
  private static url = "https://graphql.anilist.co";

  async fetchAnimesByNameAndPage(
    name: string,
    page: number
  ): Promise<
    [AnilistSmallMedia[] | null, AnilistPageInfo | null, Error | null]
  > {
    logger.trace("Fetching animes by name and page");

    const res = await fetch(AnilistRepository.url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        query: AnilistQueries.fetchAnimesByNameAndPage,
        variables: {
          page: page,
          perPage: 3,
          search: name,
        },
      }),
    });

    if (!res.ok) {
      return [null, null, new Error(res.statusText)];
    }

    const data = await res.json();
    const pageInfo = data.data.Page.pageInfo as AnilistPageInfo;
    const medias = data.data.Page.media as AnilistSmallMedia[];

    return [medias, pageInfo, null];
  }

  async fetchAnimesByName(
    name: string
  ): Promise<
    [AnilistSmallMedia[] | null, AnilistPageInfo | null, Error | null]
  > {
    logger.trace("Fetching animes by name");

    const res = await fetch(AnilistRepository.url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        query: AnilistQueries.fetchAnimesByName,
        variables: {
          perPage: 3,
          search: name,
        },
      }),
    });

    if (!res.ok) {
      return [null, null, new Error(res.statusText)];
    }

    const data = await res.json();
    const pageInfo = data.data.Page.pageInfo as AnilistPageInfo;
    const medias = data.data.Page.media as AnilistSmallMedia[];

    return [medias, pageInfo, null];
  }

  async fetchAnimeById(
    animeId: number
  ): Promise<[AnilistBigMedia | null, Error | null]> {
    logger.trace("Fetching anime by id");

    const res = await fetch(AnilistRepository.url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        query: AnilistQueries.fetchAnimeById,
        variables: {
          mediaId: animeId,
        },
      }),
    });

    if (!res.ok) {
      return [null, new Error(res.statusText)];
    }

    const data = await res.json();
    const medias = data.data.Media as AnilistBigMedia;

    return [medias, null];
  }
}
