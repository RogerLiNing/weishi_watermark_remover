package weishi_watermark_remover

import (
	"strings"
)

func WatermarkRemover(url string)(string, error)  {

	// 提取视频ID https://h5.weishi.qq.com/weishi/feed/6YKtEB29l1K37v0Kx/wsfeed?wxplay=1

	url = strings.Split(url, "?")[0]
	paths := strings.Split(url, "/")
	feedId := paths[len(paths) - 2]

	videoLink := ""
	var err error

	if len(feedId) > 0 {
		videoLink, err = getVideoInfo(feedId)
	}


	return videoLink, err
}
