package weishi_watermark_remover

import "testing"

func TestAvailableLink(t *testing.T)  {
	/*
		处理以下格式
		URL1：https://h5.weishi.qq.com/weishi/feed/6YKtEB29l1K37v0Kx/wsfeed?wxplay=1
	*/
	link := "https://h5.weishi.qq.com/weishi/feed/6YKtEB29l1K37v0Kx/wsfeed?wxplay=1"

	url, err := WatermarkRemover(link)

	if err != nil {
		t.Fail()
	}

	if len(url) > 0 {
		t.Log("test pass")
		t.Log(url)
	}

}