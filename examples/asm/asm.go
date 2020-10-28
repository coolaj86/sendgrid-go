package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go/v3"
)

// Createanewsuppressiongroup : Create a new suppression group
// POST /asm/groups
func Createanewsuppressiongroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "description": "Suggestions for products our users might like.", 
  "is_default": true, 
  "name": "Product Suggestions"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveinformationaboutmultiplesuppressiongroups : Retrieve information about multiple suppression groups
// GET /asm/groups
func Retrieveinformationaboutmultiplesuppressiongroups() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["id"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Updateasuppressiongroup : Update a suppression group.
// PATCH /asm/groups/{group_id}
func Updateasuppressiongroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}", host)
	request.Method = "PATCH"
	request.Body = []byte(` {
  "description": "Suggestions for items our users might like.", 
  "id": 103, 
  "name": "Item Suggestions"
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Getinformationonasinglesuppressiongroup : Get information on a single suppression group.
// GET /asm/groups/{group_id}
func Getinformationonasinglesuppressiongroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Deleteasuppressiongroup : Delete a suppression group.
// DELETE /asm/groups/{group_id}
func Deleteasuppressiongroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}", host)
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Addsuppressionstoasuppressiongroup : Add suppressions to a suppression group
// POST /asm/groups/{group_id}/suppressions
func Addsuppressionstoasuppressiongroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}/suppressions", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com", 
    "test2@example.com"
  ]
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveallsuppressionsforasuppressiongroup : Retrieve all suppressions for a suppression group
// GET /asm/groups/{group_id}/suppressions
func Retrieveallsuppressionsforasuppressiongroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}/suppressions", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Searchforsuppressionswithinagroup : Search for suppressions within a group
// POST /asm/groups/{group_id}/suppressions/search
func Searchforsuppressionswithinagroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}/suppressions/search", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "exists1@example.com", 
    "exists2@example.com", 
    "doesnotexists@example.com"
  ]
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Deleteasuppressionfromasuppressiongroup : Delete a suppression from a suppression group
// DELETE /asm/groups/{group_id}/suppressions/{email}
func Deleteasuppressionfromasuppressiongroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/groups/{group_id}/suppressions/{email}", host)
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveallsuppressions : Retrieve all suppressions
// GET /asm/suppressions
func Retrieveallsuppressions() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Addrecipientaddressestotheglobalsuppressiongroup : Add recipient addresses to the global suppression group.
// POST /asm/suppressions/global
func Addrecipientaddressestotheglobalsuppressiongroup() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions/global", host)
	request.Method = "POST"
	request.Body = []byte(` {
  "recipient_emails": [
    "test1@example.com", 
    "test2@example.com"
  ]
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// RetrieveaGlobalSuppression : Retrieve a Global Suppression
// GET /asm/suppressions/global/{email}
func RetrieveaGlobalSuppression() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions/global/{email}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// DeleteaGlobalSuppression : Delete a Global Suppression
// DELETE /asm/suppressions/global/{email}
func DeleteaGlobalSuppression() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions/global/{email}", host)
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Retrieveallsuppressiongroupsforanemailaddress : Retrieve all suppression groups for an email address
// GET /asm/suppressions/{email}
func Retrieveallsuppressiongroupsforanemailaddress() {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	host := "https://api.sendgrid.com"
	request := sendgrid.GetRequest(apiKey, "/v3/asm/suppressions/{email}", host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	// add your function calls here
}
