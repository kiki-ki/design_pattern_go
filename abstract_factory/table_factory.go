package absf

import "fmt"

type tableFactory struct{}

func (tf *tableFactory) CreateLink(caption, url string) item {
	tl := new(tableLink)
	tl.link = new(link)
	tl.baseItem = new(baseItem)
	tl.caption = caption
	tl.url = url
	return tl
}

func (tf *tableFactory) CreateTray(caption string) tray {
	tt := new(tableTray)
	tt.baseTray = new(baseTray)
	tt.baseItem = new(baseItem)
	tt.caption = caption
	return tt
}

func (tf *tableFactory) CreatePage(title, author string) page {
	tp := new(tablePage)
	tp.basePage = new(basePage)
	tp.title = title
	tp.author = author
	return tp
}

type tableLink struct {
	*link
}

func (tl *tableLink) makeHTML() string {
	return fmt.Sprintf("<td><a href=%s>%s</a></td>\n", tl.url, tl.caption)
}

type tableTray struct {
	*baseTray
}

func (tt *tableTray) makeHTML() string {
	buf := "<td>\n"
	buf += "<table width=\"100%\" border=\"1\"><tr>\n"
	buf += fmt.Sprintf("<td bgcolor=\"#cccccc\" algin=\"center\" colsapn=\"%d\"><b>%s</b></td>\n", len(tt.tray), tt.caption)
	buf += "</tr>\n"
	buf += "<tr>\n"
	for _, item := range tt.tray {
		buf += item.makeHTML()
	}
	buf += "</tr></table>\n"
	buf += "</td>\n"
	return buf
}

type tablePage struct {
	*basePage
}

func (tp *tablePage) makeHTML() string {
	buf := "<html>\n"
	buf += fmt.Sprintf("<head><title>%s</title></head>\n", tp.title)
	buf += "<body>\n"
	buf += fmt.Sprintf("<h1>%s</h1>", tp.title)
	buf += "<table width=\"80%\" border=\"3\">\n"
	for _, item := range tp.content {
		buf += fmt.Sprintf("<tr>%s</tr>", item.makeHTML())
	}
	buf += "</table>"
	buf += fmt.Sprintf("<hr><adress>%s</adress>", tp.author)
	buf += "</body>\n</html>\n"
	return buf
}
