package pixiv

import "time"

type Illust struct {
	Caption              string    `json:"caption"`
	CommentAccessControl int       `json:"comment_access_control"`
	CreateDate           time.Time `json:"create_date"`
	Height               int       `json:"height"`
	ID                   int       `json:"id"`
	ImageUrls            struct {
		Large        string `json:"large"`
		Medium       string `json:"medium"`
		SquareMedium string `json:"square_medium"`
	} `json:"image_urls"`
	IsBookmarked bool `json:"is_bookmarked"`
	IsMuted      bool `json:"is_muted"`
	MetaPages    []struct {
		ImageUrls struct {
			Large        string `json:"large"`
			Medium       string `json:"medium"`
			Original     string `json:"original"`
			SquareMedium string `json:"square_medium"`
		} `json:"image_urls"`
	} `json:"meta_pages"`
	MetaSinglePage struct {
		OriginalImageURL string `json:"original_image_url"`
	} `json:"meta_single_page"`
	PageCount   int         `json:"page_count"`
	Restrict    int         `json:"restrict"`
	SanityLevel int         `json:"sanity_level"`
	Series      interface{} `json:"series"`
	Tags        []struct {
		Name           string      `json:"name"`
		TranslatedName interface{} `json:"translated_name"`
	} `json:"tags"`
	Title          string        `json:"title"`
	Tools          []interface{} `json:"tools"`
	TotalBookmarks int           `json:"total_bookmarks"`
	TotalComments  int           `json:"total_comments"`
	TotalView      int           `json:"total_view"`
	Type           string        `json:"type"`
	User           struct {
		Account          string `json:"account"`
		ID               int    `json:"id"`
		IsFollowed       bool   `json:"is_followed"`
		Name             string `json:"name"`
		ProfileImageUrls struct {
			Medium string `json:"medium"`
		} `json:"profile_image_urls"`
	} `json:"user"`
	Visible   bool `json:"visible"`
	Width     int  `json:"width"`
	XRestrict int  `json:"x_restrict"`
}

type Illusts struct {
	Illusts []*Illust `json:"illusts"`
	NextUrl string    `json:"next_url"`
}


type ZipUrls struct {
	Medium string `json:"medium"`
}
type Frames struct {
	File  string `json:"file"`
	Delay int    `json:"delay"`
}
type Ugoira struct {
	ZipUrls ZipUrls  `json:"zip_urls"`
	Frames  []Frames `json:"frames"`
}
