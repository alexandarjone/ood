package main

import (
	"fmt"
	"unix/file"
	"unix/filter"
	"unix/searcher"
)

func main() {
    file1 := file.NewFile("file1", ".txt", 50)
    file2 := file.NewFile("file2", ".java", 30)
    childFolder := file.NewFolder("child-folder", []file.File{file1, file2})
    file3 := file.NewFile("file3", ".txt", 20)
    file4 := file.NewFile("file4", ".python", 30)
    rootFolder := file.NewFolder("root-folder", []file.File{childFolder, file3, file4})

    sizeFilter := filter.NewSizeFilter(40)
    extensionFilter := filter.NewExtensionFiler([]string{".java", ".txt"})
    aggregateFilter := filter.NewAggregateFilter([]filter.Filter{sizeFilter, extensionFilter})

    filteredFiles := searcher.GetFilteredFiles(rootFolder, aggregateFilter)
    for _, file := range filteredFiles {
        fmt.Println(file.GetName(), file.GetExtension())
    }
}