package pixiv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	return tk.createTime.Add(time.Second * 3000).Before(time.Now())
}

func (tk *Token) refresh() {
	tk.accessToken, _ = auth(refresh)
	tk.createTime = time.Now()
}

var refresh string
var token Token
var client http.Client

func RefreshToken() {
	token.refresh()
}

func InitAuth(refreshToken string) {
	refresh = refreshToken
}

func init() {
	client = http.Client{}
}

func Get(path string, params map[string]string) ([]byte, error) {
	apiurl := fmt.Sprintf("%v%v", HOST, path)
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	u, err := url.ParseRequestURI(apiurl + "?" + data.Encode())
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}
	if token.accessToken == "" || token.isExpire() {
		token.refresh()
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token.accessToken))
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")
	resp, err := client.Do(req)
	if err != nil {
		log.Default().Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Default().Println(resp.Status, string(b))
		return nil, fmt.Errorf("%v", resp.Status+string(b))
	}
	return b, nil
}

// 作品排行
// mode: [day, week, month, day_male, day_female, week_original, week_rookie, day_manga]
// date: '2016-08-01'
// mode (Past): [day, week, month, day_male, day_female, week_original, week_rookie,
//               day_r18, day_male_r18, day_female_r18, week_r18, week_r18g]
func IllustRanking(mode string, date string, offset string) (*Illusts, error) {
	path := "/v1/illust/ranking"
	params := map[string]string{
		"mode":   mode,
		"date":   date,
		"offset": offset,
	}
	b, err := Get(path, params)
	if err != nil {
		return nil, err
	}
	var illusts Illusts
	err = json.Unmarshal(b, &illusts)
	return &illusts, err
}

// 插画详情
func IllustDetail(illustId string) (*Illust, error) {
	path := "/v1/illust/detail"
	params := map[string]string{
		"illust_id": illustId,
	}
	b, err := Get(path, params)
	if err != nil {
		return nil, err
	}

	type Resp struct {
		Illust *Illust `json:"illust"`
	}
	var res Resp
	err = json.Unmarshal(b, &res)
	return res.Illust, err
}

// 关注用户的新作
// restrict: [public, private]
func IllustFollow(restrict string) (*Illusts, error) {
	path := "/v2/illust/follow"
	params := map[string]string{
		"restrict": restrict,
	}
	b, err := Get(path, params)
	if err != nil {
		return nil, err
	}

	var illusts Illusts
	err = json.Unmarshal(b, &illusts)
	return &illusts, err
}

// 用户插画
func UserIllusts(userId string, offset string, type_ string) (*Illusts, error) {
	path := "/v1/user/illusts"
	params := map[string]string{
		"user_id": userId,
		"offset":  offset,
		"type":    type_,
	}
	b, err := Get(path, params)
	if err != nil {
		return nil, err
	}
	var illusts Illusts
	err = json.Unmarshal(b, &illusts)
	return &illusts, err
}

// 动图信息， 包括图片压缩包下载网址、帧数信息
func UgoiraMeta(illustId string) (*Ugoira, error) {
	path := "/v1/ugoira/metadata"
	params := map[string]string{
		"illust_id": illustId,
	}
	b, err := Get(path, params)
	if err != nil {
		return nil, err
	}
	type Resp struct {
		UgoiraMetadata Ugoira `json:"ugoira_metadata"`
	}
	var resp Resp
	err = json.Unmarshal(b, &resp)
	return &resp.UgoiraMetadata, err
}

// 用户收藏作品列表
func UserBookmarkIllust(userId string, restrict string, offset string) (*Illusts, error) {
	path := "/v1/user/bookmarks/illust"
	params := map[string]string{
		"user_id":  userId,
		"restrict": restrict,
		"offset":   offset,
		"filter":   "for_ios",
	}
	b, err := Get(path, params)
	if err != nil {
		return nil, err
	}
	var illusts Illusts
	err = json.Unmarshal(b, &illusts)
	return &illusts, err
}
