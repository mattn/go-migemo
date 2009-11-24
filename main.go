package main

import "migemo"

func main() {
	m := migemo.Open("../dict/utf-8.d/migemo-dict");
	s := migemo.Query(m, "watasi");
	println(s);
}
