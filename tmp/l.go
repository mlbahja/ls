package tmp

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"syscall"

	models "my_ls/Models"
)

func LFlag(paths []string, output []models.Output) []models.Output {
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
	for _, o := range output {
	}
	files, _ := d.Readdir(-1)
	totalSize := fmt.Sprintln("total", strconv.Itoa(GetTotalSize(files)))
	for _, file := range files {
		if file.Name()[0] == '.' {
			continue
		}
		tmp := myStructs.Result{}
		lenTmp := myStructs.ResultLen{}
		symlink, info := utils.GetSymLink(file, path)
		if len(symlink) != 0 {
			tmp.Name = file.Name() + " -> " + symlink
			tmp.Perms = info.Mode().String()
			tmp.Perms = "l" + tmp.Perms[1:]
			lenTmp.Perms = len(tmp.Perms)
		} else {
			tmp.Name = file.Name()
			tmp.Perms = file.Mode().Perm().String()
			lenTmp.Perms = len(tmp.Perms)
		}
		tmp.Num = "1"
		lenTmp.Num = len(tmp.Num)
		if file.IsDir() {
			tmp.Num = strconv.Itoa((utils.CountDir(file) + 2))
			lenTmp.Num = len(tmp.Num)
			tmp.Perms = "d" + tmp.Perms[1:]
		}
		Group, err := user.LookupGroupId(strconv.Itoa(int(file.Sys().(*syscall.Stat_t).Gid)))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tmp.GroupOwn = Group.Name
		lenTmp.UserOwn = len(tmp.UserOwn)
		user, err := user.LookupId(fmt.Sprintf("%d", file.Sys().(*syscall.Stat_t).Uid))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tmp.UserOwn = user.Username
		lenTmp.GroupOwn = len(tmp.GroupOwn)
		tmp.Date = file.ModTime().Format("Jan 02 15:04")
		lenTmp.Date = len(tmp.Date)
		tmp.Size = strconv.FormatInt(file.Size(), 10)
		lenTmp.Size = len(tmp.Size)
		results = append(results, tmp)
		resultsLen = append(resultsLen, lenTmp)
	}
	return PrintResults(results, resultsLen, totalSize)
}

func addSpaces(str string, num int) string {
	for i := 0; i < num; i++ {
		str = " " + str
	}
	return str
}

func PrintResults(results []myStructs.Result, resultsLen []myStructs.ResultLen, totalSize string) []string {
	permsMax := findMaxLen(resultsLen, "Perms")
	numMax := findMaxLen(resultsLen, "Num")
	userOwnMax := findMaxLen(resultsLen, "UserOwn")
	GroupOwnMax := findMaxLen(resultsLen, "GroupOwn")
	sizeMax := findMaxLen(resultsLen, "Size")
	dateMax := findMaxLen(resultsLen, "Date")
	output := []string{}
	output = append(output, totalSize)
	for i, result := range results {
		if resultsLen[i].Perms < permsMax {
			result.Perms = addSpaces(result.Perms, permsMax-resultsLen[i].Perms)
		}
		if resultsLen[i].Num < numMax {
			result.Num = addSpaces(result.Num, numMax-resultsLen[i].Num)
		}
		if resultsLen[i].UserOwn < userOwnMax {
			result.UserOwn = addSpaces(result.UserOwn, userOwnMax-resultsLen[i].UserOwn)
		}
		if resultsLen[i].GroupOwn < GroupOwnMax {
			result.GroupOwn = addSpaces(result.GroupOwn, GroupOwnMax-resultsLen[i].GroupOwn)
		}
		if resultsLen[i].Size < sizeMax {
			result.Size = addSpaces(result.Size, sizeMax-resultsLen[i].Size)
		}
		if resultsLen[i].Date < dateMax {
			result.Date = addSpaces(result.Date, dateMax-resultsLen[i].Date)
		}
		output = append(output, fmt.Sprintln(result.Perms, result.Num, result.UserOwn, result.GroupOwn, result.Size, result.Date, result.Name))
	}
	return output
}

func CountDir(dir fs.FileInfo) int {
	path := "." + "/" + dir.Name()
	content, _ := os.ReadDir(path)
	count := 0
	for _, c := range content {
		if c.IsDir() {
			count++
		}
	}
	return count
}

func GetTotalSize(files []fs.FileInfo) int {
	totalSize := 0
	for _, f := range files {
		if f.Name()[0] == '.' {
			continue
		}
		tmp := 4
		if f.Size() > 4096 {
			tmp = int(f.Size()) / 1024
			if int(f.Size())%1024 != 0 {
				tmp += 4
			}
		}
		totalSize += tmp
	}
	return totalSize
}

func GetSymLink(file fs.FileInfo, path string) (string, fs.FileInfo) {
	fullPath := path + "/" + file.Name()
	info, err := os.Lstat(fullPath)
	if err != nil {
		os.Exit(1)
	}
	symLink := ""
	if info.Mode()&os.ModeSymlink != 0 {
		var err error
		symLink, err = os.Readlink(fullPath)
		if err != nil {
			return "", nil
		}
	}
	return symLink, info
}

func findMaxLen(resultsLen []myStructs.ResultLen, field string) int {
	max := 0
	for _, res := range resultsLen {
		switch field {
		case "Perms":
			if res.Perms > max {
				max = res.Perms
			}
		case "Num":
			if res.Num > max {
				max = res.Num
			}
		case "UserOwn":
			if res.UserOwn > max {
				max = res.UserOwn
			}
		case "GroupOwn":
			if res.GroupOwn > max {
				max = res.GroupOwn
			}
		case "Size":
			if res.Size > max {
				max = res.Size
			}
		case "Date":
			if res.Date > max {
				max = res.Date
			}
		}
	}
	return max
}
