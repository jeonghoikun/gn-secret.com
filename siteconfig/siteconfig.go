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
	DateModified:  time.Date(2024, time.Month(5), 14, 0, 0, 0, 0, time.Local),

	Title:       "강남 풀싸롱 세븐",
	Description: "강남 풀싸롱 세븐에서는 화려하고 역동적인 분위기 속에서 최상의 서비스와 고급스러운 경험을 제공합니다. 세븐의 럭셔리한 환경에서 특별한 시간을 보내보세요. 이곳은 방문객들에게 다양한 엔터테인먼트와 최고의 만족을 약속하는 장소로, 강남의 밤문화를 대표하는 고급스러운 공간입니다. 매력적인 분위기에서 편안하고 즐거운 시간을 보낼 수 있어 방문객들에게 잊지 못할 추억을 선사합니다.",
	Keywords:    "강남 풀싸롱 세븐,강남 풀싸롱,풀싸롱 세븐,파티룸,주대,가격",
	Author:      "세븐",

	Thumbnail: "/static/img/index/thumbnail.png",

	Store: &store{
		Location: "강남",
		Category: "풀싸롱",
		Name:     "세븐",
		Address:  "서울 강남구 역삼동 719-18",
		Tel:      "010-5860-0075",
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
