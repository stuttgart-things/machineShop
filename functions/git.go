/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package functions

import (
	"fmt"

	http "github.com/go-git/go-git/v5/plumbing/transport/http"

	memfs "github.com/go-git/go-billy/v5/memfs"
	memory "github.com/go-git/go-git/v5/storage/memory"

	"github.com/go-git/go-git/v5"
)

func GetFileListFromGitRepository(repository, directory string, auth *http.BasicAuth) (fileList []string) {

	// Init memory storage and fs
	storer := memory.NewStorage()
	fs := memfs.New()

	// Clone repo into memfs
	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:  repository,
		Auth: auth,
	})

	if err != nil {
		fmt.Println("Could not git clone repository")
	}
	fmt.Println("Repository cloned")

	// Get git default worktree
	w, err := r.Worktree()
	if err != nil {
		fmt.Println("Could not get git worktree: %w", err)
	}

	fmt.Println(w)
	files, _ := fs.ReadDir(directory)

	for _, file := range files {
		fileList = append(fileList, file.Name())
		fmt.Println(file.Name())
	}

	// file, _ := fs.Open(filename)
	// fileContent, _ := ioutil.ReadAll(file)

	// return string(fileContent)
	return
}
