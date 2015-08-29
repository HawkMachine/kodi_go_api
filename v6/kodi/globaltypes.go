package kodi

type FileType string

const (
	FILE_TYPE_FILE FileType = "file"
	FILE_TYPE_DIR  FileType = "directory"
)

type LibraryId int

type Filter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type ItemDetailsBase struct {
	Label string `json:"label"`
}

func FilterString(a string) *string {
	return &a
}

// 6.4 List http://kodi.wiki/view/JSON-RPC_API/v6#Audio

type AudioDetailsBase struct {
	MediaDetailsBase
	Genre []string `json:"genre"`
}

type AudioDetailsMedia struct {
}

// 6.12 List http://kodi.wiki/view/JSON-RPC_API/v6#List

type ListFieldsAll string

const (
	LIST_FILTER_FIELD_TITLE                        ListFieldsAll = "title"
	LIST_FILTER_FIELD_ARTIST                       ListFieldsAll = "artist"
	LIST_FILTER_FIELD_ALBUMARTIST                  ListFieldsAll = "albumartist"
	LIST_FILTER_FIELD_GENRE                        ListFieldsAll = "genre"
	LIST_FILTER_FIELD_YEAR                         ListFieldsAll = "year"
	LIST_FILTER_FIELD_RATING                       ListFieldsAll = "rating"
	LIST_FILTER_FIELD_ALBUM                        ListFieldsAll = "album"
	LIST_FILTER_FIELD_TRACK                        ListFieldsAll = "track"
	LIST_FILTER_FIELD_DURATION                     ListFieldsAll = "duration"
	LIST_FILTER_FIELD_COMMENT                      ListFieldsAll = "comment"
	LIST_FILTER_FIELD_LYRICS                       ListFieldsAll = "lyrics"
	LIST_FILTER_FIELD_MUSIC_BRAINZ_TRACK_ID        ListFieldsAll = "musicbrainztrackid"
	LIST_FILTER_FIELD_MUSIC_BRAINZ_ARTIST_ID       ListFieldsAll = "musicbrainzartistid"
	LIST_FILTER_FIELD_MUSIC_BRAINZ_ALBUM_ID        ListFieldsAll = "musicbrainzalbumid"
	LIST_FILTER_FIELD_MUSIC_BRAINZ_ALBUM_ARTIST_ID ListFieldsAll = "musicbrainzalbumartistid"
	LIST_FILTER_FIELD_PLAYCOUNT                    ListFieldsAll = "playcount"
	LIST_FILTER_FIELD_FANART                       ListFieldsAll = "fanart"
	LIST_FILTER_FIELD_DIRECTOR                     ListFieldsAll = "director"
	LIST_FILTER_FIELD_TRAILER                      ListFieldsAll = "trailer"
	LIST_FILTER_FIELD_TAGLINE                      ListFieldsAll = "tagline"
	LIST_FILTER_FIELD_PLOT                         ListFieldsAll = "plot"
	LIST_FILTER_FIELD_PLOT_OUTLINE                 ListFieldsAll = "plotoutline"
	LIST_FILTER_FIELD_ORIGINAL_TITLE               ListFieldsAll = "originaltitle"
	LIST_FILTER_FIELD_LAST_PLAYED                  ListFieldsAll = "lastplayed"
	LIST_FILTER_FIELD_WRITER                       ListFieldsAll = "writer"
	LIST_FILTER_FIELD_STUDIO                       ListFieldsAll = "studio"
	LIST_FILTER_FIELD_MPAA                         ListFieldsAll = "mpaa"
	LIST_FILTER_FIELD_CAST                         ListFieldsAll = "cast"
	LIST_FILTER_FIELD_COUNTRY                      ListFieldsAll = "country"
	LIST_FILTER_FIELD_IMDB_NUMBER                  ListFieldsAll = "imdbnumber"
	LIST_FILTER_FIELD_PREMIERED                    ListFieldsAll = "premiered"
	LIST_FILTER_FIELD_PRODUCTION_CODE              ListFieldsAll = "productioncode"
	LIST_FILTER_FIELD_RUNTIME                      ListFieldsAll = "runtime"
	LIST_FILTER_FIELD_SET                          ListFieldsAll = "set"
	LIST_FILTER_FIELD_SHOW_LINK                    ListFieldsAll = "showlink"
	LIST_FILTER_FIELD_STREAM_DETAILS               ListFieldsAll = "streamdetails"
	LIST_FILTER_FIELD_TOP250                       ListFieldsAll = "top250"
	LIST_FILTER_FIELD_VOTES                        ListFieldsAll = "votes"
	LIST_FILTER_FIELD_FIRST_AIRED                  ListFieldsAll = "firstaired"
	LIST_FILTER_FIELD_SEASON                       ListFieldsAll = "season"
	LIST_FILTER_FIELD_EPISODE                      ListFieldsAll = "episode"
	LIST_FILTER_FIELD_SHOW_TITLE                   ListFieldsAll = "showtitle"
	LIST_FILTER_FIELD_THUMBNAIL                    ListFieldsAll = "thumbnail"
	LIST_FILTER_FIELD_FILE                         ListFieldsAll = "file"
	LIST_FILTER_FIELD_RESUME                       ListFieldsAll = "resume"
	LIST_FILTER_FIELD_ARTIST_ID                    ListFieldsAll = "artistid"
	LIST_FILTER_FIELD_ALBUM_ID                     ListFieldsAll = "albumid"
	LIST_FILTER_FIELD_TV_SHOW_ID                   ListFieldsAll = "tvshowid"
	LIST_FILTER_FIELD_SETID                        ListFieldsAll = "setid"
	LIST_FILTER_FIELD_WATCHED_EPISODES             ListFieldsAll = "watchedepisodes"
	LIST_FILTER_FIELD_DISC                         ListFieldsAll = "disc"
	LIST_FILTER_FIELD_TAG                          ListFieldsAll = "tag"
	LIST_FILTER_FIELD_ART                          ListFieldsAll = "art"
	LIST_FILTER_FIELD_GENRE_ID                     ListFieldsAll = "genreid"
	LIST_FILTER_FIELD_DISPLAY_ARTIST               ListFieldsAll = "displayartist"
	LIST_FILTER_FIELD_ALBUM_ARTIST_ID              ListFieldsAll = "albumartistid"
	LIST_FILTER_FIELD_DESCRIPTION                  ListFieldsAll = "description"
	LIST_FILTER_FIELD_THEME                        ListFieldsAll = "theme"
	LIST_FILTER_FIELD_MOOD                         ListFieldsAll = "mood"
	LIST_FILTER_FIELD_STYLE                        ListFieldsAll = "style"
	LIST_FILTER_FIELD_ALBUM_LABEL                  ListFieldsAll = "albumlabel"
	LIST_FILTER_FIELD_SORT_TITLE                   ListFieldsAll = "sorttitle"
	LIST_FILTER_FIELD_EPISODE_GUIDE                ListFieldsAll = "episodeguide"
	LIST_FILTER_FIELD_UNIQUE_ID                    ListFieldsAll = "uniqueid"
	LIST_FILTER_FIELD_DATE_ADDED                   ListFieldsAll = "dateadded"
	LIST_FILTER_FIELD_CHANNEL                      ListFieldsAll = "channel"
	LIST_FILTER_FIELD_CHANNEL_TYPE                 ListFieldsAll = "channeltype"
	LIST_FILTER_FIELD_HIDDEN                       ListFieldsAll = "hidden"
	LIST_FILTER_FIELD_LOCKED                       ListFieldsAll = "locked"
	LIST_FILTER_FIELD_CHANNEL_NUMBER               ListFieldsAll = "channelnumber"
	LIST_FILTER_FIELD_START_TIME                   ListFieldsAll = "starttime"
	LIST_FILTER_FIELD_END_TIME                     ListFieldsAll = "endtime"
)

type ListSortmethod string

const (
	SORT_METHOD_NONE             ListSortmethod = "none"
	SORT_METHOD_LABEL            ListSortmethod = "label"
	SORT_METHOD_DATE             ListSortmethod = "date"
	SORT_METHOD_SIZE             ListSortmethod = "size"
	SORT_METHOD_FILE             ListSortmethod = "file"
	SORT_METHOD_PATH             ListSortmethod = "path"
	SORT_METHOD_DRIVE_TYPE       ListSortmethod = "drivetype"
	SORT_METHOD_TITLE            ListSortmethod = "title"
	SORT_METHOD_TRACK            ListSortmethod = "track"
	SORT_METHOD_TIME             ListSortmethod = "time"
	SORT_METHOD_ARTIST           ListSortmethod = "artist"
	SORT_METHOD_ALBUM            ListSortmethod = "album"
	SORT_METHOD_ALBUM_TYPE       ListSortmethod = "albumtype"
	SORT_METHOD_GENRE            ListSortmethod = "genre"
	SORT_METHOD_COUNTRY          ListSortmethod = "country"
	SORT_METHOD_YEAR             ListSortmethod = "year"
	SORT_METHOD_RATING           ListSortmethod = "rating"
	SORT_METHOD_VOTES            ListSortmethod = "votes"
	SORT_METHOD_TOP250           ListSortmethod = "top250"
	SORT_METHOD_PROGRAM_COUNT    ListSortmethod = "programcount"
	SORT_METHOD_PLAYLIST         ListSortmethod = "playlist"
	SORT_METHOD_EPISODE          ListSortmethod = "episode"
	SORT_METHOD_SEASON           ListSortmethod = "season"
	SORT_METHOD_TOTAL_EPISODES   ListSortmethod = "totalepisode"
	SORT_METHOD_WATCHED_EPISODES ListSortmethod = "watchedepisodes"
	SORT_METHOD_TV_SHOW_STATUS   ListSortmethod = "tvshowstatus"
	SORT_METHOD_TV_SHOW_TITLE    ListSortmethod = "tvshowtitle"
	SORT_METHOD_SORT_TITLE       ListSortmethod = "sorttitle"
	SORT_METHOD_PRODUCTION_CODE  ListSortmethod = "productioncode"
	SORT_METHOD_MPAA             ListSortmethod = "mpaa"
	SORT_METHOD_STUDIO           ListSortmethod = "studio"
	SORT_METHOD_DATE_ADDED       ListSortmethod = "dataadded"
	SORT_METHOD_LAST_PLAYED      ListSortmethod = "lastplayed"
	SORT_METHOD_PLAY_COUNT       ListSortmethod = "playcount"
	SORT_METHOD_LISTENERS        ListSortmethod = "listeners"
	SORT_METHOD_BIT_RATE         ListSortmethod = "bitrate"
	SORT_METHOD_RANDOM           ListSortmethod = "random"
)

type ListFilterFieldsMovies string

const (
	MOVIE_FILTER_FIELD_TITLE             ListFilterFieldsMovies = "title"
	MOVIE_FILTER_FIELD_PLOT              ListFilterFieldsMovies = "plot"
	MOVIE_FILTER_FIELD_PLOT_OUTLINE      ListFilterFieldsMovies = "plotoutline"
	MOVIE_FILTER_FIELD_TAGLINE           ListFilterFieldsMovies = "tagline"
	MOVIE_FILTER_FIELD_VOTES             ListFilterFieldsMovies = "votes"
	MOVIE_FILTER_FIELD_RATING            ListFilterFieldsMovies = "rating"
	MOVIE_FILTER_FIELD_TIME              ListFilterFieldsMovies = "time"
	MOVIE_FILTER_FIELD_WRITERS           ListFilterFieldsMovies = "writers"
	MOVIE_FILTER_FIELD_PLAY_COUNT        ListFilterFieldsMovies = "playcount"
	MOVIE_FILTER_FIELD_LAST_PLAYED       ListFilterFieldsMovies = "lastplayed"
	MOVIE_FILTER_FIELD_IN_PROGRESS       ListFilterFieldsMovies = "inprogress"
	MOVIE_FILTER_FIELD_GENRE             ListFilterFieldsMovies = "genre"
	MOVIE_FILTER_FIELD_COUNTRY           ListFilterFieldsMovies = "country"
	MOVIE_FILTER_FIELD_YEAR              ListFilterFieldsMovies = "year"
	MOVIE_FILTER_FIELD_DIRECTOR          ListFilterFieldsMovies = "director"
	MOVIE_FILTER_FIELD_ACTOR             ListFilterFieldsMovies = "actor"
	MOVIE_FILTER_FIELD_MPAA_RATING       ListFilterFieldsMovies = "mpaarating"
	MOVIE_FILTER_FIELD_TOP250            ListFilterFieldsMovies = "top250"
	MOVIE_FILTER_FIELD_STUDIO            ListFilterFieldsMovies = "studio"
	MOVIE_FILTER_FIELD_HAS_TRAILER       ListFilterFieldsMovies = "hastrailer"
	MOVIE_FILTER_FIELD_FILENAME          ListFilterFieldsMovies = "filename"
	MOVIE_FILTER_FIELD_PATH              ListFilterFieldsMovies = "path"
	MOVIE_FILTER_FIELD_SET               ListFilterFieldsMovies = "set"
	MOVIE_FILTER_FIELD_TAG               ListFilterFieldsMovies = "tag"
	MOVIE_FILTER_FIELD_DATE_ADDED        ListFilterFieldsMovies = "dateadded"
	MOVIE_FILTER_FIELD_VIDEO_RESOLUTION  ListFilterFieldsMovies = "videoresolution"
	MOVIE_FILTER_FIELD_AUDIO_CHANNELS    ListFilterFieldsMovies = "audiochannels"
	MOVIE_FILTER_FIELD_VIDEO_CODEC       ListFilterFieldsMovies = "videocodec"
	MOVIE_FILTER_FIELD_AUDIO_CODEC       ListFilterFieldsMovies = "audiocodec"
	MOVIE_FILTER_FIELD_AUDIO_LANGUAGE    ListFilterFieldsMovies = "audiolanguage"
	MOVIE_FILTER_FIELD_SUBTITLE_LANGUAGE ListFilterFieldsMovies = "subtitlelanguage"
	MOVIE_FILTER_FIELD_VIDEO_ASPECT      ListFilterFieldsMovies = "videoaspect"
	MOVIE_FILTER_FIELD_PLAYLIST          ListFilterFieldsMovies = "playlist"
)

type ListFilterFieldsEpisodes string

const (
	EPISODE_FILTER_FIELD_TITLE             ListFilterFieldsEpisodes = "title"
	EPISODE_FILTER_FIELD_TV_SHOW           ListFilterFieldsEpisodes = "tvshow"
	EPISODE_FILTER_FIELD_PLOT              ListFilterFieldsEpisodes = "plot"
	EPISODE_FILTER_FIELD_VOTES             ListFilterFieldsEpisodes = "votes"
	EPISODE_FILTER_FIELD_RATING            ListFilterFieldsEpisodes = "rating"
	EPISODE_FILTER_FIELD_TIME              ListFilterFieldsEpisodes = "time"
	EPISODE_FILTER_FIELD_WRITERS           ListFilterFieldsEpisodes = "writers"
	EPISODE_FILTER_FIELD_AIR_DATE          ListFilterFieldsEpisodes = "airdate"
	EPISODE_FILTER_FIELD_PLAYCOUNT         ListFilterFieldsEpisodes = "playcount"
	EPISODE_FILTER_FIELD_LAST_PLAYED       ListFilterFieldsEpisodes = "lastplayed"
	EPISODE_FILTER_FIELD_IN_PROGRESS       ListFilterFieldsEpisodes = "inprogress"
	EPISODE_FILTER_FIELD_GENRE             ListFilterFieldsEpisodes = "genre"
	EPISODE_FILTER_FIELD_YEAR              ListFilterFieldsEpisodes = "year"
	EPISODE_FILTER_FIELD_DIRECTOR          ListFilterFieldsEpisodes = "director"
	EPISODE_FILTER_FIELD_ACTOR             ListFilterFieldsEpisodes = "actor"
	EPISODE_FILTER_FIELD_EPISODE           ListFilterFieldsEpisodes = "episode"
	EPISODE_FILTER_FIELD_SEASON            ListFilterFieldsEpisodes = "season"
	EPISODE_FILTER_FIELD_FILENAME          ListFilterFieldsEpisodes = "filename"
	EPISODE_FILTER_FIELD_PATH              ListFilterFieldsEpisodes = "path"
	EPISODE_FILTER_FIELD_STUDIO            ListFilterFieldsEpisodes = "studio"
	EPISODE_FILTER_FIELD_MPAA_RATING       ListFilterFieldsEpisodes = "mpaarating"
	EPISODE_FILTER_FIELD_DATE_ADDED        ListFilterFieldsEpisodes = "dateadded"
	EPISODE_FILTER_FIELD_VIDEO_RESOLUTION  ListFilterFieldsEpisodes = "videoresolution"
	EPISODE_FILTER_FIELD_AUDIO_CHANNELS    ListFilterFieldsEpisodes = "audiochannels"
	EPISODE_FILTER_FIELD_VIDEO_CODEC       ListFilterFieldsEpisodes = "videocodec"
	EPISODE_FILTER_FIELD_AUDIO_CODEC       ListFilterFieldsEpisodes = "audiocodec"
	EPISODE_FILTER_FIELD_AUDIO_LANGUAGE    ListFilterFieldsEpisodes = "audiolanguage"
	EPISODE_FILTER_FIELD_SUBTITLE_LANGUAGE ListFilterFieldsEpisodes = "subtitlelanguage"
	EPISODE_FILTER_FIELD_VIDEO_ASPECT      ListFilterFieldsEpisodes = "videoaspect"
	EPISODE_FILTER_FIELD_PLAYLIST          ListFilterFieldsEpisodes = "playlist"
)

type ListFilterFieldsTVShows string

const (
	TV_SHOW_FILTER_FIELD_TITLE        ListFilterFieldsTVShows = "title"
	TV_SHOW_FILTER_FIELD_PLOT         ListFilterFieldsTVShows = "plot"
	TV_SHOW_FILTER_FIELD_STATUS       ListFilterFieldsTVShows = "status"
	TV_SHOW_FILTER_FIELD_VOTES        ListFilterFieldsTVShows = "votes"
	TV_SHOW_FILTER_FIELD_RATING       ListFilterFieldsTVShows = "rating"
	TV_SHOW_FILTER_FIELD_YEAR         ListFilterFieldsTVShows = "year"
	TV_SHOW_FILTER_FIELD_GENRE        ListFilterFieldsTVShows = "genre"
	TV_SHOW_FILTER_FIELD_DIRECTOR     ListFilterFieldsTVShows = "director"
	TV_SHOW_FILTER_FIELD_ACTOR        ListFilterFieldsTVShows = "actor"
	TV_SHOW_FILTER_FIELD_NUM_EPISODES ListFilterFieldsTVShows = "numepisodes"
	TV_SHOW_FILTER_FIELD_NUM_WATCHED  ListFilterFieldsTVShows = "numwatched"
	TV_SHOW_FILTER_FIELD_PLAY_COUNT   ListFilterFieldsTVShows = "playcount"
	TV_SHOW_FILTER_FIELD_PATH         ListFilterFieldsTVShows = "path"
	TV_SHOW_FILTER_FIELD_STUDIO       ListFilterFieldsTVShows = "studio"
	TV_SHOW_FILTER_FIELD_MPAA_RATING  ListFilterFieldsTVShows = "mpaarating"
	TV_SHOW_FILTER_FIELD_DATE_ADDED   ListFilterFieldsTVShows = "dateadded"
	TV_SHOW_FILTER_FIELD_LAST_PLAYED  ListFilterFieldsTVShows = "lastplayed"
	TV_SHOW_FILTER_FIELD_IN_PROGRESS  ListFilterFieldsTVShows = "inprogress"
	TV_SHOW_FILTER_FIELD_PLAYLIST     ListFilterFieldsTVShows = "playlist"
)

type ListSortOrder string

const (
	SORT_ORDER_ASCENDING  ListSortOrder = "ascending"
	SORT_ORDER_DESCENDING ListSortOrder = "descending"
)

type ListFilterOperator string

const (
	OPERATOR_CONTAINS         ListFilterOperator = "contains"
	OPERATOR_DOES_NOT_CONTAIN ListFilterOperator = "doesnotcontain"
	OPERATOR_IS               ListFilterOperator = "is"
	OPERATOR_IS_NOT           ListFilterOperator = "isnot"
	OPERATOR_STARTS_WITH      ListFilterOperator = "startswith"
	OPERATOR_ENDS_WITH        ListFilterOperator = "endswith"
	OPERATOR_GREATER_THAN     ListFilterOperator = "greaterthan"
	OPERATOR_LESS_THAN        ListFilterOperator = "lessthan"
	OPERATOR_AFTER            ListFilterOperator = "after"
	OPERATOR_BEFORE           ListFilterOperator = "before"
	OPERATOR_IN_THE_LAST      ListFilterOperator = "inthelast"
	OPERATOR_NOT_INT_THE_LAST ListFilterOperator = "notinthelast"
	OPERATOR_TRUE             ListFilterOperator = "true"
	OPERATOR_FALSE            ListFilterOperator = "false"
	OPERATOR_BETWEEN          ListFilterOperator = "between"
)

type ListSort struct {
	Order         string `json:"order"`
	Ignorearticle bool   `json:"ignorearticle"`
	Method        string `json:"method"`
}

type ListFilterRuleTVShows struct {
	Field    ListFilterFieldsTVShows `json:"field,omitempty"`
	Operator ListFilterOperator      `json:"operator"`
	Value    string                  `json:"value"`
}

type ListFilterRuleMovies struct {
	Field    ListFilterFieldsMovies `json:"field,omitempty"`
	Operator ListFilterOperator     `json:"operator"`
	Value    string                 `json:"value"`
}

type ListFilterTVShows struct {
	// one of:
	*ListFilterRuleTVShows
	And []*ListFilterTVShows `json:"and,omitempty"`
	Or  []*ListFilterTVShows `json:"or,omitempty"`
}

type ListFilterMovies struct {
	// one of:
	*ListFilterRuleMovies
	And []*ListFilterMovies `json:"and,omitempty"`
	Or  []*ListFilterMovies `json:"or,omitempty"`
}

type ListFilterRuleEpisodes struct {
	Field    ListFilterFieldsEpisodes `json:"field,omitempty"`
	Operator ListFilterOperator       `json:"operator"`
	Value    string                   `json:"value"`
}

type ListFilterEpisodes struct {
	// one of:
	*ListFilterRuleEpisodes
	And []*ListFilterEpisodes `json:"and,omitempty"`
	Or  []*ListFilterEpisodes `json:"or,omitempty"`
}

// 6.12.28 http://kodi.wiki/view/JSON-RPC_API/v6#List.Item.Base
type ListItemBase struct {
	AudioDetailsMedia
	VideoDetailsFile
	SortTitle           string              `json:"sorttitle,omitempty"`
	ProductionCode      string              `json:"productioncode,omitempty"`
	Cast                []*VideoCastElement `json:"cast,omitempty"`
	Votes               string              `json:"votes,omitempty"`
	Duration            int                 `json:"duration,omitempty"`
	Trailer             string              `json:"trailer,omitempty"`
	AlbumId             LibraryId           `json:"albumid,omitempty"`
	MusicBrainzArtistId string              `json:"musicbrainzartistid,omitempty"`
	MPAA                string              `json:"mpaa,omitempty"`
	AlbumLabel          string              `json:"albumlabel,omitempty"`
	OriginalTitle       string              `json:"originaltitle,omitempty"`
	Writer              []string            `json:"writer,omitempty"`
	AlbumArtistId       []int               `json:"albumartistid,omitempty"`
	Type                string              `json:"type,omitempty"`
	Episode             int                 `json:"episode,omitempty"`
	FirstAired          string              `json:"firstaired,omitempty"`
	ShowTitle           string              `json:"showtitle,omitempty"`
	Country             []string            `json:"country,omitempty"`
	Mood                []string            `json:"mood,omitempty"`
	Set                 string              `json:"set,omitempty"`
	MusicBrainzTrackId  string              `json:"musicbrainztrackid,omitempty"`
	Tag                 []string            `json:"tag,omitempty"`
	Lyrics              string              `json:"lyrics,omitempty"`
	Top250              int                 `json:"top250,omitempty"`
	Comment             string              `json:"comment,omitempty"`
	Premiered           string              `json:"premiered,omitempty"`
	ShowLink            []string            `json:"showlink,omitempty"`
	Style               []string            `json:"style,omitempty"`
	Album               string              `json:"album,omitempty"`
	TVShowId            LibraryId           `json:"tvshowid,omitempty"`
	Season              int                 `json:"season,omitempty"`
	Theme               []string            `json:"theme,omitempty"`
	Description         string              `json:"description,omitempty"`
	SetId               LibraryId           `json:"setid,omitempty"`
	Track               int                 `json:"track,omitempty"`
	TagLine             string              `json:"tagline,omitempty"`
	PlotOutline         string              `json:"plotoutline,omitempty"`
	WatchedEpisodes     int                 `json:"watchedepisodes,omitempty"`
	Id                  LibraryId           `json:"id,omitempty"`
	Disc                int                 `json:"disc,omitempty"`
	AlbumArtist         []string            `json:"albumartist,omitempty"`
	Studio              []string            `json:"studio,omitempty"`
	UniqueId            interface{}         `json:"uniqueid,omitempty"`
	EpisodeGuide        string              `json:"episodeguide,omitempty"`
	IMDBNumber          string              `json:"imdbnumber,omitempty,omitempty"`
}

// 6.12.27 http://kodi.wiki/view/JSON-RPC_API/v6#List.Item.All
type ListItemAll struct {
	ListItemBase
	Channeltype   PVRChannelType `json:"channeltype,omitempty"`
	Channel       string         `json:"channel,omitempty"`
	StartTime     string         `json:"starttime,omitempty"`
	EndTime       string         `json:"endtime,omitempty"`
	ChannelNumber int            `json:"channelnumber,omitempty"`
	Hidden        bool           `json:"hidden,omitempty"`
	Locked        bool           `json:"locked,omitempty"`
}

type ListLimits struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

type ListLimitsReturned struct {
	Total int `json:"total"`
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// 6.13 Media http://kodi.wiki/view/JSON-RPC_API/v6#Media

type MediaArtwork struct {
	Banner string `json:"banner"`
	Poster string `json:"poster"`
	Fanart string `json:"fanart"`
	Thumb  string `json:"thumb"`
}

type MediaDetailsBase struct {
	ItemDetailsBase
	Fanart    string `json:"fanart,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

// 6.16 http://kodi.wiki/view/JSON-RPC_API/v6#PVR

// 6.16.1 http://kodi.wiki/view/JSON-RPC_API/v6#PVR.Channel.Type
type PVRChannelType string

const (
	PVR_CHANNEL_TYPE_TV    PVRChannelType = "tv"
	PVR_CHANNEL_TYPE_RADIO PVRChannelType = "radio"
)

// 6.17 http://kodi.wiki/view/JSON-RPC_API/v6#Player

// 6.17.14 http://kodi.wiki/view/JSON-RPC_API/v6#Player.Type

type PlayerType string

const (
	PLAYER_TYPE_VIDEO  PlayerType = "video"
	PLAYER_TYPE_AUDIO  PlayerType = "audio"
	PLAYER_TYPE_PLAYER PlayerType = "picture"
)

type PlayerId int

// This is just my guess, but id has to be 0, 1 or 2.
const (
	PLAYER_AUDIO_ID   PlayerId = 0
	PLAYER_VIDEO_ID   PlayerId = 1
	PLAYER_PICTURE_ID PlayerId = 2
)

// 6.18 http://kodi.wiki/view/JSON-RPC_API/v6#Playlist

// 6.18.1 http://kodi.wiki/view/JSON-RPC_API/v6#Playlist.Type

type PlaylistId int

// This is just my guess, but id has to be 0, 1 or 2.
const (
	PLAYLIST_AUDIO_ID   PlaylistId = 0
	PLAYLIST_VIDEO_ID   PlaylistId = 1
	PLAYLIST_PICTURE_ID PlaylistId = 2
)

// 6.18.2 http://kodi.wiki/view/JSON-RPC_API/v6#Playlist.Item

type PlaylistItem struct {
	// one of
	File         *string    `json:"file"`
	Director     *string    `json:"director"`
	MovieId      *LibraryId `json:"movieid"`
	EpisodeId    *LibraryId `json:"episodeid"`
	MusicVideoId *LibraryId `json:"musicvideoid"`
	ArtistId     *LibraryId `json:"artistid"`
	AlbumId      *LibraryId `json:"albumid"`
	SongId       *LibraryId `json:"songid"`
	GenreId      *LibraryId `json:"genreid"`
}

// 6.18.4 http://kodi.wiki/view/JSON-RPC_API/v6#Playlist.Property.Name

type PlaylistPropertyName string

const (
	PLAYLIST_PROPERTY_TYPE PlaylistPropertyName = "type"
	PLAYLIST_PROPERTY_SIZE PlaylistPropertyName = "size"
)

// 6.18.5 http://kodi.wiki/view/JSON-RPC_API/v6#Playlist.Property.Value

type PlaylistPropertyValue struct {
	Size int          `json:"size"`
	Type PlaylistType `json:"type"`
}

// 6.18.6 http://kodi.wiki/view/JSON-RPC_API/v6#Playlist.Type

type PlaylistType string

const (
	PLAYLIST_TYPE_UNKNOWN PlaylistType = "unknown"
	PLAYLIST_TYPE_VIDEO   PlaylistType = "video"
	PLAYLIST_TYPE_AUDIO   PlaylistType = "audio"
	PLAYLIST_TYPE_PICTURE PlaylistType = "picture"
	PLAYLIST_TYPE_MIXED   PlaylistType = "mixed"
)

// 6.20 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details

// 6.20.1 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Cast

type VideoCastElement struct {
	Thumbnail string `json:"thumbnail"`
	Name      string `json:"name"`
	Role      string `json:"role"`
}

// 6.20.2 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.Base

type VideoDetailsBase struct {
	MediaDetailsBase
	PlayCount int `json:"playcount"`
}

// 6.20.3 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.Episode

type VideoDetailsEpisode struct {
	VideoDetailsFile
	EpisodeId      LibraryId           `json:"episodeid"` // non-optional
	Cast           []*VideoCastElement `json:"cast,omitempty"`
	ProductionCode string              `json:"productioncode,omitempty"`
	Rating         float32             `json:"rating,omitempty"`
	Votes          string              `json:"votes,omitempty"`
	Episode        int                 `json:"episode,omitempty"`
	ShowTitle      string              `json:"showtitle,omitempty"`
	TVShowId       LibraryId           `json:"tvshowid,omitempty"`
	Season         int                 `json:"season,omitempty"`
	FirstAired     string              `json:"firstaired,omitempty"`
	UniqueId       interface{}         `json:"uniqueid,omitempty"`
	Writer         []string            `json:"writer,omitempty"`
	OriginalTitle  string              `json:"originaltitle,omitempty"`
}

// 6.20.4 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.File

type VideoDetailsFile struct {
	VideoDetailsItem
	StreamDetails interface{} `json:"streamdetails,omitempty"`
	Director      []string    `json:"director,omitempty"`
	Resume        interface{} `json:"resume,omitempty"`
	Runtime       int         `json:"runtime,omitempty"`
}

// 6.20.5 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.Item

type VideoDetailsItem struct {
	VideoDetailsMedia
	DateAdded  string `json:"dateadded,omitempty"`
	File       string `json:"file,omitempty"`
	LastPlayed string `json:"lastplayed,omitempty"`
	Plot       string `json:"plot,omitempty"`
}

// 6.20.6 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.Media

type VideoDetailsMedia struct {
	VideoDetailsBase
	Title string `json:"title"`
}

// 6.20.7 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.Movie

type VideoDetailsMovie struct {
	VideoDetailsFile
	MovieId       LibraryId           `json:"movieid"` // non-optional
	PlotOutline   string              `json:"plotoutline,omitempty"`
	SortTitle     string              `json:"sorttitle,omitempty"`
	Cast          []*VideoCastElement `json:"cast,omitempty"`
	Votes         string              `json:"votes,omitempty"`
	ShowLink      []string            `json:"showlink,omitempty"`
	Top250        int                 `json:"top250,omitempty"`
	Trailer       string              `json:"trailer,omitempty"`
	Year          int                 `json:"year,omitempty"`
	Country       []string            `json:"country,omitempty"`
	Studio        []string            `json:"studio,omitempty"`
	Set           string              `json:"set,omitempty"`
	Genre         []string            `json:"genre,omitempty"`
	MPAA          string              `json:"mpaa,omitempty"`
	SetId         LibraryId           `json:"setid,omitempty"`
	Rating        float32             `json:"rating,omitempty"`
	Tag           string              `json:"tag,omitempty"`
	TagLine       string              `json:"tagline,omitempty"`
	Writer        []string            `json:"writer,omitempty"`
	OriginalTitle string              `json:"originaltitle,omitempty"`
	IMDBNumber    string              `json:"imdbnumber,omitempty"`
}

// 6.20.8 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.MovieSet

type VideoDetailsMovieSet struct {
	VideoDetailsMedia
	SetId LibraryId `json:"setid"`
}

// 6.20.9 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.MovieSet.Extended

type VideoDetailsMovieSetExtended struct {
	VideoDetailsMovieSet
	Limits ListLimitsReturned   `json:"limits"`
	Movies []*VideoDetailsMovie `json:"movies,omitempty"`
}

// 6.20.10 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.MusicVideo

type VideoDetailsMusicVideo struct {
	VideoDetailsFile
	MusicVideoId LibraryId `json:"musicvideoid"` // non-optional
	Genre        []string  `json:"genre,omitempty"`
	Artist       []string  `json:"artist,omitempty"`
	Tag          []string  `json:"tag,omitempty"`
	Album        string    `json:"album,omitempty"`
	Track        int       `json:"track,omitempty"`
	Studio       []string  `json:"studio,omitempty"`
	Year         int       `json:"year,omitempty"`
}

// 6.20.11 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.Season

type VideoDetailsSeason struct {
	VideoDetailsBase
	Season          int       `json:"season"` // non-optional
	ShowTitle       string    `json:"showtitle,omitempty"`
	WatchedEpisodes int       `json:"watchedepisodes,omitempty"`
	TVShowId        LibraryId `json:"tvshowid,omitempty"`
	Episode         int       `json:"episode,omitempty"`
}

// 6.20.12 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Details.TVShow

type VideoDetailsTVShow struct {
	VideoDetailsItem
	TVShowId        LibraryId           `json:"tvshowid"` //non-optional
	SortTitle       string              `json:"sorttitle,omitempty"`
	MPAA            string              `json:"mpaa,omitempty"`
	Premiered       string              `json:"premiered,omitempty"`
	Year            int                 `json:"year,omitempty"`
	Episode         int                 `json:"episode,omitempty"`
	WatchedEpisodes int                 `json:"watchedepisodes,omitempty"`
	Votes           string              `json:"votes,omitempty"`
	Rating          float32             `json:"rating,omitempty"`
	Studio          []string            `json:"studio,omitempty"`
	Season          int                 `json:"season,omitempty"`
	Genre           []string            `json:"genre,omitempty"`
	Cast            []*VideoCastElement `json:"cast,omitempty"`
	EpisodeGuide    string              `json:"episodeguide,omitempty"`
	Tag             []string            `json:"tag,omitempty"`
	OriginalTitle   string              `json:"originaltitle,omitempty"`
	IMDBNumber      string              `json:"imdbnumber,omitempty"`
}

// 6.20.13 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Fields.Episode

type VideoFieldsEpisode string

const (
	EPISODE_FIELD_TITLE           VideoFieldsEpisode = "title"
	EPISODE_FIELD_PLOT            VideoFieldsEpisode = "plot"
	EPISODE_FIELD_VOTES           VideoFieldsEpisode = "votes"
	EPISODE_FIELD_RATING          VideoFieldsEpisode = "rating"
	EPISODE_FIELD_WRITER          VideoFieldsEpisode = "writer"
	EPISODE_FIELD_FIRST_AIRED     VideoFieldsEpisode = "firstaired"
	EPISODE_FIELD_PLAY_COUNT      VideoFieldsEpisode = "playcount"
	EPISODE_FIELD_RUNTIME         VideoFieldsEpisode = "runtime"
	EPISODE_FIELD_DIRECTOR        VideoFieldsEpisode = "director"
	EPISODE_FIELD_PRODUCTION_CODE VideoFieldsEpisode = "productioncode"
	EPISODE_FIELD_SEASON          VideoFieldsEpisode = "season"
	EPISODE_FIELD_EPISODE         VideoFieldsEpisode = "episode"
	EPISODE_FIELD_ORIGINAL_TITLE  VideoFieldsEpisode = "originaltitle"
	EPISODE_FIELD_SHOW_TITLE      VideoFieldsEpisode = "showtitle"
	EPISODE_FIELD_CAST            VideoFieldsEpisode = "cast"
	EPISODE_FIELD_STREAM_DETAILS  VideoFieldsEpisode = "streamdetails"
	EPISODE_FIELD_LAST_PLAYED     VideoFieldsEpisode = "lastplayed"
	EPISODE_FIELD_FAN_ART         VideoFieldsEpisode = "fanart"
	EPISODE_FIELD_THUMBNAIL       VideoFieldsEpisode = "thumbnail"
	EPISODE_FIELD_FILE            VideoFieldsEpisode = "file"
	EPISODE_FIELD_RESUME          VideoFieldsEpisode = "resume"
	EPISODE_FIELD_TV_SHOW_ID      VideoFieldsEpisode = "tvshowid"
	EPISODE_FIELD_DATE_ADDED      VideoFieldsEpisode = "dateadded"
	EPISODE_FIELD_UNIQUE_ID       VideoFieldsEpisode = "uniqueid"
	EPISODE_FIELD_ART             VideoFieldsEpisode = "art"
)

// 6.20.14 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Fields.Movie

type VideoFieldsMovie string

const (
	MOVIE_FIELD_TITLE          VideoFieldsMovie = "title"
	MOVIE_FIELD_GENRE          VideoFieldsMovie = "genre"
	MOVIE_FIELD_YEAR           VideoFieldsMovie = "year"
	MOVIE_FIELD_RATING         VideoFieldsMovie = "rating"
	MOVIE_FIELD_DIRECTOR       VideoFieldsMovie = "director"
	MOVIE_FIELD_TRAILER        VideoFieldsMovie = "trailer"
	MOVIE_FIELD_TAGLINE        VideoFieldsMovie = "tagline"
	MOVIE_FIELD_PLOT           VideoFieldsMovie = "plot"
	MOVIE_FIELD_PLOT_OUTLINE   VideoFieldsMovie = "plotoutline"
	MOVIE_FIELD_ORIGINAL_TITLE VideoFieldsMovie = "originaltitle"
	MOVIE_FIELD_LAST_PLAYED    VideoFieldsMovie = "lastplayed"
	MOVIE_FIELD_PLAYCOUNT      VideoFieldsMovie = "playcount"
	MOVIE_FIELD_WRITER         VideoFieldsMovie = "writer"
	MOVIE_FIELD_STUDIO         VideoFieldsMovie = "studio"
	MOVIE_FIELD_MPAA           VideoFieldsMovie = "mpaa"
	MOVIE_FIELD_CAST           VideoFieldsMovie = "cast"
	MOVIE_FIELD_COUNTRY        VideoFieldsMovie = "country"
	MOVIE_FIELD_IMDB_NUMBER    VideoFieldsMovie = "imdbnumber"
	MOVIE_FIELD_RUNTIME        VideoFieldsMovie = "runtime"
	MOVIE_FIELD_SET            VideoFieldsMovie = "set"
	MOVIE_FIELD_SHOWLINK       VideoFieldsMovie = "showlink"
	MOVIE_FIELD_STREAM_DETAILS VideoFieldsMovie = "streamdetails"
	MOVIE_FIELD_TOP250         VideoFieldsMovie = "top250"
	MOVIE_FIELD_VOTES          VideoFieldsMovie = "votes"
	MOVIE_FIELD_FAN_ART        VideoFieldsMovie = "fanart"
	MOVIE_FIELD_THUMBNAIL      VideoFieldsMovie = "thumbnail"
	MOVIE_FIELD_FILE           VideoFieldsMovie = "file"
	MOVIE_FIELD_SORT_TITLE     VideoFieldsMovie = "sorttitle"
	MOVIE_FIELD_RESUME         VideoFieldsMovie = "resume"
	MOVIE_FIELD_SET_ID         VideoFieldsMovie = "setid"
	MOVIE_FIELD_DATE_ADDED     VideoFieldsMovie = "dateadded"
	MOVIE_FIELD_TAG            VideoFieldsMovie = "tag"
	MOVIE_FIELD_ART            VideoFieldsMovie = "art"
)

// 6.20.15 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Fields.MovieSet

type VideoFieldsMovieSet string

// TODO

// 6.20.16 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Fields.MusicVideo

type VideoFieldsMusicVideo string

// TODO

// 6.20.17 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Fields.Season

type VideoFieldsSeason string

const (
	SEASON_FIELD_SEASON           VideoFieldsSeason = "season"
	SEASON_FIELD_SHOW_TITLE       VideoFieldsSeason = "showtitle"
	SEASON_FIELD_PLAY_COUNT       VideoFieldsSeason = "playcount"
	SEASON_FIELD_EPISODE          VideoFieldsSeason = "episode"
	SEASON_FIELD_FAN_ART          VideoFieldsSeason = "fanart"
	SEASON_FIELD_THUMBNAIL        VideoFieldsSeason = "thumbnail"
	SEASON_FIELD_TV_SHOW_ID       VideoFieldsSeason = "tvshowid"
	SEASON_FIELD_WATCHED_EPISODES VideoFieldsSeason = "watchedepisodes"
	SEASON_FIELD_ART              VideoFieldsSeason = "art"
)

// 6.20.18 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Fields.TVShow

type VideoFieldsTVShow string

const (
	TV_SHOW_FIELD_TITLE            VideoFieldsTVShow = "title"
	TV_SHOW_FIELD_GENRE            VideoFieldsTVShow = "genre"
	TV_SHOW_FIELD_YEAR             VideoFieldsTVShow = "year"
	TV_SHOW_FIELD_RATING           VideoFieldsTVShow = "rating"
	TV_SHOW_FIELD_PLOT             VideoFieldsTVShow = "plot"
	TV_SHOW_FIELD_STUDIO           VideoFieldsTVShow = "studio"
	TV_SHOW_FIELD_MPAA             VideoFieldsTVShow = "mpaa"
	TV_SHOW_FIELD_CAST             VideoFieldsTVShow = "cast"
	TV_SHOW_FIELD_PLAY_COUTN       VideoFieldsTVShow = "playcount"
	TV_SHOW_FIELD_EPISODE          VideoFieldsTVShow = "episode"
	TV_SHOW_FIELD_IMDB_NUMBER      VideoFieldsTVShow = "imdbnumber"
	TV_SHOW_FIELD_PREMERED_AT      VideoFieldsTVShow = "premiered"
	TV_SHOW_FIELD_VOTES            VideoFieldsTVShow = "votes"
	TV_SHOW_FIELD_LAST_PLAYED      VideoFieldsTVShow = "lastplayed"
	TV_SHOW_FIELD_FAN_ART          VideoFieldsTVShow = "fanart"
	TV_SHOW_FIELD_THUMBNAIL        VideoFieldsTVShow = "thumbnail"
	TV_SHOW_FIELD_FILE             VideoFieldsTVShow = "file"
	TV_SHOW_FIELD_ORIGINAL_TITLE   VideoFieldsTVShow = "originaltitle"
	TV_SHOW_FIELD_SORT_TITLE       VideoFieldsTVShow = "sorttitle"
	TV_SHOW_FIELD_EPISODE_GUIDE    VideoFieldsTVShow = "episodeguide"
	TV_SHOW_FIELD_SEASON           VideoFieldsTVShow = "season"
	TV_SHOW_FIELD_WATCHED_EPISODES VideoFieldsTVShow = "watchedepisodes"
	TV_SHOW_FIELD_DATE_ADDED       VideoFieldsTVShow = "dateadded"
	TV_SHOW_FIELD_TAG              VideoFieldsTVShow = "tag"
	TV_SHOW_FIELD_ART              VideoFieldsTVShow = "art"
)

// 6.20.19 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Resume

type VideoResume struct {
	Position float32 `json:"position,omitempty"`
	Total    float32 `json:"total,omitempty"`
}

// 6.20.20 http://kodi.wiki/view/JSON-RPC_API/v6#Video.Streams

type VideoStreamsVideo struct {
	Height   int     `json:"height"`
	Width    int     `json:"width"`
	Aspect   float32 `json:"aspect"`
	Codec    string  `json:"codec"`
	Duration int     `json:"duration"`
}

type VideoStreamsAudio struct {
	Channels int    `json:"channels,omitempty"`
	Language string `json:"language,omitempty"`
	Coded    string `json:"coded,omitempty"`
}

type VideoStreamsSubtitle struct {
	Language string `json:"language,omitempty"`
}

type VideoStreams struct {
	Video    []*VideoStreamsVideo    `json:"video,omitempty"`
	Audio    []*VideoStreamsAudio    `json:"audio,omitempty"`
	Subtitle []*VideoStreamsSubtitle `json:"subtitle,omitempty"`
}
