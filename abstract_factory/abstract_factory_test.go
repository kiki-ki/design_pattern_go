package abstractFactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	expecteds := map[string]string{
		"list":  "<html>\n<head><title>MyPage</title></head>\n<body>\n<h1>MyPage</h1><ul><li>\nContents\n<ul>\n<li><a href=\"https://github.com/kiki-ki\">Github</a></li>\n</ul>\n</li>\n</ul><hr><adress>kiki-ki</adress></body>\n</html>\n",
		"table": "<html>\n<head><title>MyPage</title></head>\n<body>\n<h1>MyPage</h1><table width=\"80%\" border=\"3\">\n<tr><td>\n<table width=\"100%\" border=\"1\"><tr>\n<td bgcolor=\"#cccccc\" algin=\"center\" colsapn=\"1\"><b>Contents</b></td>\n</tr>\n<tr>\n<td><a href=https://github.com/kiki-ki>Github</a></td>\n</tr></table>\n</td>\n</tr></table><hr><adress>kiki-ki</adress></body>\n</html>\n",
	}
	t.Run("list", func(t *testing.T) {
		testAbstractFactory(t, New("list"), expecteds["list"])
	})
	t.Run("table", func(t *testing.T) {
		testAbstractFactory(t, New("table"), expecteds["table"])
	})
}

func testAbstractFactory(t *testing.T, f factory, expected string) {
	link := f.CreateLink("Github", "https://github.com/kiki-ki")
	tray := f.CreateTray("Contents")
	tray.Add(link)
	page := f.CreatePage("MyPage", "kiki-ki")
	page.Add(tray)
	result := page.makeHTML()

	if result != expected {
		t.Errorf("Expected result to %s, but %s\n", expected, result)
	}
}
