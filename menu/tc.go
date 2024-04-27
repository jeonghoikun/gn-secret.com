package menu

type tc struct {
	Name  string
	Price int
}

var TCs []*tc = []*tc{
	&tc{"룸 TC", 50_000},
	&tc{"아가씨 TC", 120_000},
	&tc{"디제이 TC", 100_000},
}
