export class AnilistSmallMedia {
  constructor(
    public readonly id: number,
    public title: {
      romaji: string;
      english: string;
      native: string;
    },
    public genres: string[],
    public bannerImage: string,
    public averageScore: number
  ) {}
}
