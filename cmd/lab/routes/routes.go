package routes

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
	routesCmd := &cobra.Command{
		Use:   "topo",
		Short: "Topology commands.",
	}
	routesCmd.AddCommand(deployCmd)
	routesCmd.AddCommand(generateCmd)
	return routesCmd
}
