package main

import (
	"container/heap"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"io"
	"sort"

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
			//repoHeap := make(Repos, TopNum)
			repoHeap := &Repos{}
			heap.Init(repoHeap)
			for i, repo := range repositories {
				tags := Manager.Tags(repo)
				glog.Infof("handling %d/%d repo: %s tags: %d", i+1, repositoriesCount, repo,len(tags))
				r := Repo{Name: repo, TagsNum: len(tags)}
				heap.Push(repoHeap, r)
				//pop出小的，最后剩下大的
				if len(*repoHeap)> TopNum {
					heap.Pop(repoHeap)
				}
			}
			sort.Sort(repoHeap)
			for i := len(*repoHeap) - 1; i >= 0; i-- {
				fmt.Fprint(out,(*repoHeap)[i].Name, (*repoHeap)[i].TagsNum)
			}
		},
	}
	f := cmd.Flags()
	f.IntVarP(&TopNum, "num","n",10,"top n to list the top n images with most tags")

	return cmd
}