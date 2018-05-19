package matome
import "time"

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


