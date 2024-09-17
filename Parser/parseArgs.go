package parser

import (
	"fmt"
	"os"

	models "my_ls/Models"
)

func initFlags(flags *models.Flags) {
	flags.AllFlag = false
	flags.LformatFlag = false
	flags.RecuFlag = false
	flags.RevOrderFlag = false
	flags.TimeFlag = false
}

func ParseArgs(args []string) ([]string, models.Flags) {
	flags := models.Flags{}
	paths := []string{}
	initFlags(&flags)
	for _, arg := range args {
		if len(arg) > 1 && arg[0] == '-' {
			arg = arg[1:]
			for _, c := range arg {
				if c == 'a' {
					flags.AllFlag = true
				} else if c == 'l' {
					flags.LformatFlag = true
				} else if c == 'R' {
					flags.RecuFlag = true
				} else if c == 'r' {
					flags.RevOrderFlag = true
				} else if c == 't' {
					flags.TimeFlag = true
				} else {
					fmt.Fprintln(os.Stderr, "my_ls: unrecognized option")
					os.Exit(2)
				}
			}
		} else {
			paths = append(paths, arg)
		}
	}
	return paths, flags
}
