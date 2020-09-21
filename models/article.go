// models.article.go

package models

import "errors"

type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	BPM       int    `json:"bpm"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
}

// ArticleList For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var ArticleList = []Article{
	{ID: 1,
		Title:     "Article 1",
		Content:   "Article 1 body",
		Index:     0,
		Timestamp: "2020-08-21 05:26:22.774889 -0800 IST m=+0.001658999",
		BPM:       0,
		Hash:      "",
		PrevHash:  "",
	},
	{
		ID:        2,
		Title:     "Article 2",
		Content:   "Article 2 body",
		Index:     1,
		Timestamp: "2020-08-21 05:23:45.345889 -0800 IST m=+0.234658999",
		BPM:       22,
		Hash:      "45c9a6614fccd4f9592d8283a4f25bff84076fd43ee9f90eaa07746ebbed02ca",
		PrevHash:  "",
	},
	{
		ID:        3,
		Title:     "Article 2",
		Content:   "Article 2 body",
		Index:     2,
		Timestamp: "2020-08-21 05:24:21.927845 -0800 IST m=+0.299308376",
		BPM:       25,
		Hash:      "b6fecd38e1af90e1d117bfe6694e0ab54d4c9447db0c8412fe66de3bd9e43b9f",
		PrevHash:  "45c9a6614fccd4f9592d8283a4f25bff84076fd43ee9f90eaa07746ebbed02ca",
	},
}

// Return a list of all the articles
func GetAllArticles() []Article {
	return ArticleList
}

// Fetch an article based on the ID supplied
func GetArticleByID(id int) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// Create a new article with the title and content provided
func CreateNewArticle(title, content string) (*Article, error) {
	// Set the ID of a new article to one more than the number of articles
	a := Article{ID: len(ArticleList) + 1, Title: title, Content: content}

	// Add the article to the list of articles
	ArticleList = append(ArticleList, a)

	return &a, nil
}
