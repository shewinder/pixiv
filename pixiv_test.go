package pixiv

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	token := auth("5_wLcosaJG103dcOR_ES8ybX3NTwKVxEjH7nFVF9YRA")
	fmt.Println(token)
}

func TestRanking(t *testing.T) {
	InitAuth("5_wLcosaJG103dcOR_ES8ybX3NTwKVxEjH7nFVF9YRA")
	m,_ := IllustRanking("day", "2022-2-5", "1")
	fmt.Println(m)
}

func TestUserIllusts(t *testing.T) {
	InitAuth("5_wLcosaJG103dcOR_ES8ybX3NTwKVxEjH7nFVF9YRA")
	m, _ := UserIllusts("4588267", "10", "illust")
	fmt.Println(m)
}

func TestIllust(t *testing.T) {
	InitAuth("5_wLcosaJG103dcOR_ES8ybX3NTwKVxEjH7nFVF9YRA")
	ill, _ := IllustDetail("96273113")
	fmt.Println(ill.PageCount, ill.Title, ill.Caption, ill.Tags[0].Name)
}


//func TestToken(t testing.T) {
//	//
//} 
