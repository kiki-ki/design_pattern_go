package absf

import "fmt"

type listFactory struct{}

func (lf *listFactory) CreateLink(caption, url string) item {
	ll := new(listLink)
	ll.link = new(link)
	ll.baseItem = new(baseItem)
	ll.caption = caption
	ll.url = url
	return ll
}

func (lf *listFactory) CreateTray(caption string) tray {
	lt := new(listTray)
	lt.baseTray = new(baseTray)
	lt.baseItem = new(baseItem)
	lt.caption = caption
	return lt
}

func (lf *listFactory) CreatePage(title, author string) page {
	lp := new(listPage)
	lp.basePage = new(basePage)
	lp.title = title
	lp.author = author
	return lp
}

type listLink struct {
	*link
}

func (ll *listLink) makeHTML() string {
	return fmt.Sprintf("<li><a href=\"%s\">%s</a></li>\n", ll.url, ll.caption)
}

type listTray struct {
	*baseTray
}

func (lt *listTray) makeHTML() string {
	buf := "<li>\n"
	buf += fmt.Sprintf("%s\n", lt.caption)
	buf += "<ul>\n"
	for _, item := range lt.tray {
		buf += item.makeHTML()
	}
	buf += "</ul>\n"
	buf += "</li>\n"
	return buf
}

type listPage struct {
	*basePage
}

func (lp *listPage) makeHTML() string {
	buf := "<html>\n"
	buf += fmt.Sprintf("<head><title>%s</title></head>\n", lp.title)
	buf += "<body>\n"
	buf += fmt.Sprintf("<h1>%s</h1>", lp.title)
	buf += "<ul>"
	for _, item := range lp.content {
		buf += item.makeHTML()
	}
	buf += "</ul>"
	buf += fmt.Sprintf("<hr><adress>%s</adress>", lp.author)
	buf += "</body>\n</html>\n"
	return buf
}
