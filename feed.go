package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type video struct {
	URL         string      `json:"url"`
	Rendition   string      `json:"rendition"`
	Comments    []string    `json:"comments"`
	PostedDate  json.Number `json:"postedDate"`
	Likes       json.Number `json:"likes"`
	CommentsNum json.Number `json:"commentsNum"`
}

func Feed(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	encoder := json.NewEncoder(w)

	var mockVideos []video
	err := json.Unmarshal([]byte(mockFeed), &mockVideos)
	if err != nil {
		log.Printf("Error decoding mock vidoes, Error=%v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = encoder.Encode(mockVideos)
	if err != nil {
		log.Printf("Error encoding vidoes, Error=%v", err)
	}
}

const mockFeed = `[
	{
		"url":"https://s3.amazonaws.com/jelo-videos/a.mp4",
		"rendition":"https://s3.amazonaws.com/jelo-renditions/a.jpg",
		"comments": ["a","lol"],
		"postedDate":13,
		"likes":1,
		"commentsNum":2		
	},
	{
		"url":"https://s3.amazonaws.com/jelo-videos/b.mp4",
		"rendition":"https://s3.amazonaws.com/jelo-renditions/b.jpg",
		"comments": ["b","lol22"],
		"postedDate":133,
		"likes":3,
		"commentsNum":131		
	},
	{
		"url":"https://s3.amazonaws.com/jelo-videos/c.mp4",
		"rendition":"https://s3.amazonaws.com/jelo-renditions/c.jpg",
		"comments": ["c","lol22"],
		"postedDate":133,
		"likes":3,
		"commentsNum":131		
	},
	{
		"url":"https://s3.amazonaws.com/jelo-videos/d.mp4",
		"rendition":"https://s3.amazonaws.com/jelo-renditions/d.jpg",
		"comments": ["d","lol22"],
		"postedDate":133,
		"likes":3,
		"commentsNum":131		
	}
]`
