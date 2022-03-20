package pixiv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const HOST = "https://app-api.pixiv.net"

type Token struct {
	accessToken string
	createTime  time.Time
}

func (tk *Token) isExpire() bool {
	return tk.createTime.Add(time.Second * 3500).Before(time.Now())
}

func (tk *Token) refresh() {
	tk.accessToken = auth(REFRESHTOKEN)
	tk.createTime = time.Now()
}

var REFRESHTOKEN string
var token Token
var client http.Client

func RefreshToken() {
	token.refresh()
}

func InitAuth(refreshToken string) {
	REFRESHTOKEN = refreshToken
}

func init() {
	client = http.Client{}
}

func pixGet(apiUrl string) (*http.Response, error) {
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("new request err", err)
		return nil, err
	}
	if token.accessToken == "" || token.isExpire() {
		token.refresh()
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token.accessToken))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get response err", err)
	}
	return resp, nil
}

// 作品排行
// mode: [day, week, month, day_male, day_female, week_original, week_rookie, day_manga]
// date: '2016-08-01'
// mode (Past): [day, week, month, day_male, day_female, week_original, week_rookie,
//               day_r18, day_male_r18, day_female_r18, week_r18, week_r18g]
func IllustRanking(mode string, date string, offset string) ([]*Illust, error) {
	apiUrl := fmt.Sprintf("%v/v1/illust/ranking", HOST)
	data := url.Values{}
	data.Set("mode", mode)
	data.Set("date", date)
	data.Set("offset", offset)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("parse url err", err)
		return nil, err
	}
	u.RawQuery = data.Encode()
	resp, err := pixGet(u.String())
	if err != nil {
		fmt.Println("get response err", err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err", err)
		return nil, err
	}
	type Resp struct {
		Illusts []*Illust `json:"illusts"`
	}
	var res Resp
	json.Unmarshal(b, &res)
	return res.Illusts, nil
}

func IllustDetail(illustId string) (*Illust, error) {
	apiUrl := fmt.Sprintf("%v/v1/illust/detail", HOST)
	data := url.Values{}
	data.Set("illust_id", illustId)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("parse url err", err)
		return nil, err
	}
	u.RawQuery = data.Encode()
	resp, err := pixGet(u.String())
	if err != nil {
		fmt.Println("get response err", err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err", err)
		return nil, err
	}
	//var m map[string]interface{}
	type Resp struct {
		Illust *Illust `json:"illust"`
	}
	var res Resp
	json.Unmarshal(b, &res)
	return res.Illust, nil
}

// 关注用户的新作
// restrict: [public, private]
func IllustFollow(restrict string) (map[string]interface{}, error) {
	apiUrl := fmt.Sprintf("%v/v2/illust/follow", HOST)
	data := url.Values{}
	data.Set("restrict", restrict)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("parse url err", err)
		return nil, err
	}
	u.RawQuery = data.Encode()
	resp, err := pixGet(u.String())
	if err != nil {
		fmt.Println("get response err", err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err", err)
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m, nil
}

func UserIllusts(userId string, offset string, type_ string) (map[string]interface{}, error) {
	apiUrl := fmt.Sprintf("%v/v1/user/illusts", HOST)
	data := url.Values{}
	data.Set("user_id", userId)
	data.Set("offset", offset)
	data.Set("type", type_)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("parse url err", err)
		return nil, err
	}
	u.RawQuery = data.Encode()
	resp, err := pixGet(u.String())
	if err != nil {
		fmt.Println("get response err", err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err", err)
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m, nil
}
