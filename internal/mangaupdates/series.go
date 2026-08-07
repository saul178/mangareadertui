package mangaupdates

/*
TODO:
- implement relevant api functions for the app
- implement api to search mangas
- get a specific manga from search results
- implement caching results
- at some point move the models and defined types to their own file as this file gets bigger
*/

const BaseAPIURL string = "https://api.mangaupdates.com/v1"
const searchSeries string = "/series/search"
const getSeriesByID string = "/search/%v"

// WARNING: might not be needed ///////////////////////////////
type Response struct {
	StatusCode int                        `json:"-"`
	Status     string                     `json:"status"`
	Reason     string                     `json:"reason"`
	Context    map[string][]ContextErrors `json:"context"`
}

type ContextErrors struct {
	Index  int      `json:"index"`
	Errors []string `json:"errors"`
}

////////////////////////////////////////////////////////////

// NOTE: struct to send the payload to the search endpoint and get a list of mangas
// TODO: all marked as enumtype create a type enum of the related fields to
type SearchSeriesRequest struct {
	Search                string           `json:"search"`
	AddedBy               int              `json:"added_by"`
	SType                 SearchType       `json:"stype"`    // <- enumtype
	Licensed              IsSeriesLicensed `json:"licensed"` // <- enumtype
	Type                  []SeriesType     `json:"type"`
	Year                  string           `json:"year"`
	FilterType            []SeriesType     `json:"filter_types"` // <- NOTE: filters OUT certain series type EX: filter manga == to no manga on search response
	Category              []string         `json:"category"`     // <- NOTE: this will be an interesting field to filter from since it's voted by community
	PubName               string           `json:"pubname"`
	Filters               []SearchFilters  `json:"filters"` // <- enumtype
	List                  string           `json:"list"`
	Page                  int              `json:"page"`
	PerPage               int              `json:"perpage"`
	Letter                string           `json:"letter"`
	Genre                 []SeriesGenres   `json:"genre"` // <- TODO: find out all the genres from the site
	ExcludeGenre          []SeriesGenres   `json:"exclude_genre"`
	OrderBy               SearchOrderBy    `json:"orderby"` // <- enumtype
	Pending               bool             `json:"pending"`
	IncludeRankMetadata   bool             `json:"include_rank_metadata"`
	ExcludeFilteredGenres bool             `json:"exclude_filtered_genres"`
}

type SearchMangaResponse struct {
	TotalHits int             `json:"total_hits"`
	Page      int             `json:"page"`
	PerPage   int             `json:"per_page"`
	Results   []SearchResults `json:"results"`
}

type SearchResults struct {
	Record   MangaSeriesSearchSummary `json:"record"`
	HitTitle string                   `json:"hit_title"`
}

// NOTE: populates struct with a summary of the search results
type MangaSeriesSearchSummary struct {
	ID             int           `json:"series_id"`
	Title          string        `json:"title"`
	URL            string        `json:"url"`
	Description    string        `json:"description"`
	Image          MangaCover    `json:"image"`
	Type           SeriesType    `json:"type"` // <- enumtype
	Year           string        `json:"year"`
	BayesianRating float64       `json:"bayesian_rating"`
	RatingVotes    int           `json:"rating_votes"`
	Genres         []MangaGenres `json:"genres"`
	LatestChapter  int           `json:"latest_chapter"`
}

type MangaCover struct {
	Url    ImageURL `json:"url"`
	Height int      `json:"height"`
	Width  int      `json:"width"`
}

type ImageURL struct {
	Original string `json:"original"`
	Thumb    string `json:"thumb"`
}

type MangaGenres struct {
	Genre string `json:"genre"`
}

// NOTE: main manga metadata when getting a specific series by it's ID given from search results
type MangaSeries struct {
	MangaSeriesSearchSummary                    // <- NOTE: adopts the fields from the series struct to avoid duplication of fields
	Associated               []AssociatedTitles `json:"associated"`
	Status                   string             `json:"status"`
	Licensed                 bool               `json:"licensed"`
	Completed                bool               `json:"completed"`
	Anime                    AnimeDetails       `json:"anime"`
	RelatedSeries            []RelatedSeries    `json:"related_series"`
	Authors                  []SeriesAuthors    `json:"authors"`
	Publishers               []Publishers       `json:"publishers"`
	Publications             []Publications     `json:"publications"`
	Recommendations          []Recommendations  `json:"recommendations"`
	LastUpdated              LastUpdated        `json:"last_updated"`
	// Categories     []Categories       `json:"categories"` <- NOTE: idk if i care for this, can be added anytime
	// CategoryRecommendations []CategoryRecommendation `json:"category_recommendations"` <- NOTE: same as above
}

type LastUpdated struct {
	Timestamp int    `json:"timestamp"`
	AsRFC3339 string `json:"as_rfc3339"`
	AsString  string `json:"as_string"`
}

type Recommendations struct {
	SeriesName  string     `json:"series_name"`
	SeriesURL   string     `json:"series_url"`
	SeriesID    int        `json:"series_id"`
	SeriesImage MangaCover `json:"series_image"`
	Weight      int        `json:"weight"`
}

type Publications struct {
	PublicationName string `json:"publication_name"`
	PublisherName   string `json:"publisher_name"`
	PublisherID     int    `json:"publisher_id"`
}

type Publishers struct {
	PublisherName string        `json:"publisher_name"`
	PublisherID   int           `json:"publisher_id"`
	Url           string        `json:"url"`
	Type          PublisherType `json:"type"` // <- enumtype
	Notes         string        `json:"notes,omitempty"`
}

type SeriesAuthors struct {
	Name     string           `json:"name"`
	AuthorID int              `json:"author_id"`
	Url      string           `json:"url"`
	Type     SeriesAuthorType `json:"type"` // <- enumtype
}

type RelatedSeries struct {
	RelationType      SeriesRelationship `json:"relation_type"` // <- enumtype
	RelatedSeriesID   int                `json:"related_series_id"`
	RelatedSeriesName string             `json:"related_series_name"`
	RelatedSeriesURL  string             `json:"related_series_url"`
}

type AnimeDetails struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type AssociatedTitles struct {
	Title string `json:"title"`
}

// NOTE: these are defined types that the api expects from certain fields

// Search by title or description
type SearchType string

const (
	SearchByTitle       SearchType = "title"
	SearchByDescription SearchType = "description"
)

// Filter by licensed or unlicensed series
type IsSeriesLicensed string

const (
	IsLicensed    IsSeriesLicensed = "yes"
	IsNotLicensed IsSeriesLicensed = "no"
)

// Filter search series by "scanlated" "completed" etc
type SearchFilters string

const (
	SeriesScanlated    SearchFilters = "scanlated"
	SeriesCompleted    SearchFilters = "completed"
	SeriesOneshots     SearchFilters = "oneshots"
	SeriesNoOneshots   SearchFilters = "no_oneshots"
	SeriesSomeReleases SearchFilters = "some_releases"
	SeriesNoReleases   SearchFilters = "no_releases"
)

// Order a search by a given value
type SearchOrderBy string

const (
	OrderByScore              SearchOrderBy = "score"
	OrderByTitle              SearchOrderBy = "title"
	OrderByRank               SearchOrderBy = "rank"
	OrderByRating             SearchOrderBy = "rating"
	OrderByYear               SearchOrderBy = "year"
	OrderByDateAdded          SearchOrderBy = "date_added"
	OrderByWeekePosition      SearchOrderBy = "week_pos"
	OrderByMonthOnePosition   SearchOrderBy = "month1_pos"
	OrderByMonthThreePosition SearchOrderBy = "month3_pos"
	OrderByMonthSixPosition   SearchOrderBy = "month6_pos"
	OrderByYearPosition       SearchOrderBy = "year_pos"
	OrderByListReading        SearchOrderBy = "list_reading"
	OrderByListWish           SearchOrderBy = "list_wish"
	OrderByListComplete       SearchOrderBy = "list_complete"
	OrderByListUnfinished     SearchOrderBy = "list_unfinished"
)

// The type a valid series is defined as
type SeriesType string

const (
	ArtbookType    SeriesType = "Artbook"
	DoujinshiType  SeriesType = "Doujinshi"
	DramaCDType    SeriesType = "Drama CD"
	FilipinoType   SeriesType = "Filipino"
	IndonesianType SeriesType = "Indonesian"
	MangaType      SeriesType = "Manga"
	ManhwaType     SeriesType = "Manhwa"
	ManhuaType     SeriesType = "Manhua"
	NovelType      SeriesType = "Novel"
	OELType        SeriesType = "OEL"
	ThaiType       SeriesType = "Thai"
	VietnameseType SeriesType = "Vietnamese"
	MalaysianType  SeriesType = "Malaysian"
	NordicType     SeriesType = "Nordic"
	FrenchType     SeriesType = "French"
	SpanishType    SeriesType = "Spanish"
	GermanType     SeriesType = "German"
)

// The Genres stated by the api to filter searches
type SeriesGenres string

const (
	Josei         SeriesGenres = "Josei"
	Lolicon       SeriesGenres = "Lolicon"
	Seinen        SeriesGenres = "Seinen"
	Shotacon      SeriesGenres = "Shotacon"
	Shoujo        SeriesGenres = "Shoujo"
	ShoujoAi      SeriesGenres = "Shoujo Ai"
	Shounen       SeriesGenres = "Shounen"
	ShounenAi     SeriesGenres = "Shounen Ai"
	Yaoi          SeriesGenres = "Yaoi"
	Yuri          SeriesGenres = "Yuri"
	Action        SeriesGenres = "Action"
	Adult         SeriesGenres = "Adult"
	Adventure     SeriesGenres = "Adventure"
	Comedy        SeriesGenres = "Comedy"
	Doujinshi     SeriesGenres = "Doujinshi"
	Drama         SeriesGenres = "Drama"
	Ecchi         SeriesGenres = "Ecchi"
	Fantasy       SeriesGenres = "Fantasy"
	GenderBender  SeriesGenres = "Gender Bender"
	Harem         SeriesGenres = "Harem"
	Hentai        SeriesGenres = "Hentai"
	Historical    SeriesGenres = "Historical"
	Horror        SeriesGenres = "Horror"
	MartialArts   SeriesGenres = "MartialArts"
	Mature        SeriesGenres = "Mature"
	Mecha         SeriesGenres = "Mecha"
	Mystery       SeriesGenres = "Mystery"
	Psychological SeriesGenres = "Psychological"
	Romance       SeriesGenres = "Romance"
	SchoolLife    SeriesGenres = "School Life"
	SciFi         SeriesGenres = "Sci-fi"
	SliceOfLife   SeriesGenres = "Slice of Life"
	Smut          SeriesGenres = "Smut"
	Sports        SeriesGenres = "Sports"
	Supernatural  SeriesGenres = "Supernatural"
	Tragedy       SeriesGenres = "Tragedy"
)

// The role a creator has towards a series
type SeriesAuthorType string

const (
	SeriesAuthor SeriesAuthorType = "Author"
	SeriesArtist SeriesAuthorType = "Artist"
)

// The Publisher to a series
type PublisherType string

const (
	OriginalPublisher PublisherType = "Original"
	EnglishPublisher  PublisherType = "English"
)

// Defines the relationship a series has
type SeriesRelationship string

const (
	SeriesIsPrequel          SeriesRelationship = "Prequel"
	SeriesIsSequel           SeriesRelationship = "Sequel"
	SeriesIsSpinOff          SeriesRelationship = "Spin-Off"
	SeriesIsAdaptedFrom      SeriesRelationship = "Adapted From"
	SeriesIsAlternateVersion SeriesRelationship = "Alternate Version"
	SeriesIsPartOfAnthology  SeriesRelationship = "Part of Anthology"
	SeriesIsMainStory        SeriesRelationship = "Main Story"
	SeriesIsSideStory        SeriesRelationship = "Side Story"
	SeriesIsFullAnthology    SeriesRelationship = "Full Anthology"
	SeriesIsOther            SeriesRelationship = "Other"
)

