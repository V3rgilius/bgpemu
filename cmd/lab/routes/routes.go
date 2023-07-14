package routes

import (
	"fmt"

	"github.com/v3rgilius/bgpemu/helper"
	"github.com/v3rgilius/bgpemu/lab"
	"github.com/spf13/cobra"
	// log "github.com/sirupsen/logrus"
)

func New() *cobra.Command {
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy bgp route info to routers on cluster",
		RunE:  deployFn,
	}
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate route info from AS data",
		RunE:  generateFn,
	}
	routesCmd := &cobra.Command{
		Use:   "routes",
		Short: "Topology commands.",
	}
	routesCmd.AddCommand(deployCmd)
	routesCmd.AddCommand(generateCmd)
	return routesCmd
}

func deployFn(cmd *cobra.Command, args []string) error {
	_, err := helper.FileRelative(args[0])
	if err != nil {
		return err
	}
	_, err = lab.LoadRoutes(args[0])

	if err != nil {
		return fmt.Errorf("%s: %w", cmd.Use, err)
	}

	// 加载路由
	// 创建Manager
	// Deploy

	// topopb, err := topo.LoadToKneTopo(args[0])

	// tm, err := knetopo.New(topopb, knetopo.WithKubecfg(helper.DefaultKubeCfg()), knetopo.WithBasePath(bp))
	// if err != nil {
	// 	return fmt.Errorf("%s: %w", cmd.Use, err)
	// }
	// return tm.Create(cmd.Context(), 0)
	return nil
}
func generateFn(cmd *cobra.Command, args []string) error {
	return nil
}
