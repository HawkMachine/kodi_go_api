package kodi

// Implementation of VideoLibrary RPCs.
type VideoLibrary struct {
	k *Kodi
}

func (vl *VideoLibrary) doRPC(rpcName string, params interface{}, response interface{}) error {
	req := &request{
		Jsonrpc: "2.0",
		Method:  "VideoLibrary." + rpcName,
		Id:      1,
		Params:  params,
	}
	return vl.k.doRPC(req, response)
}

// VideoLibrary.GetMovies

type VideoLibraryGetMoviesFilter struct {
	// one of:
	*ListFilterMovies
	GenreId  *LibraryId `json:"genreid,omitempty"`
	Genre    *string    `json:"genre,omitempty"`
	Year     *int       `json:"year,omitempty"`
	Actor    *string    `json:"actor,omitempty"`
	Director *string    `json:"director,omitempty"`
	Studio   *string    `json:"studio,omitempty"`
	Country  *string    `json:"country,omitempty"`
	SetId    *LibraryId `json:"setid,omitempty"`
	Set      *string    `json:"set,omitempty"`
	Tag      *string    `json:"tag,omitempty"`
}

type VideoLibraryGetMoviesParams struct {
	Properties []VideoFieldsMovie           `json:"properties"`
	Limits     *ListLimits                  `json:"limits,omitempty"`
	Sort       *ListSort                    `json:"sort,omitempty"`
	Filter     *VideoLibraryGetMoviesFilter `json:"filter,omitempty"`
}

type VideoLibraryGetMoviesResult struct {
	Limits ListLimitsReturned   `json:"limits,omitempty"`
	Movies []*VideoDetailsMovie `json:"movies,omitempty"`
}

type VideoLibraryGetMoviesResponse struct {
	ResponseBase
	Result *VideoLibraryGetMoviesResult `json:"result,omitempty"`
}

func (vl *VideoLibrary) GetMovies(params *VideoLibraryGetMoviesParams) (*VideoLibraryGetMoviesResponse, error) {
	response := &VideoLibraryGetMoviesResponse{}
	err := vl.doRPC("GetMovies", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// VideoLibrary.GetMovieDetails

type VideoLibraryGetMovieDetailsParams struct {
	MovieId    LibraryId          `json:"movieid"`
	Properties []VideoFieldsMovie `json:"properties,omitempty"`
}

type VideoLibraryGetMovieDetailsResult struct {
	MovieDetails *VideoDetailsMovie `json:"moviedetails,omitempty"`
}

type VideoLibraryGetMovieDetailsResponse struct {
	ResponseBase
	Result *VideoLibraryGetMovieDetailsResult
}

func (vl *VideoLibrary) GetMovieDetails(params *VideoLibraryGetMovieDetailsParams) (*VideoLibraryGetMovieDetailsResponse, error) {
	response := &VideoLibraryGetMovieDetailsResponse{}
	err := vl.doRPC("GetMovieDetailsResponse", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// VideoLibrary.GetEpisodes

type VideoLibraryGetEpisodesFilter struct {
	// one of:
	*ListFilterEpisodes
	GenreId  *LibraryId `json:"genreid,omitempty"`
	Genre    *string    `json:"genre,omitempty"`
	Year     *int       `json:"year,omitempty"`
	Actor    *string    `json:"actor,omitempty"`
	Director *string    `json:"director,omitempty"`
}

func VideoLibraryCreateEpisodesFilterByField(field ListFilterFieldsEpisodes, operator ListFilterOperator, value string) *VideoLibraryGetEpisodesFilter {
	return &VideoLibraryGetEpisodesFilter{
		ListFilterEpisodes: &ListFilterEpisodes{
			ListFilterRuleEpisodes: &ListFilterRuleEpisodes{
				Field:    field,
				Operator: operator,
				Value:    value,
			},
		},
	}
}

type VideoLibraryGetEpisodesParams struct {
	TVShowId   *LibraryId                     `json:"tvshowid,omitempty"`
	Season     *int                           `json:"season,omitempty"`
	Properties []VideoFieldsEpisode           `json:"properties,omitempty"`
	Limits     *ListLimits                    `json:"limits,omitempty"`
	Sort       *ListSort                      `json:"sort,omitempty"`
	Filter     *VideoLibraryGetEpisodesFilter `json:"filter,omitempty"`
}

type VideoLibraryGetEpisodesResult struct {
	Episodes []*VideoDetailsEpisode `json:"episodes,omitempty"`
	Limits   ListLimitsReturned     `json:"limits"`
}

type VideoLibraryGetEpisodesResponse struct {
	ResponseBase
	Result *VideoLibraryGetEpisodesResult `json:"result,omitempty"`
}

func (vl *VideoLibrary) GetEpisodes(params *VideoLibraryGetEpisodesParams) (*VideoLibraryGetEpisodesResponse, error) {
	response := &VideoLibraryGetEpisodesResponse{}
	err := vl.doRPC("GetEpisodes", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// VideoLibrary.GetTVShows

type VideoLibraryGetTVShowsFilter struct {
	// one of
	*ListFilterTVShows
	GenreId *LibraryId `json:"genreid,omitempty"`
	Genre   *string    `json:"genre,omitempty"`
	Year    *int       `json:"year,omitempty"`
	Actor   *string    `json:"actor,omitempty"`
	Studio  *string    `json:"studio,omitempty"`
	Tag     *string    `json:"tag,omitempty"`
}

type VideoLibraryGetTVShowsParams struct {
	Properties []VideoFieldsTVShow `json:"properties,omitempty"`
	Limits     *ListLimits         `json:"limits,omitempty"`
	Sort       *ListSort           `json:"sort,omitempty"`
	Filter     *Filter             `json:"filter,omitempty"`
}

type VideoLibraryGetTVShowsResult struct {
	Limits  ListLimitsReturned    `json:"limits,omitempty"`
	TVShows []*VideoDetailsTVShow `json:"tvshows,omitempty"`
}

type VideoLibraryGetTVShowsResponse struct {
	ResponseBase
	Result *VideoLibraryGetTVShowsResult `json:"result,omitempty"`
}

func (vl *VideoLibrary) GetTVShows(params *VideoLibraryGetTVShowsParams) (*VideoLibraryGetTVShowsResponse, error) {
	response := &VideoLibraryGetTVShowsResponse{}
	err := vl.doRPC("GetTVShows", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// VideoLibrary.GetTVShowDetails

type VideoLibraryGetTVShowDetailsParams struct {
	TVShowId   LibraryId           `json:"tvshowid"`
	Properties []VideoFieldsTVShow `json:"properties,omitempty"`
}

type VideoLibraryGetTVShowDetailsResult struct {
	TVShowDetails *VideoDetailsTVShow `json:"tvshowdetails,omitempty"`
}

type VideoLibraryGetTVShowDetailsResponse struct {
	ResponseBase
	Result *VideoLibraryGetTVShowDetailsResult `json:"result,omitempty"`
}

func (vl *VideoLibrary) GetTVShowDetails(params *VideoLibraryGetTVShowDetailsParams) (*VideoLibraryGetTVShowDetailsResponse, error) {
	response := &VideoLibraryGetTVShowDetailsResponse{}
	err := vl.doRPC("GetTVShowDetails", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// VideoLibrary.GetSeasons

type VideoLibraryGetSeasonsParams struct {
	TVShowId   LibraryId           `json:"tvshowid"`
	Properties []VideoFieldsSeason `json:"properties,omitempty"`
	Limits     *ListLimits         `json:"limits,omitempty"`
	Sort       *ListSort           `json:"sort,omitempty"`
}

type VideoLibraryGetSeasonsResult struct {
	Limits  ListLimitsReturned    `json:"limits,omitempty"`
	Seasons []*VideoDetailsSeason `json:"seasons,omitempty"`
}

type VideoLibraryGetSeasonsResponse struct {
	ResponseBase
	Result *VideoLibraryGetSeasonsResult `json:"result,omitempty"`
}

func (vl *VideoLibrary) GetSeasons(params *VideoLibraryGetSeasonsParams) (*VideoLibraryGetSeasonsResponse, error) {
	response := &VideoLibraryGetSeasonsResponse{}
	err := vl.doRPC("GetSeasons", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// 5.12 VideoLibrary.Scan

type VideoLibraryScanParams struct {
	Directory   string `json:"directory,omitempty"`
	ShowDialogs bool   `json:"showdialogs,omitempty"`
}

type VideoLibraryScanResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}

func (vl *VideoLibrary) Scan(params *VideoLibraryScanParams) (*VideoLibraryScanResponse, error) {
	response := &VideoLibraryScanResponse{}
	err := vl.doRPC("Scan", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
