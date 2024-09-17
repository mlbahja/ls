package execution

import (
	models "my_ls/Models"
	utils "my_ls/Utils"
)

func NoFlags(paths []string, output []models.Output) []models.Output {
	if len(output) == 0 {
		dirs := paths
		for _, dir := range dirs {
			tmp := utils.GetContent(dir, false)
			if tmp.Path == "" {
				continue
			}
			output = append(output, tmp)
		}
		utils.SortStruct(&output)
	}
	utils.SortStruct(&output)
	return output
}
