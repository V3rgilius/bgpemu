package routes

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/v3rgilius/bgpemu/lab"
	// log "github.com/sirupsen/logrus"
)

func New() *cobra.Command {
	deployCmd := &cobra.Command{
		Use:     "deploy <routes file>",
		Short:   "Deploy bgp route info to routers on cluster",
		RunE:    deployFn,
		PreRunE: ValidateRoutes,
	}
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate route info from AS data",
		RunE:  generateFn,
	}
	routesCmd := &cobra.Command{
		Use:   "routes",
		Short: "Routes commands.",
	}
	routesCmd.AddCommand(deployCmd)
	routesCmd.AddCommand(generateCmd)
	return routesCmd
}

func ValidateRoutes(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: routes info file must be provided", cmd.Use)
	}
	return nil
}

func deployFn(cmd *cobra.Command, args []string) error {
	rts, err := lab.LoadRoutes(args[0])
	if err != nil {
		return err
	}
	m, err := lab.New(rts.TopoName)
	if err != nil {
		return err
	}
	err = m.GetGrpcServersAll()
	if err != nil {
		return err
	}
	err = m.DeployRoutes(rts)
	if err != nil {
		return err
	}
	return nil
}
func generateFn(cmd *cobra.Command, args []string) error {
	return nil
}
