package main
import (
	"net/http"
	"fmt"
	"time"
	"strings"
)
func main() {
	var webhook = "https://discord.com/api/webhooks/824939217502732300/3t0_c9E_X2QUrDWOaKZ7CJHIeDcDYVdjhaYrLvlN06Wgur3Y9XpwcjBbxHg9teQaup92"
	var i = 0
	var payload =`{"content": "This has been going on for %d seconds"}`
	for {
		_, err := http.Post(webhook, "application/json", strings.NewReader(fmt.Sprintf(payload, i)))
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Duration(1)*time.Second)
		i += 1
	}
}
