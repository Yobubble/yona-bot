export class AnilistBigMedia {
  constructor(
    public readonly id: number,
    public title: {
      romaji: string;
      english: string;
      native: string;
    },
    public description: string,
    public episodes: number,
    public genres: string[],
    public bannerImage: string,
    public averageScore: number,
    public startDate: {
      day: number;
      month: number;
      year: number;
    },
    public endDate: {
      day: number;
      month: number;
      year: number;
    },
    public trailer: {
      id: string;
      site: string;
      thumbnail: string;
    },
    public status: string
  ) {}
}
