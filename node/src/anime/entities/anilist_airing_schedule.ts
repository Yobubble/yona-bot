export class AnilistAiringSchedule {
  constructor(
    public readonly airingSchedule: {
      nodes: [
        {
          timeUntilAiring: number;
        }
      ];
    }
  ) {}
}
