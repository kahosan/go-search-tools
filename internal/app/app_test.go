package app

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSearchText(t *testing.T) {
	workPath := "./"
	text := "app"

	var searchText = &SearchText{}

	result, err := searchText.SearchHandler(workPath, text)
	if err != nil {
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	buf.WriteString(result[0])

	// 验证输出
	if buf.String() != "文件名 app_test.go" {
		t.Errorf("unexpected output: %s", buf.String())
	}
}

func TestSearchFile(t *testing.T) {
	workPath := "./"
	text := "app"

	var searchFile = &SearchFile{}

	result, err := searchFile.SearchHandler(workPath, text)
	if err != nil {
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	buf.WriteString(result[1])

	// 验证输出
	if buf.String() != "文件名 app_test.go" {
		t.Errorf("unexpected output: %s", buf.String())
	}
}
