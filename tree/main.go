package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Ref: https://www.youtube.com/watch?v=XbKSssBftLM

var (
	dirCount, fileCount int
)

func main() {
	// define default arg
	args := []string{"."}
	// overwrite default arg
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		if err := tree(arg, "  "); err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
		fmt.Printf("%d directories, %d files\n\n", dirCount, fileCount)
		dirCount, fileCount = 0, 0
	}
}

// tree lists the contents of directories in a tree-like format.
// ident is used to customize indentation at the head of each line.
func tree(root, ident string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}

	if fi.IsDir() {
		dirCount++
	} else {
		fileCount++
	}

	fmt.Println(fi.Name())

	// check the specified root indicates directory or not.
	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", root, err)
	}

	// filter every dot file in early stage.
	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		add := "│  "
		if i == len(names)-1 {
			fmt.Printf(ident + "└──")
			add = "   "
		} else {
			fmt.Printf(ident + "├──")
		}

		if tree(filepath.Join(root, name), ident+add); err != nil {
			return nil
		}
	}

	return nil
}
