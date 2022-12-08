package main

import (
	"flag"
	"fmt"
	"tools/internal/app"
)

func main() {
	workPath := flag.String("d", ".", "工作目录")
	searchText := flag.String("s", "", "要搜索的文本")
	isSearchFile := flag.Bool("f", false, "搜索文件")
	flag.Parse()

	var searchTools app.SearchTools

	if *isSearchFile {
		searchTools = &app.SearchFile{}
	} else {
		searchTools = &app.SearchText{}
	}

	result, err := searchTools.SearchHandler(*workPath, *searchText)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
