package pixiv

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func auth(refreshToken string) (token string) {
	// use refresh token to acquire a new bearer token
	clientId := "MOBrBDS8blbauoSck0ZfDbtuzpyT"
	clientSecret := "lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj"
	hashSecret := "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"
	localTime := time.Now().UTC().Format("2006-01-02T15:04:05+00:00")
	hashArr := md5.Sum([]byte(localTime + hashSecret))
	hashStr := hex.EncodeToString(hashArr[:])
	url := "https://oauth.secure.pixiv.net/auth/token"
	data := fmt.Sprintf("get_secure_url=1&client_id=%v&client_secret=%v&grant_type=refresh_token&refresh_token=%v", clientId, clientSecret, refreshToken)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		fmt.Println("req err")
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-client-time", localTime)
	req.Header.Set("x-client-hash", hashStr)
	req.Header.Set("app-os", "ios")
	req.Header.Set("app-os-version", "14.6")
	req.Header.Set("user-agent", "PixivIOSApp/7.13.3 (iOS 14.6; iPhone13,2)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get resp err", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp err", err.Error())
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("unmashal err", err)
		return
	}
	//fmt.Println(m)
	token, ok := m["access_token"].(string)
	if !ok {
		fmt.Println("error", m)
		return
	}
	return token
}
