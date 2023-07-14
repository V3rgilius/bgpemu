/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/v3rgilius/bgpemu/cmd/data"
	"github.com/v3rgilius/bgpemu/cmd/lab"
	"github.com/v3rgilius/bgpemu/cmd/topo"
	"os"
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
	rootCmd.AddCommand(data.New())
}
