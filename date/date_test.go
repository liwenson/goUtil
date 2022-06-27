package date

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

var (
	layout = [...]string{"y", "M", "w", "W", "D", "d", "F", "E", "a", "H", "k", "K", "h", "m", "s", "S", "z", "Z",
	}
)

func TestDate_Format(t *testing.T) {
	d := WithMillisecond(1219055318888)
	fmt.Println("|layout     |result |")
	fmt.Println("|:----    |:----        |")

	for _, lyt := range layout {
		lyts := make([]string, 0)
		for i := 0; i < 5; i++ {
			lyts = append(lyts, lyt)
			l := strings.Join(lyts, "")
			s, err := d.Format(l, false)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("|%s|%s|\n", l, s)
		}
	}
}
