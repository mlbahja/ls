package execution

import (
	models "my_ls/Models"
	utils "my_ls/Utils"
)

func AllFlag(paths []string, flags models.Flags) []models.Output {
	output := []models.Output{}
	dirs := paths
	for _, dir := range dirs {
		tmp := utils.GetContent(dir, true)
		if tmp.Path == "" {
			continue
		}
		output = append(output, tmp)
	}
	utils.SortStruct(&output)
	return output
}
