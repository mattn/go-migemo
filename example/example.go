package main

import (
  "fmt";
  "regexp";
  "github.com/mattn/go-migemo/migemo";
)

func main() {
	var pattern = "goGengo";
	var match = "go言語";
	m := migemo.Open("./dict/utf-8.d/migemo-dict");
	s := migemo.Query(m, pattern);
	println(s)
	if (regexp.MustCompile(s).Match([]byte(match))) {
	    fmt.Printf("%s は %s にマッチします！\n", pattern, match);
	} else {
	    fmt.Printf("%s は %s にマッチしません！\n", pattern, match);
	}
}
