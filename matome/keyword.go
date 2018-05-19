package matome

//+migu
type Keyword struct {
	ID      int     `migu:"autoincrement,index:ix_keyword_id"`
	SiteID  *int    `migu:"index:ix_keyword_site_id"`
	Keyword *string `migu:"index:ix_keyword_keyword"`
	Count   *int
}

