package weishi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type JsonData struct {
	Data Data `json:"data"`
}

type Data struct {
	Feeds []Feed `json:"feeds"`
}

type Feed struct {
	VideoUrl string `json:"video_url"`
}
/*
{
    "feedid": "6YKtEB29l1K37v0Kx",
    "recommendtype": 0,
    "datalvl": "all",
    "_weishi_mapExt": {}
}
*/


func getVideoInfo(videoId string) (string, error) {

	client := &http.Client{}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	url := "https://h5.weishi.qq.com/webapp/json/weishi/WSH5GetPlayPage"

	postData := fmt.Sprintf(`{"feedid": "%s","recommendtype": 0,"datalvl": "all","_weishi_mapExt": {}}`, videoId)

	request, err := http.NewRequest("POST", url, strings.NewReader(postData))

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Content-type", "application/json")

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)

	}

	jsonData := JsonData{}
	json.Unmarshal(body, &jsonData)

	var videolInk string

	if len(jsonData.Data.Feeds) > 0 {
		videolInk = jsonData.Data.Feeds[0].VideoUrl
	}


	return videolInk, nil

}
