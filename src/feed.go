package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type video struct {
	URL         string      `json:"url"`
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
		"url":"lols",
		"comments": ["hi","lol"],
		"postedDate":13,
		"likes":1,
		"commentsNum":2		
	},
	{
		"url":"lols2",
		"comments": ["ahhh","lol22"],
		"postedDate":133,
		"likes":3,
		"commentsNum":131		
	}
]`
