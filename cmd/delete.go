package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"io"

	. "github.com/hzliangbin/harbor-cli/pkg/types"
)

func newDeleteCmd(out io.Writer) *cobra.Command{
	cmd := &cobra.Command{
		Use:  "delete",
		Aliases: []string{"del"},
		Short:  "delete images",
		Long:  "delete images in harbor based some policy",
		RunE: func(cmd *cobra.Command, args []string) error{
			repositories := Manager.Repositories()
			repositoriesCount := len(repositories)
			glog.Infof("fetched repos totol: %d", repositoriesCount)
			deletedNum := 0
			for i, repo := range repositories {
				glog.Infof("handling %d/%d repo: %s", i+1, repositoriesCount, repo)
				tags := Manager.Tags(repo)
				if len(tags) <= Manager.DeleteController.MixCount {
					continue
				}
				tagObjs := Manager.TagObjs(repo)
				count := 0

				for _, tag := range tagObjs {
					if Manager.DeleteController.NeedDeleteTag(tag,&count) {
						glog.Infof("delete image: %s:%s", repo, tag.Name)
						if !Manager.DeleteController.DryRun {
							Manager.DeleteManifest(repo, tag.Name)
							deletedNum++
						}
					}
				}
			}
			fmt.Fprintf(out,"%d images deleted",deletedNum)
			return nil
		},
	}
	return cmd
}
