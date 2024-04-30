package siteconfig

import (
	"fmt"
	"os"
	"time"
)

type searchEngineConnection struct {
	Google string
}

type store struct {
	Location string
	Category string
	Name     string
	Address  string
	Tel      string
	Images   []string
}

func parseImages() []string {
	var ss []string
	entry, err := os.ReadDir("./static/img/index/store")
	if err != nil {
		panic(err)
	}
	for _, e := range entry {
		ss = append(ss, e.Name())
	}
	return ss
}

type config struct {
	ModeDev bool

	Port   uint32
	Domain string

	DatePublished time.Time
	DateModified  time.Time

	Title       string
	Description string
	Keywords    string
	Author      string

	Thumbnail string

	Store *store

	SearchEngineConnection *searchEngineConnection
}

var Config *config = &config{
	ModeDev: false,

	Port:   uint32(8021),
	Domain: "gn-secret.com",

	DatePublished: time.Date(2022, time.Month(11), 28, 0, 0, 0, 0, time.Local),
	DateModified:  time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.Local),

	Title:       "강남 풀싸롱 씨크릿 - Gangnam Room Salon Secret",
	Description: "강남 풀싸롱 씨크릿에서는 최상의 서비스와 프라이빗한 공간을 제공합니다. 고급스러운 인테리어와 전문 스태프가 만들어내는 최고의 환대를 경험하세요. 강남 최고의 명소에서 특별한 시간을 보내실 수 있습니다.",
	Keywords:    "강남 풀싸롱 씨크릿,강남 풀싸롱,풀싸롱 씨크릿,파티룸,주대,가격",
	Author:      "씨크릿",

	Thumbnail: "/static/img/index/thumbnail.png",

	Store: &store{
		Location: "강남",
		Category: "풀싸롱",
		Name:     "씨크릿",
		Address:  "서울 강남구 역삼동 719-18",
		Tel:      "010-7970-9057",
		Images:   parseImages(),
	},
	SearchEngineConnection: &searchEngineConnection{
		Google: "vca6lAnDIG-285kKpSxmm47sVeAOO5U-IM-_fvTAsaU",
	},
}

func Host() string {
	if Config.ModeDev {
		return fmt.Sprintf("http://localhost:%d", Config.Port)
	}
	return fmt.Sprintf("https://%s", Config.Domain)
}
