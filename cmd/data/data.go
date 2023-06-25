package data

import (
	"fmt"

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
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop collecting data.",
		RunE:  stopFn,
	}
	dataCmd.AddCommand(startCmd)
	dataCmd.AddCommand(stopCmd)
	dataCmd.AddCommand(dumpCmd)
	return dataCmd
}

func startFn(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: topology must be provided", cmd.Use)
	}
	return data.Start(args[0])
}

func dumpFn(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: topology must be provided", cmd.Use)
	}
	return data.Dump(args[0])
}

func stopFn(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: topology must be provided", cmd.Use)
	}
	return data.Stop(args[0])
}
