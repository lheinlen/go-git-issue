package main

import (
	"fmt"

	"gopkg.in/src-d/go-billy.v4/memfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func main() {
	// Clone master where .gitignore contains the /b directory
	repo, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL: "https://github.com/lheinlen/go-git-issue",
	})
	poe(err)

	worktree, err := repo.Worktree()
	poe(err)

	// Switch to a commit where .gitignore does not contain /b and and a file has been committed in that directory
	poe(worktree.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash("9842e001835f3e1022d717f6f0f49e205dc7b05d"),
	}))

	status, err := worktree.Status()
	poe(err)

	for f, fs := range status {
		fmt.Printf("%s %c\n", f, fs.Worktree)
	}
}

func poe(err error) {
	if err != nil {
		panic(err)
	}
}
