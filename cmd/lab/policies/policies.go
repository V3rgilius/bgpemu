package policies

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/v3rgilius/bgpemu/lab"
)

func New() *cobra.Command {
	deployCmd := &cobra.Command{
		Use:     "deploy <policies file>",
		Short:   "Deploy policies to routers on cluster",
		RunE:    deployFn,
		PreRunE: ValidatePolicies,
	}
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate policies from AS data",
		RunE:  generateFn,
	}
	policiesCmd := &cobra.Command{
		Use:   "policies",
		Short: "Policies commands.",
	}
	policiesCmd.AddCommand(deployCmd)
	policiesCmd.AddCommand(generateCmd)
	return policiesCmd
}

func ValidatePolicies(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: policies file must be provided", cmd.Use)
	}
	return nil
}

func deployFn(cmd *cobra.Command, args []string) error {
	pds, err := lab.LoadPolicies(args[0])
	if err != nil {
		return err
	}
	m, err := lab.New(pds.TopoName)
	if err != nil {
		return err
	}
	err = m.GetGrpcServersAll()
	if err != nil {
		return err
	}
	err = m.DeployPolicies(pds)
	if err != nil {
		return err
	}
	return nil
}
func generateFn(cmd *cobra.Command, args []string) error {
	return nil
}
