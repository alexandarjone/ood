package searcher

import (
	"unix/file"
	"unix/matcher"
)


type Searcher struct {
	matchers []matcher.Matcher
}

func NewSearcher(matchers []matcher.Matcher) *Searcher {
	return &Searcher{matchers}
}

func (s Searcher) SearchFiles(directory file.File) []file.File {
	validFiles := []file.File{}
	var dfs func(file file.File)
	dfs = func(file file.File) {
		if !file.IsDirectory() {
			isMatched := true
			for _, matcher := range s.matchers {
				if !matcher.IsMatched(file) {
					isMatched = false
					break
				}
			}
			if isMatched {
				validFiles = append(validFiles, file)
			}
			return
		}
		for _, child := range file.ListOfSubDirector() {
			dfs(*child)
		}
	}
	dfs(directory)
	return validFiles
}