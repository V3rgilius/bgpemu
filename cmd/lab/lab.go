package lab

import (
	"github.com/p3rdy/bgpemu/cmd/lab/policies"
	"github.com/p3rdy/bgpemu/cmd/lab/routes"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	labCmd := &cobra.Command{
		Use:   "lab",
		Short: "Lab commands.",
	}
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a lab",
		RunE:  deployFn,
	}
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a lab scene from a topo file",
		RunE:  generateFn,
	}
	labCmd.AddCommand(policies.New())
	labCmd.AddCommand(routes.New())
	labCmd.AddCommand(deployCmd)
	labCmd.AddCommand(generateCmd)
	return labCmd
}

func deployFn(cmd *cobra.Command, args []string) error {

	return nil
}

func generateFn(cmd *cobra.Command, args []string) error {
	return nil
}
