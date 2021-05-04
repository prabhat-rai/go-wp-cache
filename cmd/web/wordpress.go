package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Post struct {
	ID      int    `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
}

type PostResponse struct {
	Found int    `json:"found"`
	Posts []Post `json:"posts"`
}

type CategoryTag struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	PostCount int    `json:"post_count"`
}

type CategoryResponse struct {
	Found      int           `json:"found"`
	Categories []CategoryTag `json:"categories"`
}

type TagResponse struct {
	Found int           `json:"found"`
	Tags  []CategoryTag `json:"tags"`
}

var responseArray map[string]string
var concurrentRequest bool

func fetchWordPressData(wpSite string, authorId string) map[string]string {
	responseArray = map[string]string{"posts": "", "categories": "", "tags": ""}
	concurrentRequest = false
	lengthOfResponse := len(responseArray)
	wpUrls := getWordPressUrls(wpSite, authorId)
	ch := make(chan string, 10)

	for key, url := range wpUrls {
		if concurrentRequest {
			go callWpApi(url, key, ch)
		} else {
			callWpApi(url, key, ch)
		}
	}

	for i := 0; i < lengthOfResponse; i++ {
		fmt.Printf("\n %s\n", <-ch)
	}

	return responseArray
}

func getWordPressUrls(wpSite string, authorId string) map[string]string {
	wordPressBaseUrl := "https://public-api.wordpress.com/rest/v1.1/sites/" + wpSite
	postQueryParams := "&number=3&status=publish&fields=ID,title,date,excerpt"
	categoryTagParams := "?order_by=count&order=DESC&number=&fields=slug,name,post_count"

	postUrl := wordPressBaseUrl + "/posts/?author=" + authorId + postQueryParams
	categoryUrl := wordPressBaseUrl + "/categories/" + categoryTagParams
	tagUrl := wordPressBaseUrl + "/tags/" + categoryTagParams

	return map[string]string{"posts": postUrl, "categories": categoryUrl, "tags": tagUrl}
}

func callWpApi(url string, callType string, ch chan string) {
	startTime := time.Now()
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

	apiResponse := convertResponse(response, callType)
	responseArray[callType] = apiResponse
	ch <- "Time taken for loading " + callType + " is " + time.Since(startTime).String()
}

func convertResponse(response []byte, callType string) string {
	var responseObject interface{}

	switch callType {
	case "posts":
		responseObject = PostResponse{}
		break
	case "categories":
		responseObject = CategoryResponse{}
		break
	case "tags":
		responseObject = TagResponse{}
		break
	}

	json.Unmarshal(response, &responseObject)
	responseStr := prettyPrint(responseObject)
	// fmt.Printf("\n\nNew API Response as struct %s\n", responseStr)
	return responseStr
}

func prettyPrint(responseObject interface{}) string {
	responseStr, _ := json.MarshalIndent(responseObject, "", "  ")
	return string(responseStr)
}
