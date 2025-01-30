import { AnilistRepository } from "../repositories/anilist_repository";
import { IRepository } from "../repositories/irepository";
import { GetAnimeById } from "./get_anime_by_id";
import { GetAnimeByName } from "./get_animes_by_name";
import { GetBestRatingAnime } from "./get_best_rating_anime";
import { GetEmbedFields } from "./get_embed_fields";
import { HandlePagination } from "./handle_pagination";

export class AnimeUseCases {
  private static repo: IRepository = new AnilistRepository();

  static getAnimesByNameUseCase() {
    return new GetAnimeByName(this.repo);
  }

  static handlePaginationUseCase() {
    return new HandlePagination(this.repo);
  }

  static getAnimesByIdUseCase() {
    return new GetAnimeById(this.repo);
  }

  static getBestRatingAnimeUseCase() {
    return new GetBestRatingAnime();
  }

  static getEmbedFieldsUseCase() {
    return new GetEmbedFields();
  }
}
