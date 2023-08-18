package domain

type Recognition struct {
	IdRecognition  int
	IdUser         int
	Key            string
	Name           string
	LocationLeft   string
	LocationTop    string
	LocationRight  string
	LocationBottom string
	Embeddings     string
	Distance       string
}
