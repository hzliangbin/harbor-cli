package main

import (
	"github.com/spf13/cobra"
	"io"
)

func newSaveCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "save",
		Short:  "save images",
		Long: "save images from harbor to tar.gz",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
