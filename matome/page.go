package matome
import "time"

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

