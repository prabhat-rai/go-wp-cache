package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//type Post struct {
//	ID      int    `json:"id"`
//	Date    string `json:"date"`
//	Title   string `json:"title"`
//	Excerpt string `json:"excerpt"`
//}

type PostResponse struct {
	Found int    `json:"found"`
	Posts []interface{} `json:"posts"`
}

//type CategoryTag struct {
//	Name      string `json:"name"`
//	Slug      string `json:"slug"`
//	PostCount int    `json:"post_count"`
//}
//
//type CategoryResponse struct {
//	Found      int           `json:"found"`
//	Categories []interface{} `json:"categories"` // Try tonight
//}
//
//type TagResponse struct {
//	Found int           `json:"found"`
//	Tags  []interface{} `json:"tags"`
//}

func fetchWordpressData(wpSite string, authorId string) []string {
	WordpressBaseUrl := "https://public-api.wordpress.com/rest/v1.1/sites/" + wpSite
	postQueryParams := "&number=3&status=publish&fields=ID,title,date,excerpt"
	categoryTagParams := "?order_by=count&order=DESC&number=&fields=slug,name,post_count"

	postUrl := WordpressBaseUrl + "/posts/?author=" + authorId + postQueryParams
	categoryUrl := WordpressBaseUrl + "/categories/" + categoryTagParams
	tagUrl := WordpressBaseUrl + "/tags/" + categoryTagParams

	response := callWpApi(postUrl, "posts")
	categoryResponse := callWpApi(categoryUrl, "categories")
	tagResponse := callWpApi(tagUrl, "tags")

	responseArray := []string{
		response,
		categoryResponse,
		tagResponse,
	}
	return responseArray
}

func callWpApi(url string, callType string) string {
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

	return convertResponse(response, callType)
}

func convertResponse(response []byte, callType string) string {
	var responseObject interface{}

	responseObject = PostResponse{}

	//switch callType {
	//	case "posts":
	//		responseObject = PostResponse{}
	//		break
	//	case "categories":
	//		responseObject = CategoryResponse{}
	//		break
	//	case "tags":
	//		responseObject = TagResponse{}
	//		break
	//}

	json.Unmarshal(response, &responseObject)
	responseStr := prettyPrint(responseObject)
	fmt.Printf("\n\nNew API Response as struct %s\n", responseStr)
	return responseStr
}

func prettyPrint(responseObject interface{}) string {
	responseStr, _ := json.MarshalIndent(responseObject, "", "  ")
	return string(responseStr)
}
