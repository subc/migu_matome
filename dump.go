package matome
import (
	"time"
)

//+migu
type InspectionWord struct {
	ID   int     `migu:"autoincrement,index:ix_inspection_word_id"`
	Word *string `migu:"unique,size:30"`
}

//+migu
type Keyword struct {
	ID      int     `migu:"autoincrement,index:ix_keyword_id"`
	SiteID  *int    `migu:"index:ix_keyword_site_id"`
	Keyword *string `migu:"index:ix_keyword_keyword"`
	Count   *int
}

//+migu
type Page struct {
	ID        int `migu:"autoincrement,index:ix_page_id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	SiteID    *int `migu:"index:ix_page_site_id"`
	DatID     *int `migu:"index:ix_page_dat_id"`
	Page      *string
	ViewCount *int `migu:"index:ix_page_view_count"`
	PageTop   *string
	Type      *int    `migu:"index:ix_page_type"`
	Keywords  *string `migu:"size:1000"`
	StartAt   *time.Time
	DeleteAt  *time.Time
}

//+migu
type PageKeywordRelation struct {
	ID        int `migu:"autoincrement,index:ix_page_keyword_relation_id"`
	PageID    *int
	KeywordID *int `migu:"index:ix_page_keyword_relation_keyword_id"`
}

//+migu
type Site struct {
	ID                   int `migu:"autoincrement,index:ix_site_id"`
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
	Name                 *string `migu:"size:50"`
	ShortName            *string `migu:"size:20"`
	Title                *string `migu:"index:ix_site_title,size:10"`
	URL                  *string `migu:"size:200"`
	BackgroundImageCount *int
	AdType               *int
}

//+migu
type Site3 struct {
	ID                   int `migu:"pk"`
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
	Name                 *string `migu:"size:50"`
	ShortName            *string `migu:"size:20"`
	Title                *string `migu:"index:ix_site_title,size:10"`
	URL                  *string `migu:"size:200"`
	BackgroundImageCount *int
	AdType               *int
}
