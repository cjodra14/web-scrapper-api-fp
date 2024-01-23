package services

import "web-scrapper-api-fp/api/models"

type ScrapperService interface {
	ScrapePosts(uri string) []models.Post
}
