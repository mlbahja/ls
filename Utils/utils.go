package utils

import (
	"fmt"
	"os"
	"sort"
	"strings"

	models "my_ls/Models"
)

func filterStr(str string) string {
	result := ""
	if str == "." || str == ".." {
		return str
	}
	for _, c := range str {
		if (c <= 'z' && c >= 'a') || (c <= 'Z' && c >= 'A') || (c <= '9' && c >= '0') {
			result += string(c)
		}
	}
	return result
}

func SortArr(arr *[]string) {
	sort.SliceStable(*arr, func(i, j int) bool {
		return strings.ToLower(filterStr((*arr)[i])) < strings.ToLower(filterStr((*arr)[j]))
	})
}

func SortStruct(arr *[]models.Output) {
	sort.SliceStable(*arr, func(i, j int) bool {
		return strings.ToLower(filterStr((*arr)[i].Path)) < strings.ToLower(filterStr((*arr)[j].Path))
	})
	for i := range *arr {
		SortArr(&(*arr)[i].Content)
	}
}

func RevSortArr(arr *[]string) {
	sort.SliceStable(*arr, func(i, j int) bool {
		return strings.ToLower(filterStr((*arr)[i])) > strings.ToLower(filterStr((*arr)[j]))
	})
}

func RevSortStruct(arr *[]models.Output) {
	sort.SliceStable(*arr, func(i, j int) bool {
		return strings.ToLower(filterStr((*arr)[i].Path)) > strings.ToLower(filterStr((*arr)[j].Path))
	})
	for i := range *arr {
		RevSortArr(&(*arr)[i].Content)
	}
}

func GetContent(dir string, allFlag bool) models.Output {
	output := models.Output{}
	openedDir, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ls: cannot access '%s'\n", dir)
		output.Path = ""
		return output
	}
	defer openedDir.Close()
	files, _ := openedDir.Readdir(-1)
	output.Path = dir
	tmp, _ := os.Stat(dir)
	if !tmp.IsDir() {
		output.Content = append(output.Content, dir)
		output.PathIsDir = false
		return output
	}
	for _, f := range files {
		if !allFlag && len(f.Name()) > 1 && f.Name()[0] == '.' {
			continue
		}
		output.Content = append(output.Content, f.Name())
	}
	if allFlag {
		output.Content = append(output.Content, ".")
		output.Content = append(output.Content, "..")
	}
	output.PathIsDir = true
	return output
}
