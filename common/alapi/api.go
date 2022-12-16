package alapi

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"wc_robot/common"
	"wc_robot/common/utils"
)

// 请求名言
func GetMingYan() (string, error) {
	uri, err := url.Parse(host + mingyanPath)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("token", common.GetConfig().ALAPI.Token)
	params.Add("format", "json")
	// 45种类型的名言，详见http://www.alapi.cn/api/view/7
	params.Add("typeid", strconv.Itoa(rand.Intn(46)))
	uri.RawQuery = params.Encode()
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}
	var ar AlapiResp
	if err := utils.ScanJson(resp, &ar); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s --%s", ar.Data.Content, ar.Data.Author), nil
}

// 请求情话
func GetQinghua() (string, error) {
	uri, err := url.Parse(host + qinghuaPath)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("token", common.GetConfig().ALAPI.Token)
	params.Add("format", "json")
	uri.RawQuery = params.Encode()
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}
	var ar AlapiResp
	if err := utils.ScanJson(resp, &ar); err != nil {
		return "", err
	}
	if ar.Code == Success || ar.Code == OverQPS || ar.Code == OverLimit {
		return ar.Data.Content, nil
	}
	return "", fmt.Errorf("请求响应失败, code:%d, desc:%s", ar.Code, GetCodeDesc(ar.Code))
}

// 请求心灵鸡汤
func GetSoul() (string, error) {
	uri, err := url.Parse(host + soulPath)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("token", common.GetConfig().ALAPI.Token)
	params.Add("format", "json")
	uri.RawQuery = params.Encode()
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}
	var ar AlapiResp
	if err := utils.ScanJson(resp, &ar); err != nil {
		return "", err
	}
	if ar.Code == Success || ar.Code == OverQPS || ar.Code == OverLimit {
		return ar.Data.Content, nil
	}
	return "", fmt.Errorf("请求响应失败, code:%d, desc:%s", ar.Code, GetCodeDesc(ar.Code))
}

//WeiboHotSearch 微博热搜
func WeiboHotSearch() (string, error) {
	uri, err := url.Parse(weiboHotSearch)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}
	var ar WeiboResp
	if err := utils.ScanJson(resp, &ar); err != nil {
		return "", err
	}
	if ar.Code == Success || ar.Code == OverQPS || ar.Code == OverLimit {
		return ar.toString(), nil
	}
	return "", fmt.Errorf("请求响应失败, code:%d, desc:%s", ar.Code, GetCodeDesc(ar.Code))
}
func Gjmj() (string, error) {
	uri, err := url.Parse(gjmj)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}
	var ar GjmjResp
	if err := utils.ScanJson(resp, &ar); err != nil {
		return "", err
	}
	if ar.Code == Success || ar.Code == OverQPS || ar.Code == OverLimit {
		return ar.toString(), nil
	}
	return "", fmt.Errorf("请求响应失败, code:%d, desc:%s", ar.Code, GetCodeDesc(ar.Code))
}
func Caijing() (string, error) {
	uri, err := url.Parse(caijing)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}
	var ar TopnewsResp
	if err := utils.ScanJson(resp, &ar); err != nil {
		return "", err
	}
	if ar.Code == Success || ar.Code == OverQPS || ar.Code == OverLimit {
		return ar.toString("财经新闻"), nil
	}
	return "", fmt.Errorf("请求响应失败, code:%d, desc:%s", ar.Code, GetCodeDesc(ar.Code))
}
func Topnews() (string, error) {
	uri, err := url.Parse(topnews)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}
	var ar TopnewsResp
	if err := utils.ScanJson(resp, &ar); err != nil {
		return "", err
	}
	if ar.Code == Success || ar.Code == OverQPS || ar.Code == OverLimit {
		return ar.toString("头条新闻"), nil
	}
	return "", fmt.Errorf("请求响应失败, code:%d, desc:%s", ar.Code, GetCodeDesc(ar.Code))
}
func Networkhot() (string, error) {
	uri, err := url.Parse(networkhot)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(uri.String())
	if err != nil {
		return "", err
	}
	var ar NetworkhotResp
	if err := utils.ScanJson(resp, &ar); err != nil {
		return "", err
	}
	if ar.Code == Success || ar.Code == OverQPS || ar.Code == OverLimit {
		return ar.toString(), nil
	}
	return "", fmt.Errorf("请求响应失败, code:%d, desc:%s", ar.Code, GetCodeDesc(ar.Code))
}

type WeiboResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		List []struct {
			Hotword    string `json:"hotword"`
			Hotwordnum string `json:"hotwordnum"`
			Hottag     string `json:"hottag"`
		} `json:"list"`
	} `json:"result"`
}

func (resp *WeiboResp) toString() string {

	result := "微博热榜\n"
	for i, item := range resp.Result.List {
		result += fmt.Sprintf("热榜%d【%s】\n %s\n ", i, item.Hotword, item.Hottag)
	}
	return result
}

type TopnewsResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		List []struct {
			Title       string `json:"title"`
			Hotnum      string `json:"hotnum"`
			Source      string `json:"source"`
			Url         string `json:"url"`
			PicUrl      string `json:"picUrl"`
			Description string `json:"description"`
		} `json:"list"`
	} `json:"result"`
}

func (resp *TopnewsResp) toString(title string) string {

	result := title + "\n"
	for i, item := range resp.Result.List {
		result += fmt.Sprintf("标题%d【%s】\n %s\n 链接：%s\n", i, item.Title, item.Description, item.Url)
	}
	return result
}

type NetworkhotResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		List []struct {
			Title  string `json:"title"`
			Hotnum string `json:"hotnum"`
			Digest string `json:"digest"`
		} `json:"list"`
	} `json:"result"`
}

func (resp *NetworkhotResp) toString() string {

	result := "全网热搜\n"
	for i, item := range resp.Result.List {
		result += fmt.Sprintf("标题%d【%s】\n %s", i, item.Title, item.Digest)
	}
	return result
}

type GjmjResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		Content string
		Source  string
	} `json:"result"`
}

func (resp *GjmjResp) toString() string {
	return fmt.Sprintf("%s(来自-%s)", resp.Result.Content, resp.Result.Source)
}
