package topo

import (
	"fmt"

	knetopo "github.com/openconfig/kne/topo"
	"github.com/p3rdy/bgpemu/helper"
	"github.com/p3rdy/bgpemu/topo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	createCmd := &cobra.Command{
		Use:       "create",
		Short:     "Create Topology on cluster",
		PreRunE:   ValidateTopology,
		RunE:      createFn,
		ValidArgs: []string{"topology"},
	}
	generateCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate topology from AS data",
		RunE:  generateFn,
	}
	topoCmd := &cobra.Command{
		Use:   "topo",
		Short: "Topology commands.",
	}
	topoCmd.AddCommand(createCmd)
	topoCmd.AddCommand(generateCmd)
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
	topopb, err := topo.LoadToKneTopo(args[0])
	if err != nil {
		return fmt.Errorf("%s: %w", cmd.Use, err)
	}

	tm, err := knetopo.New(topopb, knetopo.WithKubecfg(helper.DefaultKubeCfg()), knetopo.WithBasePath(bp))
	if err != nil {
		return fmt.Errorf("%s: %w", cmd.Use, err)
	}
	return tm.Create(cmd.Context(), 0)
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
