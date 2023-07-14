package topo

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/v3rgilius/bgpemu/helper"
	"github.com/v3rgilius/bgpemu/topo"
)

func New() *cobra.Command {
	createCmd := &cobra.Command{
		Use:       "create",
		Short:     "Create topology on cluster",
		PreRunE:   ValidateTopology,
		RunE:      createFn,
		ValidArgs: []string{"topology"},
	}
	generateCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate topology from AS data",
		RunE:  generateFn,
	}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update topology on cluster",
		RunE:  updateFn,
	}
	topoCmd := &cobra.Command{
		Use:   "topo",
		Short: "Topology commands.",
	}
	topoCmd.AddCommand(createCmd)
	topoCmd.AddCommand(generateCmd)
	topoCmd.AddCommand(updateCmd)
	return topoCmd
}

func ValidateTopology(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: topology must be provided", cmd.Use)
	}
	return nil
}
func createFn(cmd *cobra.Command, args []string) error {
	bp, err := helper.FileRelative(args[0])
	if err != nil {
		return err
	}
	log.Infof(bp)
	t, err := topo.Load(args[0])
	if err != nil {
		return fmt.Errorf("%s: %w", cmd.Use, err)
	}
	kt, err := topo.KneTopo(t)
	if err != nil {
		return fmt.Errorf("%s: %w", cmd.Use, err)
	}
	tm, err := topo.New(t, kt, 0)
	if err != nil {
		return fmt.Errorf("%s: %w", cmd.Use, err)
	}
	err = tm.Create(cmd.Context(), 0)
	// if err != nil {
	// 	return fmt.Errorf("%s: %w", cmd.Use, err)
	// }
	// err = topo.UpdatePods(t, tm)
	return err
}

func generateFn(cmd *cobra.Command, args []string) error {
	bp, err := helper.FileRelative(args[0])
	if err != nil {
		return err
	}
	bop, err := helper.FileRelative(args[1])
	if err != nil {
		return err
	}
	_, err = topo.GenerateFromAS(bp, bop)
	return err
}

func updateFn(cmd *cobra.Command, args []string) error {
	topopb, err := topo.Load(args[0])
	if err != nil {
		return fmt.Errorf("%s: %w", cmd.Use, err)
	}
	err = topo.Update(topopb)
	if err != nil {
		return fmt.Errorf("%s: %w", cmd.Use, err)
	}
	return err
}
