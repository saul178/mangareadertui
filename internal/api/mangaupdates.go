package api

/*
TODO:
- implement relevant api functions for the app
- implement api to search mangas
- get a specific manga from search results
- implement caching results
*/

const apiURL string = "https://api.mangaupdates.com/v1"
const searchSeriesEndpoint string = "/series/search"
const getSeries string = "/search/%v"

// WARNING: might not be needed ///////////////////////////////
type Response struct {
	Status  string                     `json:"status"`
	Reason  string                     `json:"reason"`
	Context map[string][]ContextErrors `json:"context"`
}

type ContextErrors struct {
	Index  int      `json:"index"`
	Errors []string `json:"errors"`
}

////////////////////////////////////////////////////////////

// NOTE: struct to send the payload to the search endpoint and get a list of mangas
// TODO: all marked as enumtype create a type enum of the related fields to
type SearchManga struct {
	Search                string   `json:"search"`
	AddedBy               int      `json:"added_by"`
	SType                 string   `json:"stype"`    // <- enumtype
	Licensed              string   `json:"licensed"` // <- enumtype
	Type                  []string `json:"type"`
	Year                  string   `json:"year"`
	FilterType            []string `json:"filter_types"`
	Category              []string `json:"category"`
	PubName               string   `json:"pubname"`
	Filters               []string `json:"filters"` // <- enumtype
	List                  string   `json:"list"`
	Page                  int      `json:"page"`
	PerPage               int      `json:"perpage"`
	Letter                string   `json:"letter"`
	Genre                 []string `json:"genre"`
	ExcludeGenre          []string `json:"exclude_genre"`
	OrderBy               string   `json:"orderby"` // <- enumtype
	Pending               bool     `json:"pending"`
	IncludeRankMetadata   bool     `json:"include_rank_metadata"`
	ExcludeFilteredGenres bool     `json:"exclude_filtered_genres"`
}

type SearchMangaResponse struct {
	TotalHits int             `json:"total_hits"`
	Page      int             `json:"page"`
	PerPage   int             `json:"per_page"`
	Results   []SearchResults `json:"results"`
}

type SearchResults struct {
	Record   MangaSeriesSummary `json:"record"`
	HitTitle string             `json:"hit_title"`
}

// NOTE: populates struct with a summary of the search results
type MangaSeriesSummary struct {
	ID             int           `json:"series_id"`
	Title          string        `json:"title"`
	Url            string        `json:"url"`
	Description    string        `json:"description"`
	Image          MangaCover    `json:"image"`
	Type           string        `json:"type"` // <- enumtype
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

// NOTE: main manga metadata when getting a specific series
type MangaSeries struct {
	MangaSeriesSummary                    // <- NOTE: adopts the fields from the series struct to avoid duplication of fields
	Associated         []AssociatedTitles `json:"associated"`
	Status             string             `json:"status"`
	Licensed           bool               `json:"licensed"`
	Completed          bool               `json:"completed"`
	Anime              Anime              `json:"anime"`
	RelatedSeries      []RelatedSeries    `json:"related_series"`
	Authors            []MangaAuthors     `json:"authors"`
	Publishers         []Publishers       `json:"publishers"`
	Publications       []Publications     `json:"publications"`
	Recommendations    []Recommendations  `json:"recommendations"`
	LastUpdated        LastUpdated        `json:"last_updated"`
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
	PublisherName string `json:"publisher_name"`
	PublisherID   int    `json:"publisher_id"`
	Url           string `json:"url"`
	Type          string `json:"type"` // <- enumtype
	Notes         string `json:"notes,omitempty"`
}

type MangaAuthors struct {
	Name     string `json:"name"`
	AuthorID int    `json:"author_id"`
	Url      string `json:"url"`
	Type     string `json:"type"` // <- enumtype
}

type RelatedSeries struct {
	RelationType      string `json:"relation_type"` // <- enumtype
	RelatedSeriesID   int    `json:"related_series_id"`
	RelatedSeriesName string `json:"related_series_name"`
	RelatedSeriesURL  string `json:"related_series_url"`
}

type Anime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type AssociatedTitles struct {
	Title string `json:"title"`
}
