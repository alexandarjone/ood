package searcher

import (
	"unix/file"
	"unix/filter"
)


type Searcher interface {
	GetFilteredFiles(file.File) []file.File
}

type UnixSearcher struct {
	filter filter.Filter
}

func (u UnixSearcher) GetFilteredFiles(f file.File) []file.File {
	filteredFiles := []file.File{}
	var dfs func(currentFile file.File)
	dfs = func(currentFile file.File) {
		if !currentFile.IsDirectory() && u.filter.IsMatched(currentFile){
			filteredFiles = append(filteredFiles, currentFile)
			return
		}
		for _, child := range currentFile.ListOfSubDirectory() {
			dfs(child)
		}
	}
	dfs(f)
	return filteredFiles
}

func GetFilteredFiles(rootFolder file.File, searchFilter filter.Filter) []file.File {
	filteredFiles := []file.File{}
	var dfs func(currentFile file.File)
	dfs = func(currentFile file.File) {
		if !currentFile.IsDirectory() && searchFilter.IsMatched(currentFile){
			filteredFiles = append(filteredFiles, currentFile)
			return
		}
		for _, child := range currentFile.ListOfSubDirectory() {
			dfs(child)
		}
	}
	dfs(rootFolder)
	return filteredFiles
}