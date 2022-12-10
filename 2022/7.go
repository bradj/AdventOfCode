package twenttwo

import (
	"fmt"
	"strings"

	"github.com/bradj/AdventOfCode/util"
)

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	files       []file
	directories map[string]*directory
	parent      *directory
	size        int
}

var rootDir directory

func createDirStructure(items []string) {
	rootDir = directory{
		name:        "/",
		files:       []file{},
		directories: map[string]*directory{},
		size:        0,
	}

	var currentDir *directory

	for _, line := range items {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			parts := strings.Split(line, " ")
			newDirPath := parts[2]

			if newDirPath == "/" {
				currentDir = &rootDir
				continue
			}

			if newDirPath == ".." {
				currentDir = currentDir.parent
				continue
			}

			currentDir = currentDir.directories[newDirPath]
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "dir") {
			parts := strings.Split(line, " ")
			name := parts[1]

			newDir := directory{
				name:        name,
				parent:      currentDir,
				directories: map[string]*directory{},
			}

			currentDir.directories[name] = &newDir
		} else { // files
			parts := strings.Split(line, " ")
			name := parts[1]
			size := util.Intme(parts[0])

			f := file{
				size: size,
				name: name,
			}

			currentDir.files = append(currentDir.files, f)
			currentDir.size += f.size
		}
	}
}

func totalRecurse(current *directory) (int, []int) {
	total := 0
	sizes := []int{}

	for _, dir := range current.directories {
		childDirTotal, dirs := totalRecurse(dir)
		sizes = append(sizes, dirs...)
		sizes = append(sizes, childDirTotal)

		total += childDirTotal
	}

	return total + current.size, sizes
}

func D7p1(items []string) {
	createDirStructure(items)

	total := 0
	_, dirs := totalRecurse(&rootDir)

	for _, val := range dirs {
		if val < 100000 {
			total += val
		}
	}

	fmt.Println(total)
}

func D7p2(items []string) {
	createDirStructure(items)

	diskSpace := 70000000
	needed := 30000000
	used, dirs := totalRecurse(&rootDir)
	unused := diskSpace - used
	needed -= unused
	delete := 10000000000

	for _, val := range dirs {
		if val > needed {
			if val < delete {
				delete = val
			}
		}
	}

	fmt.Println(delete)
}
