package matome

//+migu
type InspectionWord struct {
	ID   int     `migu:"autoincrement,index:ix_inspection_word_id"`
	Word *string `migu:"unique,size:30"`
}


