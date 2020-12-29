package main

import (
	"bufio"
	"fmt"
	. "github.com/hzliangbin/harbor-cli/pkg/types"
	"github.com/hzliangbin/harbor-cli/pkg/utils"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
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
			var harborAddress string
			if len(registry) != 0 {
				harborAddress = registry
			} else {
				harborAddress = Manager.HarborAddress()
			}
			//fmt.Fprint(out, harborAddress)
			if strings.HasPrefix(harborAddress, "https") {
				harborAddress = strings.Replace(harborAddress,"https://","",1)
			}
			if strings.HasPrefix(harborAddress,"http") {
				harborAddress = strings.Replace(harborAddress, "http://","",1)
			}

			fi, err := os.Open(ImagesList)
			if err != nil {
				utils.Error(cmd, args, err)
			}
			defer fi.Close()

			images :=[]string{}
			br := bufio.NewReader(fi)
			for {
				a, _, c := br.ReadLine()
				if c == io.EOF {
					break
				}
				images = append(images, harborAddress+string(a))

			}
			images = append(images,"| gzip -c > "+ OutputTar)
			fmt.Println(images)
			output, err := utils.ExecuteCommand("docker","save",images...)
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
