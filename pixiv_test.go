package pixiv

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func init() {
	InitAuth("5_wLcosaJG103dcOR_ES8ybX3NTwKVxEjH7nFVF9YRA")
}

func TestAuth(t *testing.T) {
	token, _ := auth("5_wLcosaJG103dcOR_ES8ybX3NTwKVxEjH7nFVF9YRA")
	assert.NotEqual(t, "", token)
}

func TestRanking(t *testing.T) {
	illusts, _ := IllustRanking("day", "2022-4-25", "0")
	assert.Equal(t, 30, len(illusts.Illusts))
}

func TestUserIllusts(t *testing.T) {
	illusts, _ := UserIllusts("16731", "0", "illust")
	assert.Equal(t, 30, len(illusts.Illusts))
}

func TestIllustDetai(t *testing.T) {
	ill, _ := IllustDetail("87872301")
	assert.Equal(t, 87872301, ill.ID)
}

func TestUgoira(t *testing.T) {
	// InitAuth("5_wLcosaJG103dcOR_ES8ybX3NTwKVxEjH7nFVF9YRA")
	ug, _ := UgoiraMeta("87872301")
	assert.Equal(t, "https://i.pximg.net/img-zip-ugoira/img/2021/02/18/21/02/05/87872301_ugoira600x600.zip", ug.ZipUrls.Medium)
}

func TestUserBookmarkIllust(t *testing.T) {
	illusts, _ := UserBookmarkIllust("56339055", "public", "0")
	assert.Greater(t, len(illusts.Illusts), 0)
}