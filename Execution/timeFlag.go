package execution

import (
	"fmt"
	"os"
	"sort"
	"time"

	models "my_ls/Models"
	utils "my_ls/Utils"
)

func getPathModTime(output []models.Output) []SortPathWithTime {
	times := []SortPathWithTime{}
	tmp := SortPathWithTime{}
	for _, path := range output {
		stats, err := os.Stat(path.Path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		tmp.output = path
		tmp.Time = stats.ModTime()
		times = append(times, tmp)
	}
	return times
}

func getFileModTime(path string, output []string) []SortFileWithTime {
	times := []SortFileWithTime{}
	tmp := SortFileWithTime{}
	for _, file := range output {
		stats, err := os.Stat(path + "/" + file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		tmp.output = file
		tmp.Time = stats.ModTime()
		times = append(times, tmp)
	}
	return times
}

// func sortBasedOnTimes[T any](items []T, times []time.Time, swap func(int, int)) {
// 	if len(items) != len(times) {
// 		panic("items and times slices must be of the same length")
// 	}

// 	sort.SliceStable(times, func(i, j int) bool {
// 		return times[i].Before(times[j])
// 	})

// 	sort.SliceStable(items, func(i, j int) bool {
// 		return times[i].Before(times[j])
// 	})
// }

// func SortContent(content *[]string, times []time.Time) {
// 	sortBasedOnTimes(*content, times, func(i, j int) {
// 		(*content)[i], (*content)[j] = (*content)[j], (*content)[i]
// 	})
// }

func sortPathByTime(slice []SortPathWithTime) {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i].Time.After(slice[j].Time)
	})
}

func sortFileByTime(slice []SortFileWithTime) {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i].Time.After(slice[j].Time)
	})
}

// func SortContent(arr *[]string, times []time.Time) {
// 	sort.SliceStable(*arr, func(i, j int) bool {
// 		return times[i].Unix() < times[j].Unix()
// 	})
// }

type SortPathWithTime struct {
	Time   time.Time
	output models.Output
}

type SortFileWithTime struct {
	Time   time.Time
	output string
}

func TimeFlag(paths []string, output []models.Output) []models.Output {
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
	times := getPathModTime(output)
	sortPathByTime(times)
	for i := 0; i < len(times); i++ {
		output[i] = times[i].output
		times := getFileModTime(output[i].Path, output[i].Content)
		sortFileByTime(times)
		for j := 0; j < len(times); j++ {
			output[i].Content[j] = times[j].output
		}
	}
	tmp := []models.Output{}
	index := 0
	for ; index < len(output); index++ {
		if output[index].Path == "." {
			tmp = append(tmp, output[index])
			break
		}
	}
	tmp = append(tmp, output[:index]...)
	if index+1 < len(output) {
		output = append(tmp, output[index+1:]...)
	}
	return output
}
