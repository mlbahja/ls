package models

type Output struct {
	PathIsDir bool
	TotalSize int64
	Path      string
	Content   []string
}
