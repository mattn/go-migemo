package main

import (
  "fmt";
  "regexp";
  "strings";
  "migemo";
)

func main() {
	var pattern = "goGengo";
	var match = "go言語";
	m := migemo.Open("../dict/utf-8.d/migemo-dict");
	s := migemo.Query(m, pattern);
	if (regexp.MustCompile(s).Match(strings.Bytes(match))) {
	    fmt.Printf("%s は %s にマッチします！\n", pattern, match);
	} else {
	    fmt.Printf("%s は %s にマッチしません！\n", pattern, match);
	}
}
