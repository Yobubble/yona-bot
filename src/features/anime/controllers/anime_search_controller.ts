import {
  InteractionResponseType,
  MessageComponentTypes,
} from "discord-interactions";
import { logger } from "../../../utils/services/logger_service";
import { Colors } from "../../../utils/constants/colors";
import { FeatureTags } from "../../../utils/enums/feature_tags";
import { AnimeUseCases } from "../use_cases/anime_use_cases";
import { Response } from "express";
import { MonthService } from "../../../utils/services/month_service";
import { TextService } from "../../../utils/services/text_service";

export class AnimeSearchController {
  async getAnimePage(animeId: number, res: Response): Promise<void> {
    logger.trace("Call getAnimePage");

    const [media, error] = await AnimeUseCases.getAnimesByIdUseCase().execute(
      animeId
    );
    if (error) {
      logger.error("Failed to fetch anime by id", error);
      res.status(500).send({
        type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
        data: {
          content: "internal server error",
        },
      });
      return;
    }

    res.status(200).send({
      type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
      data: {
        embeds: [
          {
            title: media!.title.romaji + " (" + media!.title.native + ")",
            description: TextService.postProcessing(media!.description),
            color: Colors.lightBlue,
            thumbnail: {
              url: media!.trailer?.thumbnail || media!.bannerImage,
            },
            image: {
              url: media!.bannerImage,
            },
            url: "https://www.youtube.com/watch?v=" + media!.trailer?.id || "",
            fields: [
              {
                name: "Genres",
                value: media!.genres.join(", "),
                inline: false,
              },
              {
                name: "Average Score",
                value: media!.averageScore + "%",
                inline: false,
              },
              {
                name: "Episodes",
                value: media!.episodes,
                inline: false,
              },
              {
                name: "Start Date",
                value:
                  media!.startDate!.day +
                  " " +
                  MonthService.getMonthName(media!.startDate!.month) +
                  " " +
                  media!.startDate!.year,
                inline: true,
              },
              {
                name: "End Date",
                value:
                  media!.endDate!.day +
                  " " +
                  MonthService.getMonthName(media!.endDate!.month) +
                  " " +
                  media!.endDate!.year,
                inline: true,
              },
              {
                name: "Status",
                value: media!.status,
                inline: false,
              },
            ],
            timestamp: new Date(),
          },
        ],
      },
    });
  }
  async pagination(
    currentPage: string,
    hasNextPage: string,
    animeName: string,
    ctx: string,
    res: Response
  ): Promise<void> {
    logger.trace("Call pagination");

    const [medias, pageInfo, error] =
      await AnimeUseCases.handlePaginationUseCase().execute(ctx, animeName, {
        currentPage: parseInt(currentPage),
        hasNextPage: hasNextPage === "true\n",
      });

    if (error) {
      logger.error("Failed to fetch paginated animes", error);
      res.status(200).send({
        type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
        data: {
          content:
            "😎 Next or Previoous page not avaiable, sorry but I'm too lazy",
        },
      });
      return;
    }

    if (medias === null) {
      logger.error("Please check the page input", new Error("No animes found"));
      res.status(400).send({
        type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
        data: {
          content: "No animes found",
        },
      });
      return;
    }

    const [fields, fieldsBtn] = AnimeUseCases.getEmbedFieldsUseCase().execute(
      medias!
    );
    res.status(200).send({
      type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
      data: {
        embeds: [
          {
            description: "List of animes that match your search",
            color: Colors.white,
            fields: fields,
            timestamp: new Date(),
            image: {
              url: AnimeUseCases.getBestRatingAnimeUseCase().execute(medias!)
                .bannerImage,
            },
          },
        ],
        components: [
          {
            type: MessageComponentTypes.ACTION_ROW,
            components: fieldsBtn,
          },
          {
            type: MessageComponentTypes.ACTION_ROW,
            components: [
              {
                type: MessageComponentTypes.BUTTON,
                label: "Prev",
                style: 1,
                custom_id: `${FeatureTags.ANIME_SEARCH} prev ${
                  pageInfo!.currentPage
                } ${pageInfo!.hasNextPage}
                ${animeName}`,
              },
              {
                type: MessageComponentTypes.BUTTON,
                label: "Next",
                style: 3,
                custom_id: `${FeatureTags.ANIME_SEARCH} next ${
                  pageInfo!.currentPage
                } ${pageInfo!.hasNextPage}
                ${animeName}`,
              },
            ],
          },
        ],
      },
    });
  }

  async getFirstPage(animeName: string, res: Response): Promise<void> {
    logger.trace("Call getFirstPage");

    const [medias, pageInfo, error] =
      await AnimeUseCases.getAnimesByNameUseCase().execute(animeName);

    if (error) {
      logger.error("Failed to fetch first page animes", error);
      res.status(500).send({
        type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
        data: {
          content: "internal server error",
        },
      });
      return;
    }

    if (medias === null) {
      logger.error("Please check the name input", new Error("No animes found"));
      res.status(400).send({
        type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
        data: {
          content: "No animes found",
        },
      });
      return;
    }

    const [fields, fieldsBtn] = AnimeUseCases.getEmbedFieldsUseCase().execute(
      medias!
    );

    res.status(200).send({
      type: InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
      data: {
        embeds: [
          {
            description: "List of animes that match your search",
            color: Colors.white,
            fields: fields,
            timestamp: new Date(),
            image: {
              url: AnimeUseCases.getBestRatingAnimeUseCase().execute(medias!)
                .bannerImage,
            },
          },
        ],
        components: [
          {
            type: MessageComponentTypes.ACTION_ROW,
            components: fieldsBtn,
          },
          {
            type: MessageComponentTypes.ACTION_ROW,
            components: [
              {
                type: MessageComponentTypes.BUTTON,
                label: "Prev",
                style: 1,
                custom_id: `${FeatureTags.ANIME_SEARCH} prev ${
                  pageInfo!.currentPage
                } ${pageInfo!.hasNextPage}
          ${animeName}`,
              },
              {
                type: MessageComponentTypes.BUTTON,
                label: "Next",
                style: 3,
                custom_id: `${FeatureTags.ANIME_SEARCH} next ${
                  pageInfo!.currentPage
                } ${pageInfo!.hasNextPage}
          ${animeName}`,
              },
            ],
          },
        ],
      },
    });
  }
}
