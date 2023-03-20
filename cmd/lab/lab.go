package lab

import (
	"github.com/p3rdy/bgpemu/cmd/lab/policies"
	"github.com/p3rdy/bgpemu/cmd/lab/routes"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	labCmd := &cobra.Command{
		Use:   "topo",
		Short: "Topology commands.",
	}
	labCmd.AddCommand(policies.New())
	labCmd.AddCommand(routes.New())
	return labCmd
}
