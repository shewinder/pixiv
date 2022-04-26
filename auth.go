package pixiv

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// use refresh token to acquire a new bearer token
func auth(refreshToken string) (string, error) {
	clientId := "MOBrBDS8blbauoSck0ZfDbtuzpyT"
	clientSecret := "lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj"
	hashSecret := "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"
	localTime := time.Now().UTC().Format("2006-01-02T15:04:05+00:00")
	hashArr := md5.Sum([]byte(localTime + hashSecret))
	hashStr := hex.EncodeToString(hashArr[:])

	apiurl := "https://oauth.secure.pixiv.net/auth/token"
	data := url.Values{}
	data.Set("get_secure_url", "1")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	
	req, err := http.NewRequest("POST", apiurl, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-client-time", localTime)
	req.Header.Set("x-client-hash", hashStr)
	req.Header.Set("app-os", "ios")
	req.Header.Set("app-os-version", "14.6")
	req.Header.Set("user-agent", "PixivIOSApp/7.13.3 (iOS 14.6; iPhone13,2)")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return "", err
	}

	token, ok := m["access_token"].(string)
	if !ok {
		return "", err
	}
	return token, nil
}
