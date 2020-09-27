package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/hzliangbin/harbor-cli/pkg/types"
	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"
)

func main() {

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//flag.CommandLine.Parse([]string{})
	flag.Set("logtostderr","true")

	cmd, err := newRootCmd(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	cobra.OnInitialize(initConfig)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if types.CfgFile != "" {
		viper.SetConfigFile(types.CfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("registry-manager")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		glog.Infof("use config file: %s", viper.ConfigFileUsed())
	}
	err := viper.Unmarshal(&types.Manager)
	if err != nil {
		panic(err)
	}
	types.Manager.Init()
}