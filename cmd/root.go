package main

import (
	"github.com/hzliangbin/harbor-cli/pkg/types"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func newRootCmd(out io.Writer, args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "harbor-cli",
		Short: "A tool-cli for managing image to harbor",
		Long: "Use harbor-cli to list/delete/save/load images to harbor",
	}

	cmd.PersistentFlags().StringVarP(&types.CfgFile, "config","c","",
		"config file path，default to current dir to find registry-manager.yml")
	flags := cmd.PersistentFlags()
	flags.ParseErrorsWhitelist.UnknownFlags = true
	err := flags.Parse(args)
	if err != nil {
		os.Exit(1)
	}

	cmd.AddCommand(
		newDeleteCmd(out),
		newTopCmd(out))
	return cmd, nil
}