import { Response } from "express";
import { logger } from "../../../utils/services/logger_service";

export class AnimeAiringSchedule {
  async getAiringSchedulePage(res: Response): Promise<void> {
    logger.trace("Call getAiringSchedulePage");
    // TODO: Implement getAiringSchedulePage
  }
}
