package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println
	fmt.Println("\xe6\x97\xa5") //日	\xNN specifies a byte
	fmt.Println(s.EqualFold("Japan", "JAPAN"))

	//Search (contains, prefix/suffix, index)
	fmt.Println(s.Contains("India", "abc"))    //false Is abc in India?
	fmt.Println(s.ContainsAny("India", "abc")) //true	Is a, b or c in India?)

	p(s.Count("Banana", "ana"))

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("testeste", "te"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	p("split:", s.Split("http://google.com//asdf.gyan//.chand", "//"))
	p("Replace:", s.Replace("foooouck", "o", "*", 2))
	p("Replace:", s.Replace("foooouck", "o", "*", -1))

	p("Upper:", s.ToUpper("Japan"))
	p("Lower:", s.ToLower("JAPaN"))
	p("Title:", s.Title("ja pan"))
	p("Trim space:", s.TrimSpace("      foo        ")) //Strip leading and trailing white space
	p("Trim:", s.Trim("foo", "fo"))                    //Strip leading and trailing f:s and o:s
	p("Trim Left:", s.TrimLeft("foo", "f"))            //Strip leading and trailing f:s and o:s
	p("Trim right:", s.TrimRight("foo", "o"))          //Strip leading and trailing f:s and o:s
}
