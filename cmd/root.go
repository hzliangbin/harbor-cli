package main

import (
	"github.com/hzliangbin/harbor-cli/pkg/types"
	"github.com/spf13/cobra"
	"io"
)

func newRootCmd(out io.Writer, args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "harbor-cli",
		Short: "A tool-cli for managing image to harbor",
		Long: "Use harbor-cli to list/delete/save/load images to harbor",
	}

	cmd.PersistentFlags().StringVarP(&types.CfgFile, "config","c","",
		"配置文件路程，默认当前目录下 registry-manager.yml文件")
	flags := cmd.PersistentFlags()
	flags.ParseErrorsWhitelist.UnknownFlags = true
	flags.Parse(args)

	cmd.AddCommand(
		newDeleteCmd(out),
		newTopCmd(out))
	return cmd, nil
}