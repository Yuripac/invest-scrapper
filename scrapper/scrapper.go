package scrapper

type Scrapper interface {
	FetchValue(string) (float64, error)
}
