package main

import (
	"fmt"
	"os"
	"strings"

	execution "my_ls/Execution"
	models "my_ls/Models"
	Parser "my_ls/Parser"
)

func printResult(Output []models.Output, flags models.Flags) {
	tmp := []models.Output{}
	line := ""
	for _, o := range Output {
		if o.PathIsDir {
			tmp = append(tmp, o)
		} else {
			if flags.LformatFlag {
				line += o.Content[0] + "\n"
			} else {
				line += o.Content[0] + "  "
			}
		}
	}
	if line != "" {
		line = strings.TrimRight(line, " ")
		line = strings.TrimRight(line, "\n")
		fmt.Println(line)
	}
	for i, o := range tmp {
		if o.Path == "." && flags.RecuFlag {
			fmt.Println(o.Path + ":")
		} else if len(Output) > 1 {
			if line != "" {
				fmt.Println()
			}
			fmt.Println(o.Path + ":")
		}
		if flags.LformatFlag {
			fmt.Println("total", o.TotalSize)
			for _, line := range o.Content {
				fmt.Println(line)
			}
			if i+1 < len(tmp) {
				fmt.Println()
			}
		} else {
			if len(o.Content) > 0 {
				fmt.Println(strings.Join(o.Content, "  "))
			}
			if i+1 < len(tmp) && flags.RecuFlag {
				fmt.Println()
			}
		}
	}
}

func callFlags(flags models.Flags, paths []string) {
	output := []models.Output{}
	if flags.RecuFlag {
		output = execution.RecuFlag(paths, flags)
	} else if flags.AllFlag {
		output = execution.AllFlag(paths, flags)
	}
	if flags.TimeFlag {
		output = execution.TimeFlag(paths, output)
	}
	if flags.RevOrderFlag {
		output = execution.RevOrderFlag(paths, output)
	}
	if flags.LformatFlag {
		output = execution.LongFormatFlag(paths, output)
	}
	if len(output) == 0 {
		output = execution.NoFlags(paths, output)
	}
	printResult(output, flags)
}

func main() {
	paths, flags := Parser.ParseArgs(os.Args[1:])
	if len(paths) == 0 {
		paths = append(paths, ".")
	}
	callFlags(flags, paths)
}
