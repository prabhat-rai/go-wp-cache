package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	ID 		int 	`json:"id"`
	Date 	string 	`json:"date"`
	Title 	string 	`json:"title"`
	Excerpt string 	`json:"excerpt"`
}

type PostResponse struct {
	Found	int 	`json:"found"`
	Posts   []Post 	`json:"posts"`
	Status 	int    	`json:"status"`
}

type CategoryTag struct {
	Name 		string 	`json:"name"`
	Slug 		string 	`json:"slug"`
	PostCount 	int 	`json:"post_count"`
}

type CategoryResponse struct {
	Found     	int 			`json:"found"`
	Categories  []CategoryTag 	`json:"categories"`
}

type TagResponse struct {
	Found	int 			`json:"found"`
	Tags	[]CategoryTag 	`json:"tags"`
}

func fetchWordpressData(wpSite string, authorId string) []string  {
	WordpressBaseUrl := "https://public-api.wordpress.com/rest/v1.1/sites/" + wpSite
	categoryTagFieldList := "?order_by=count&order=DESC&number=&fields=slug,name,post_count"

	postUrl := WordpressBaseUrl + "/posts/?number=3&author=" + authorId + "&status=publish&fields=ID,title,date,excerpt"
	categoryUrl := WordpressBaseUrl + "/categories/" + categoryTagFieldList
	tagUrl := WordpressBaseUrl + "/tags/" + categoryTagFieldList

	response := string(callWpApi(postUrl, "posts"))
	categoryResponse := string(callWpApi(categoryUrl, "categories"))
	tagResponse := string(callWpApi(tagUrl, "tags"))

	responseArray := []string {
		response,
		categoryResponse,
		tagResponse,
	}
	return responseArray
}

func callWpApi (url string, callType string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	// logWpResponse(response, callType)

	return response
}

func logWpResponse(response []byte, callType string) {
	switch callType {
		case "posts":
			var responseObject PostResponse
			json.Unmarshal(response, &responseObject)
			fmt.Printf("\n\nAPI Response as struct %+v\n", responseObject)
			break
		case "categories":
			var responseObject CategoryResponse
			json.Unmarshal(response, &responseObject)
			fmt.Printf("\n\nAPI Response as struct %+v\n", responseObject)
			break
		case "tags":
			var responseObject TagResponse
			json.Unmarshal(response, &responseObject)
			fmt.Printf("\n\nAPI Response as struct %+v\n", responseObject)
			break
	}
}