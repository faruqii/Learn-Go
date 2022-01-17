package main

import "fmt"

type (
	TreeNode struct {
	name string
	left *TreeNode
	right *TreeNode
	}
	FileSystem struct {
		root *FileSystem
		dirs map[string]*FileSystem
	}	
)

func (tn *TreeNode) add_subdir(name string) {
	if tn.left == nil {
		tn.left = &TreeNode{name, nil, nil}
	} else if tn.right == nil {
		tn.right = &TreeNode{name, nil, nil}
	} else {
		panic("Too many subdirectories")
	}
}

func (tn *TreeNode) print_name() {
	fmt.Print(tn.name)
}

func (fs *FileSystem) create_dir(name string) {
	if fs.root == nil {
		fs.root = &TreeNode(name)
		fs.dirs.append(fs.root)
		return
	}

func (fs *FileSystem) create_child(parent_name string, new_directory *TreeNode) {
	par = nil
	for _, dir := range fs.dirs {
		if dir.name == parent_name {
			par = dir
		}
	}
	if par == nil {
		panic("Parent directory not found")
	}
	par.add_subdir(new_directory)
	fs.dirs.append(new_directory)
}