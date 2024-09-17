package execution

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"

	models "my_ls/Models"
	utils "my_ls/Utils"
)

func CountDir(path string) int {
	content, _ := os.ReadDir(path)
	count := 0
	for _, c := range content {
		if c.IsDir() {
			count++
		}
	}
	return count
}

// func GetTotalSize(files []fs.FileInfo) int {
// 	totalSize := 0
// 	for _, f := range files {
// 		if f.Name()[0] == '.' {
// 			continue
// 		}
// 		tmp := 4
// 		if f.Size() > 4096 {
// 			tmp = int(f.Size()) / 1024
// 			if int(f.Size())%1024 != 0 {
// 				tmp += 4
// 			}
// 		}
// 		totalSize += tmp
// 	}
// 	return totalSize
// }

func LongFormatFlag(paths []string, output []models.Output) []models.Output {
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
	var totalSize int64
	for i, o := range output {
		totalSize = 0
		for i := 0; i < len(o.Content); i++ {
			fullPath := ""
			if !o.PathIsDir {
				fullPath = o.Path
			} else {
				fullPath = o.Path + "/" + o.Content[i]
			}
			symLink := ""
			info, _ := os.Stat(fullPath)
			link, _ := os.Lstat(fullPath)
			if link.Mode()&os.ModeSymlink != 0 {
				symLink, _ = os.Readlink(fullPath)
				o.Content[i] = strings.ToLower(link.Mode().String() + "  ")
			} else {
				o.Content[i] = info.Mode().Perm().String() + "  "
				if info.IsDir() {
					o.Content[i] = "d" + o.Content[i][1:]
				}
			}
			if symLink != "" {
				count := 1
				if link.IsDir() {
					count += 1
				}
				o.Content[i] += strconv.Itoa(count+(CountDir(fullPath))) + "  "
				Group, err := user.LookupGroupId(strconv.Itoa(int(link.Sys().(*syscall.Stat_t).Gid)))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				o.Content[i] += Group.Name + "  "
				User, err := user.LookupId(strconv.Itoa(int(link.Sys().(*syscall.Stat_t).Uid)))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				o.Content[i] += User.Username + "  "
				tmp := 4
				if link.Size() > 4096 {
					tmp = int(link.Size()) / 1024
					if int(link.Size())%1024 != 0 {
						tmp += 4
					}
				}
				totalSize += int64(tmp)
				o.Content[i] += strconv.FormatInt(link.Size(), 10) + "  "
				o.Content[i] += link.ModTime().Format("Jan 02 15:04") + "  "
				o.Content[i] += link.Name() + " -> " + symLink
			} else {
				count := 1
				if info.IsDir() {
					count += 1
				}
				o.Content[i] += strconv.Itoa(count+(CountDir(fullPath))) + "  "
				Group, err := user.LookupGroupId(strconv.Itoa(int(info.Sys().(*syscall.Stat_t).Gid)))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				o.Content[i] += Group.Name + "  "
				User, err := user.LookupId(strconv.Itoa(int(info.Sys().(*syscall.Stat_t).Uid)))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				o.Content[i] += User.Username + "  "
				tmp := 4
				if info.Size() > 4096 {
					tmp = int(info.Size()) / 1024
					if int(info.Size())%1024 != 0 {
						tmp += 4
					}
				}
				totalSize += int64(tmp)
				o.Content[i] += strconv.FormatInt(info.Size(), 10) + "  "
				o.Content[i] += info.ModTime().Format("Jan 02 15:04") + "  "
				o.Content[i] += info.Name()
			}
		}
		output[i].TotalSize = totalSize
	}
	return output
}
