package retrieve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getWordList(book Book, openbd OpenBDResponse) ([]string, error) {
	var query []string
	query = []string{
		book.Title,
		book.Author,
		book.Publisher,
	}
	if len(openbd) != 0 {
		query = append(
			query,
			openbd[0].Onix.DescriptiveDetail.TitleDetail.TitleElement.Subtitle.Content,
			returnCollateralTexts(openbd[0].Onix.CollateralDetail.TextContent),
			openbd[0].Hanmoto.Maegakinado,
			openbd[0].Hanmoto.Kaisetsu105W,
		)
	}
	word := strings.ReplaceAll(strings.Join(query, " "), "\n", "")
	req, err := http.NewRequest(
		"POST",
		"https://jlp.yahooapis.jp/MAService/V2/parse",
		bytes.NewBuffer([]byte(`{
			"id": "1234-1",
			"jsonrpc" : "2.0",
			"method" : "jlp.maservice.parse",
			"params" : {
				"q" : "`+word+`"
			}
		}`)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Yahoo AppID: "+os.Getenv("YAHOO_APP_ID"))
	client := new(http.Client)
	var wordList []string
	res, err := client.Do(req)
	if err != nil {
		return wordList, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return wordList, err
	}
	var data YahooMAApiResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return wordList, err
	}
	fmt.Println(data)
	for _, v := range data.Result.Tokens {
		if v[0] != " " && v[0] != "　" && v[3] != "動詞" && v[3] != "助詞" && v[3] != "接尾辞" && v[3] != "特殊" && v[4] != "数詞" {
			wordList = append(wordList, v[0], v[1])
		}
	}
	return wordList, nil
}

func returnCollateralTexts(contents []CollateralDetailTextContent) string {
	var newContentText string
	for _, s := range contents {
		newContentText += s.Text
	}
	return newContentText
}
