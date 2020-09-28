package main

import (
	"github.com/spf13/cobra"
	"io"
)

func newLoadCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "load",
		Short: "load images",
		Long:  "load images from *.tar.gz and push to harbor",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
