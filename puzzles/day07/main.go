package main

import (
	"fmt"
	"github.com/kylehoehns/aoc-2022-go/utils/files"
	"github.com/kylehoehns/aoc-2022-go/utils/ints"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(name string) int {
	lines := files.ReadLines(name)

	root := buildFileDirectory(lines)

	sumOfDirectorySizes := 0
	for _, dir := range root.AllSubDirectories() {
		dirSize := dir.TotalSize()
		if dirSize <= 100_000 {
			sumOfDirectorySizes += dirSize
		}
	}
	return sumOfDirectorySizes
}

func part2(name string) int {

	lines := files.ReadLines(name)
	root := buildFileDirectory(lines)
	unusedSpace := 70_000_000 - root.TotalSize()
	amountNeededToFree := 30_000_000 - unusedSpace
	candidatesForDeletion := make([]*Dir, 0)
	for _, dir := range root.AllSubDirectories() {
		if dir.TotalSize() >= amountNeededToFree {
			candidatesForDeletion = append(candidatesForDeletion, dir)
		}
	}

	smallestDir := candidatesForDeletion[0]
	for _, dir := range candidatesForDeletion {
		if dir.TotalSize() < smallestDir.TotalSize() {
			smallestDir = dir
		}
	}

	return smallestDir.TotalSize()
}

func buildFileDirectory(lines []string) *Dir {
	root := NewDir()
	curr := root
	for _, line := range lines {
		if string(line[0]) == "$" {
			command := line[2:]
			if command == "ls" {
				continue
			} else {
				cdArg := line[5:]
				if cdArg == ".." {
					curr = curr.parent
				} else if cdArg == "/" {
					curr = root
				} else {
					curr = curr.children[cdArg]
				}
			}
		} else {
			if line[:3] == "dir" {
				directoryName := line[4:]
				if _, ok := curr.children[directoryName]; !ok {
					curr.children[directoryName] = NewDirWithParent(curr)
				}
			} else {
				parts := strings.Split(line, " ")
				fileSize := ints.FromString(parts[0])
				fileName := parts[1]
				if _, ok := curr.files[fileName]; !ok {
					curr.files[fileName] = fileSize
				}
			}
		}
	}
	return root
}

type Dir struct {
	parent   *Dir
	children map[string]*Dir
	files    map[string]int
}

func NewDir() *Dir {
	return &Dir{
		children: map[string]*Dir{},
		files:    map[string]int{},
	}
}

func NewDirWithParent(parent *Dir) *Dir {
	return &Dir{
		parent:   parent,
		children: map[string]*Dir{},
		files:    map[string]int{},
	}
}

func (d *Dir) TotalSize() int {
	childDirSizes := 0
	for _, childDir := range d.children {
		childDirSizes += childDir.TotalSize()
	}
	fileSizes := 0
	for _, fileSize := range d.files {
		fileSizes += fileSize
	}
	return childDirSizes + fileSizes
}

func (d *Dir) AllSubDirectories() []*Dir {
	dirs := make([]*Dir, 0)
	dirs = append(dirs, d)
	for _, child := range d.children {
		dirs = append(dirs, child.AllSubDirectories()...)
	}
	return dirs
}
