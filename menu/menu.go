package menu

type menu struct {
	Name  string
	Price string
}

var Menus []*menu = []*menu{
	&menu{"1부 주대", "문의"},
	&menu{"2부 주대", "문의"},
	&menu{"TC", "400,000"},
	&menu{"RT", "문의"},
}
