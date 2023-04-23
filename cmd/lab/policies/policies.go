package policies

import "github.com/spf13/cobra"

func New() *cobra.Command {
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy policies to routers on cluster",
		RunE:  deployFn,
	}
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate policies from AS data",
		RunE:  generateFn,
	}
	policiesCmd := &cobra.Command{
		Use:   "topo",
		Short: "Topology commands.",
	}
	policiesCmd.AddCommand(deployCmd)
	policiesCmd.AddCommand(generateCmd)
	return policiesCmd
}

func deployFn(cmd *cobra.Command, args []string) error {
	return nil
}
func generateFn(cmd *cobra.Command, args []string) error {
	return nil
}
