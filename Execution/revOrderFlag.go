package execution

import (
	models "my_ls/Models"
	utils "my_ls/Utils"
)

func RevArrNoSort(output *[]string) {
	start := 0
	end := len(*output) - 1
	for start < end {
		(*output)[start], (*output)[end] = (*output)[end], (*output)[start]
		start++
		end--
	}
}

func RevNoSort(output *[]models.Output) {
	start := 0
	end := len(*output) - 1
	for _, o := range *output {
		RevArrNoSort(&o.Content)
	}
	for start < end {
		(*output)[start], (*output)[end] = (*output)[end], (*output)[start]
		start++
		end--
	}
}

func RevOrderFlag(paths []string, output []models.Output) []models.Output {
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
		tmp := []models.Output{}
		for i := 0; i < len(output); i++ {
			if output[i].Path == "." {
				tmp = append(tmp, output[i])
				if i+1 < len(output) {
					output = append(output[:i], output[i+1:]...)
				} else {
					output = output[:i]
				}
				break
			}
		}
		utils.RevSortStruct(&output)
		utils.RevSortStruct(&tmp)
		output = append(tmp, output...)
	} else {
		tmp := []models.Output{}
		for i := 0; i < len(output); i++ {
			if output[i].Path == "." {
				tmp = append(tmp, output[i])
				if i+1 < len(output) {
					output = append(output[:i], output[i+1:]...)
				} else {
					output = output[:i]
				}
				break
			}
		}
		RevNoSort(&output)
		output = append(tmp, output...)
	}
	return output
}
