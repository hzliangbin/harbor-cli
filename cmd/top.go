package main

import (
	"container/heap"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"io"

	. "github.com/hzliangbin/harbor-cli/pkg/types"
)

func newTopCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "top",
		Short: "list the images with most tags",
		Long: "list the images with most tags",
		Run: func(cmd *cobra.Command, args []string) {
			repositories := Manager.Repositories()
			repositoriesCount := len(repositories)
			repoHeap := make(Repos, TopNum)
			heap.Init(&repoHeap)
			for i, repo := range repositories {
				glog.Infof("handling %d/%d repo: %s", i+1, repositoriesCount, repo)
				tags := Manager.Tags(repo)
				r := Repo{Name: repo, TagsNum: len(tags)}
				heap.Push(&repoHeap, r)
				if len(repoHeap) > TopNum {
					heap.Pop(&repoHeap)
				}
			}
			for _, repo := range repoHeap {
				fmt.Println(out, repo)
			}
		},
	}
	f := cmd.Flags()
	f.IntVarP(&TopNum, "num","n",10,"top n to list the top n images with most tags")

	return cmd
}