package alapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestWeiboHotSearch(t *testing.T) {
	WeiboHotSearch()
}

func TestWeiboHot(t *testing.T) {
	var resp WeiboResp
	err := json.Unmarshal([]byte(text), &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}

var text = `{
  "code": 200,
  "msg": "success",
  "result": {
    "list": [
      {
        "hotword": "张靓颖道歉",
        "hotwordnum": " 3168645",
        "hottag": "新"
      },
      {
        "hotword": "女子发烧敷20分钟面膜揭下变3D立体",
        "hotwordnum": " 1406032",
        "hottag": "热"
      }
    ]
  }
}`
