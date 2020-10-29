package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("contains:  ", s.Contains("test", "es"))
	p("count:     ", s.Count("test", "t"))
	p("hasprefix: ", s.HasPrefix("test", "te"))
	p("hassuffix: ", s.HasSuffix("test", "st"))
	p("index:     ", s.Index("test", "e"))
	p("join:      ", s.Join([]string{"best", "test"}, "-"))
	p("repeat:    ", s.Repeat("test-", 4))
	p("replace:   ", s.Replace("test-test", "te", "twi", -1))
	p("replace:   ", s.Replace("test-test", "t", "s", 2))
	p("split:     ", s.Split("t-e-s-t", "-"))
	p("tolower:   ", s.ToLower("TEST"))
	p("toupper:   ", s.ToUpper("test"))

	p()

	p("len:       ", len("test"))
	p("char:      ", "test"[1])
	// Indexing and len work at byte level. Go uses UTF-8 encoded strings.
	// For multi byte characters, we might need to use encoding aware operations
}
