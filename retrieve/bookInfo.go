package retrieve

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetBookInfo(isbn string) (BookInfo, OpenBDResponse, error) {
	info, data, err := getFromOpenBD(isbn)
	if err != nil {
		info, err = getFromNDLApiApi(isbn)
		if err != nil {
			return info, data, err
		}
	}
	return info, data, nil
}

func getFromOpenBD(isbn string) (BookInfo, OpenBDResponse, error) {
	res, err := http.Get("https://api.openbd.jp/v1/get?isbn=" + isbn)
	info := BookInfo{
		Isbn: isbn,
	}
	if err != nil {
		return info, nil, err
	} else if res.StatusCode != 200 {
		return info, nil, errors.New(res.Status)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return info, nil, err
	}

	var data OpenBDResponse
	if err := json.Unmarshal(body, &data); err != nil || len(data) == 0 {
		return info, nil, err
	}
	info = BookInfo{
		Title:       data[0].Summary.Title,
		Author:      data[0].Summary.Author,
		Publisher:   data[0].Summary.Publisher,
		Pubdate:     data[0].Summary.Pubdate,
		Lanove:      data[0].Hanmoto.Lanove,
		Tameshiyomi: data[0].Hanmoto.Hastameshiyomi,
		Isbn:        isbn,
	}
	return info, data, nil
}

func getFromNDLApiApi(isbn string) (BookInfo, error) {
	info := BookInfo{
		Isbn: isbn,
	}
	res, err := http.Get("https://iss.ndl.go.jp/api/sru?operation=searchRetrieve&query=isbn=" + isbn)
	if err != nil {
		return info, err
	} else if res.StatusCode != 200 {
		return info, errors.New(res.Status)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return info, err
	}

	var data NDLApiResponse
	if err := xml.Unmarshal(body, &data); err != nil || data.NumberOfRecords == "0" {
		return info, err
	}
	info = BookInfo{
		Title:     data.Records.Record[0].RecordData.Dc.Title,
		Author:    data.Records.Record[0].RecordData.Dc.Creator,
		Publisher: data.Records.Record[0].RecordData.Dc.Publisher,
		Isbn:      isbn,
	}
	return info, nil
}
