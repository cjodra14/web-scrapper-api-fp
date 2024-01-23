package scrapper

import (
	"web-scrapper-api-fp/api/models"

	"github.com/gocolly/colly"
)

type ScrappingService struct{}

func NewScrapper() *ScrappingService {
	return &ScrappingService{}
}

func (scrappingService *ScrappingService) ScrapePosts(uri string) []models.Post {
	var forumPosts []models.Post
	// initializing a Colly instance
	c := colly.NewCollector()
	// setting a valid User-Agent header
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// iterating over the list of pagination links to implement the crawling logic
	scrapePostPage(&forumPosts, uri)

	return forumPosts
}

func scrapePostPage(forumPosts *[]models.Post, postURL string) {
	// Initialize a new Colly collector for scraping individual posts
	postCollector := colly.NewCollector()
	// ... [set up postCollector, similar to the main collector]

	// Logic to scrape the contents of an individual post
	// scraping the post data
	postCollector.OnHTML("li.postcontainer", func(e *colly.HTMLElement) {
		forumPost := models.Post{}

		forumPost.MessageContent = e.ChildText("blockquote")
		forumPost.MessageTitle = e.ChildText(".title")
		forumPost.Author = e.ChildText(".username")

		*forumPosts = append(*forumPosts, forumPost)
	})

	// Visit the post page
	postCollector.Visit(postURL)
}
