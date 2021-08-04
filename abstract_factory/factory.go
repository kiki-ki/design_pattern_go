package absf

import (
	"fmt"
	"os"
)

type factoryType string

const (
	list factoryType = "list"
	table factoryType = "table"
)

func New(factoryType factoryType) factory {
	switch factoryType {
	case list:
		return &listFactory{}
	case table:
		return &tableFactory{}
	default:
		panic(fmt.Sprintf("%s is not allowed type", factoryType))
	}
}

type factory interface {
	CreateLink(caption, url string) item
	CreateTray(caption string) tray
	CreatePage(title, author string) page
}

type item interface {
	makeHTML() string
}

type baseItem struct {
	caption string
}

type link struct {
	*baseItem
	url string
}

type tray interface {
	item
	Add(item item)
}

type baseTray struct {
	*baseItem
	tray []item
}

func (t *baseTray) Add(item item) {
	t.tray = append(t.tray, item)
}

type page interface {
	item
	Add(item item)
	Output(o item)
}

type basePage struct {
	title, author string
	content []item
}

func (p *basePage) Add(item item) {
	p.content = append(p.content, item)
}

func (p *basePage) Output(o item) {
	filename := fmt.Sprintf("./%s.html", p.title)
	file, _ := os.Create(filename)
	defer file.Close()
	b := []byte(o.makeHTML())
	file.Write(b)
	fmt.Printf("file '%s' was created.\n", filename)
}
