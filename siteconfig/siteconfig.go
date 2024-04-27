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

	Port:   uint32(8012),
	Domain: "gn-secret.com",

	DatePublished: time.Date(2022, time.Month(11), 28, 0, 0, 0, 0, time.Local),
	DateModified:  time.Date(2024, time.Month(1), 5, 0, 0, 0, 0, time.Local),

	Title:       "강남 퍼펙트 하이퍼블릭 - Gangnam High Public Perfect",
	Description: "강남 퍼펙트 하이퍼블릭 공식홈페이지. 퍼펙트 퍼블릭은 엘리에나 호텔 부속 초호화 하이퍼블릭 입니다. 프라이빗한 파티 공간을 제공 합니다. 시스템 소개 및 서비스 가격(주대) 안내. 고객님의 특별한 하루를 더욱 특별하게 만들어드리는 퍼펙트에서 즐거운 추억 만들어가세요.",
	Keywords:    "강남 퍼펙트 하이퍼블릭,강남 퍼펙트,퍼펙트 하이퍼블릭,파티룸,주대,가격",
	Author:      "퍼펙트",

	Thumbnail: "/static/img/index/thumbnail.png",

	Store: &store{
		Location: "강남",
		Category: "하이퍼블릭",
		Name:     "퍼펙트",
		Address:  "서울시 강남구 논현동 151-30 엘리에나호텔 지하",
		Tel:      "010-6590-7589",
		Images:   parseImages(),
	},

	SearchEngineConnection: &searchEngineConnection{
		Google: "MymVF9NWXKDeOkux9wwNkGAGfahosPwju227NyeStwg",
	},
}

func Host() string {
	if Config.ModeDev {
		return fmt.Sprintf("http://localhost:%d", Config.Port)
	}
	return fmt.Sprintf("https://%s", Config.Domain)
}
