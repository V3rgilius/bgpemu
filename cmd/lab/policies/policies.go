package policies

import "github.com/spf13/cobra"

func New() *cobra.Command {
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Create Topology on cluster",
	}
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate topology from AS data",
	}
	policiesCmd := &cobra.Command{
		Use:   "topo",
		Short: "Topology commands.",
	}
	policiesCmd.AddCommand(deployCmd)
	policiesCmd.AddCommand(generateCmd)
	return policiesCmd
}
