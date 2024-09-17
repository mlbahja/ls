package execution

import (
	"os"
	"strings"

	models "my_ls/Models"
	utils "my_ls/Utils"
)

func dirTravers(path string, dirs *[]string, allFlag bool) {
	openedPath, err := os.Open(path)
	if err != nil {
		return
	}
	*dirs = append(*dirs, path)
	files, _ := openedPath.Readdir(-1)
	for _, f := range files {
		if len(f.Name()) > 1 && f.Name()[0] == '.' && !allFlag {
			continue
		}
		if f.IsDir() {
			if strings.HasSuffix(path, "/") {
				path = strings.TrimRight(path, "/")
			}
			dirTravers(path+"/"+f.Name(), dirs, allFlag)
		}
	}
	openedPath.Close()
}

func RecuFlag(paths []string, flags models.Flags) []models.Output {
	output := []models.Output{}
	dirs := []string{}
	for _, path := range paths {
		dirTravers(path, &dirs, flags.AllFlag)
	}
	for _, dir := range dirs {
		tmp := utils.GetContent(dir, flags.AllFlag)
		if tmp.Path == "" {
			continue
		}
		output = append(output, tmp)
	}
	utils.SortStruct(&output)
	return output
}
