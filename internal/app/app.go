package app

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
	"tools/pkg/errors"
	"tools/pkg/util"
)

type SearchTools interface {
	SearchHandler(path, key string) ([]string, error)
}

type SearchFile struct {
	Result []string
}

type SearchText struct {
	Result []string
}

func NewSearchTools(f SearchTools) SearchTools {
	return f
}

func (s *SearchText) SearchHandler(path, key string) ([]string, error) {
	if key == "" {
		return nil, errors.NewError("searchText Error", fmt.Errorf("请输入要搜索的文本"))
	}

	fileNames, err := util.GetAllFileName(path)
	if err != nil {
		return nil, errors.NewError("searchText Error", fmt.Errorf("获取文件名失败: %v", err))
	}

	var wg sync.WaitGroup
	wg.Add(len(fileNames))

	var mu sync.Mutex

	for _, fileName := range fileNames {
		go func(fileName string) {
			defer wg.Done()

			fileContent, err := os.Open(fileName)
			if err != nil {
				fmt.Printf("读取文件 %s 失败: %v", fileName, err)
				return
			}
			defer fileContent.Close()

			scanner := bufio.NewScanner(fileContent)

			pattern, regErr := regexp.Compile(key)
			if regErr != nil {
				fmt.Printf("无效的文本: %v", regErr)
			}
			for scanner.Scan() {
				if pattern.MatchString(scanner.Text()) {
					mu.Lock()
					s.Result = append(s.Result, "文件名 "+fileName)
					mu.Unlock()
					break
				}
			}

		}(fileName)
	}

	wg.Wait()

	if s.Result == nil {
		return nil, errors.NewError("searchText Error", fmt.Errorf("未找到文本"))
	}

	return s.Result, nil
}

func (s *SearchFile) SearchHandler(path, key string) ([]string, error) {
	if key == "" {
		return nil, errors.NewError("searchFile Error", fmt.Errorf("请输入要搜索的文件名"))
	}

	fileNames, err := util.GetAllFileName(path)
	if err != nil {
		return nil, errors.NewError("searchText Error", fmt.Errorf("获取文件名失败: %v", err))
	}

	pattern, regErr := regexp.Compile(key)
	if regErr != nil {
		return nil, errors.NewError("searchFile Error", fmt.Errorf("无效的文件名: %v", regErr))
	}

	for _, fileName := range fileNames {
		if pattern.MatchString(fileName) {
			s.Result = append(s.Result, "文件名 "+fileName)
		}
	}

	if s.Result == nil {
		return nil, errors.NewError("searchFile Error", fmt.Errorf("未找到文件名"))
	}

	return s.Result, nil
}
