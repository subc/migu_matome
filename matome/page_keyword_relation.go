package matome

//+migu
type PageKeywordRelation struct {
	ID        int `migu:"autoincrement,index:ix_page_keyword_relation_id"`
	PageID    *int
	KeywordID *int `migu:"index:ix_page_keyword_relation_keyword_id"`
}


