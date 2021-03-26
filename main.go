package main
import (
	"fmt"
	"github.com/anaskhan96/soup"
)

var support = "https://rms-support-letter.github.io/"
var open = "https://rms-open-letter.github.io/"
func get_supporters(url string) int {
	resp, err := soup.Get(url)
	if err != nil {
	}

	body := soup.HTMLParse(resp)
	var signs soup.Root
	if url == support {
		signs = body.Find("main","class","page-content").Find("div", "class", "wrapper").Find("div", "class","home").Find("ol")
	} else {
		signs = body.Find("main","class","page-content").Find("div", "class", "wrapper").Find("div", "class","home").FindAll("ol")[1]
	}
	// Do you like divisions? I do.
	return len(signs.Children())/2 // B)
	// Why? Who knows. Black magic. If it works, don't break it.


	// Note: Fuck HTML.
}



func main() {
		fmt.Printf("Supporters: %d, Opposers: %d\n", get_supporters(support), get_supporters(open))
}
