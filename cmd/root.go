/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	knecmd "github.com/openconfig/kne/cmd"
	"github.com/p3rdy/bgpemu/cmd/lab"
	"github.com/p3rdy/bgpemu/cmd/topo"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bgpemu",
	Short: "A tool for BGP emulation",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(topo.New())
	rootCmd.AddCommand(lab.New())
	rootCmd.AddCommand(knecmd.RootCmd)
}
