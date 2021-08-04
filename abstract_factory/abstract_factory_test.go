package absf

import "testing"

func TestListFactory(t *testing.T) {
	f := New("list")
	page := setPage(f)
	result := page.makeHTML()
	expect := "<html>\n<head><title>MyPage</title></head>\n<body>\n<h1>MyPage</h1><ul><li>\nContents\n<ul>\n<li><a href=\"https://github.com/kiki-ki\">Github</a></li>\n</ul>\n</li>\n</ul><hr><adress>kiki-ki</adress></body>\n</html>\n"
	if result != expect {
		t.Errorf("Expect result to %s, but %s\n", expect, result)
	}
}

func TestTableFactory(t *testing.T) {
	f := New("table")
	page := setPage(f)
	page.Output(page)
	result := page.makeHTML()
	expect := "<html>\n<head><title>MyPage</title></head>\n<body>\n<h1>MyPage</h1><table width=\"80%\" border=\"3\">\n<tr><td>\n<table width=\"100%\" border=\"1\"><tr>\n<td bgcolor=\"#cccccc\" algin=\"center\" colsapn=\"1\"><b>Contents</b></td>\n</tr>\n<tr>\n<td><a href=https://github.com/kiki-ki>Github</a></td>\n</tr></table>\n</td>\n</tr></table><hr><adress>kiki-ki</adress></body>\n</html>\n"
	if result != expect {
		t.Errorf("Expect result to %s, but %s\n", expect, result)
	}
}

func setPage(f factory) page {
	link := f.CreateLink("Github", "https://github.com/kiki-ki")
	tray := f.CreateTray("Contents")
	tray.Add(link)
	page := f.CreatePage("MyPage", "kiki-ki")
	page.Add(tray)
	return page
}
