package main

import (
	"fmt"
	"os"
)

/*
aert
boiy
*/
func afficheDirFile(arr []string) {
	for i, v := range arr {
		if i < len(arr)-1 {

			fmt.Print(v + "  ")
		} else {
			fmt.Println(v)
		}
	}
}
func sortstring(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
func revSorte(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
func main() {
	dir := "."

	args := os.Args[1:]
	if len(os.Args) < 1 || len(os.Args) > 2 {
		fmt.Println("number of arument not vlid")
		return
	}
	globalFlage := ""
	flags := []string{"-l", "-R", "-r", "-a", "-t"}
	for _, str := range args {
		for _, flag := range flags {
			if str == flag {
				globalFlage += str
			}
		}
	}
	fmt.Println(globalFlage, "if valid global")
	file, _ := os.Open(dir)
	f, _ := file.Readdirnames(0)
	if globalFlage == "-r" {
		f = revSorte(f)
		afficheDirFile(f)
	} else {
		f = sortstring(f)
		afficheDirFile(f)
	}

	// if args == []string{"-r"} {
	// 	f = revSorte(f)
	// }
	// for i, v := range f {
	// 	if i < len(f)-1 {

	// 		fmt.Print(v + "  ")
	// 	} else {
	// 		fmt.Println(v)
	// 	}
	// }

}
