package json_url

import (
	"fmt"
	"regexp"
)

func Regular() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)                    // true
	r, _ := regexp.Compile("p([a-z]+)ch") // так быстрее, скомпилитровать регулярку в переменную и ее использвать
	fmt.Println(r.MatchString("peach"))
}
