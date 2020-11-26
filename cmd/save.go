package main

import (
	"fmt"
	. "github.com/hzliangbin/harbor-cli/pkg/types"
	"github.com/hzliangbin/harbor-cli/pkg/utils"
	"github.com/spf13/cobra"
	"io"
)
var OutputTar string
var ImagesList string
var registry string
func newSaveCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "save",
		Short:  "save images",
		Long: "save images from harbor to tar.gz",
		Run: func(cmd *cobra.Command, args []string) {
			if registry != nil {

			}
			harborAddress := Manager.HarborAddress()
			fmt.Fprint(out, harborAddress)

			output, err := utils.ExecuteCommand("docker","ps","-a")
			if err != nil {
				utils.Error(cmd, args, err)
			}

			fmt.Fprint(out, output)
		},
	}
	f := cmd.Flags()
	f.StringVarP(&OutputTar, "output", "o", "images.tar.gz","save images to this file")
	f.StringVarP(&ImagesList, "imageslist", "l", "imageslist.txt","images to be save, one images per line")
	f.StringVarP(&registry, "registry", "r","","registry that pull images from")

	return cmd
}
