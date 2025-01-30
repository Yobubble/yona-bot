export class AnilistQueries {
  public static readonly fetchAnimesByName = `query ($perPage: Int, $search: String) {
                                          Page (page: 1, perPage: $perPage) {
                                            pageInfo {
                                              currentPage
                                              hasNextPage
                                            }
                                            media (search: $search, type: ANIME) {
                                              id
                                              title {
                                                romaji
                                                english
                                                native
                                              }
                                              genres
                                              bannerImage
                                              averageScore
                                            }
                                          }
                                        }`;
  public static readonly fetchAnimesByNameAndPage = `query ($page: Int, $perPage: Int, $search: String) {
                                          Page (page: $page, perPage: $perPage) {
                                            pageInfo {
                                              currentPage
                                              hasNextPage
                                            }
                                            media (search: $search, type: ANIME) {
                                              id
                                              title {
                                                romaji
                                                english
                                                native
                                              }
                                              genres
                                              bannerImage
                                              averageScore
                                            }
                                          }
                                        }`;
  public static readonly fetchAnimeById = `query($mediaId: Int)  {
                                            Media(id: $mediaId) {
                                              id
                                              title {
                                                english
                                                native
                                                romaji
                                              }
                                              description
                                              episodes
                                              genres
                                              bannerImage
                                              averageScore
                                              startDate {
                                                day
                                                month
                                                year
                                              }
                                              endDate {
                                                day
                                                month
                                                year
                                              }
                                              trailer {
                                                id
                                                site
                                                thumbnail
                                              }
                                              status
                                            }
                                          }`;
}
