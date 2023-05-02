package data

import (
	"github.com/p3rdy/bgpemu/data"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	dataCmd := &cobra.Command{
		Use:   "data",
		Short: "Data commands.",
	}
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start collecting data.",
		RunE:  startFn,
	}
	dumpCmd := &cobra.Command{
		Use:   "dump",
		Short: "Dump routing data.",
		RunE:  dumpFn,
	}
	dataCmd.AddCommand(startCmd)
	dataCmd.AddCommand(dumpCmd)
	return dataCmd
}

func startFn(cmd *cobra.Command, args []string) error {
	return data.Start(args[0])
}

func dumpFn(cmd *cobra.Command, args []string) error {
	return nil
}
