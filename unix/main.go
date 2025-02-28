package main

import (
	"fmt"
	"unix/file"
	"unix/matcher"
	"unix/searcher"
)

func main() {
	// Simulated file system
	fs := file.NewFolder(
		"/home/user",
		[]*file.File{
			file.NewFile("test1", ".java", 400),
			file.NewFile("test2", ".java", 500),
			file.NewFolder("subdir", []*file.File{
				file.NewFile("test3", ".txt", 400),
				file.NewFile("test4", ".java", 200),
			}),
		},
	)

	// Create a filter to find files <= 5MB with .java extension
	sizeFilter := matcher.NewSizeFilter(400)
	extensionFiler := matcher.NewExtensionFilter([]string{".java"})


	searcher := searcher.NewSearcher([]matcher.Matcher{sizeFilter, extensionFiler})
	matchingFiles := searcher.SearchFiles(*fs)

	fmt.Println("Matching files:")
	for _, file := range matchingFiles {
		fmt.Println(file.GetFileName())
	}
}