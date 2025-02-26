import { AnilistSmallMedia } from "../entities/anilist_small_media";
import { AnilistPageInfo } from "../entities/anilist_page_info";
import { AnilistBigMedia } from "../entities/anilist_big_media";
import { AnilistAiringSchedule } from "../entities/anilist_airing_schedule";

export interface IRepository {
  fetchAnimesByName(
    name: string
  ): Promise<
    [AnilistSmallMedia[] | null, AnilistPageInfo | null, Error | null]
  >;
  fetchAnimeById(
    animeId: number
  ): Promise<[AnilistBigMedia | null, Error | null]>;
  fetchAnimesByNameAndPage(
    name: string,
    page: number
  ): Promise<
    [AnilistSmallMedia[] | null, AnilistPageInfo | null, Error | null]
  >;
  fetchSeasonalAnimes(
    season: string
  ): Promise<[AnilistSmallMedia[] | null, Error | null]>;
  fetchAiringSchedule(
    animeId: number
  ): Promise<[AnilistAiringSchedule | null, Error | null]>;
}
