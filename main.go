package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/xmlquery"
)

var cookie string = "foo"                                        // Todo Change
var incomesUrl string = "/users/42/basic_incomes"                // ToDo Change
var transfersUrl string = "/users/42/transfers"                  // ToDo Change
var pushoverKey string = "token=foo&user=bar&message=PalimPalim" // ToDo change

func main() {
	fmt.Println("Start")

	content := downloadstring()

	ok := strings.Contains(content, fmt.Sprintf("<form class=\"button_to\" method=\"post\" action=\"%v\">", incomesUrl))
	fmt.Println(ok)
	if ok {
		root, _ := xmlquery.Parse(strings.NewReader(content))
		formNode := xmlquery.FindOne(root, fmt.Sprintf("//form[@action='%v']", incomesUrl))
		authNode := xmlquery.FindOne(formNode, ".//input[@name='authenticity_token']")
		authToken := authNode.SelectAttr("value")
		fmt.Println(authToken)
	}
}

func downloadstring() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://palai.org%v", incomesUrl), nil)
	req.Header.Add("Cookie", cookie)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := string(body[:len(body)])

	fmt.Println(result) // Todo remove

	return result
}
