package lab

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/v3rgilius/bgpemu/cmd/lab/policies"
	"github.com/v3rgilius/bgpemu/cmd/lab/routes"
	"github.com/v3rgilius/bgpemu/lab"
	"os/exec"
)

func New() *cobra.Command {
	labCmd := &cobra.Command{
		Use:   "lab",
		Short: "Lab commands.",
	}
	deployCmd := &cobra.Command{
		Use:     "deploy <scene file>",
		Short:   "Deploy a lab",
		RunE:    deployFn,
		PreRunE: ValidateScene,
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

func ValidateScene(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: scene file must be provided", cmd.Use)
	}
	return nil
}

func deployFn(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: Scene file must be provided", cmd.Use)
	}
	spb, err := lab.LoadScene(args[0])
	if err != nil {
		return err
	}
	err = lab.DeployScene(spb)
	if err != nil {
		return err
	}
	return nil
}

func generateFn(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No option specified: %s", cmd.Use)
	}
	genArgs := make([]string, 0, 16)
	genArgs = append(genArgs, "py/gen_policy.py")
	for _, arg := range args {
		genArgs = append(genArgs, "--"+arg)
	}
	pycmd := exec.Command("python3", genArgs...)
	out, err := pycmd.Output() // 执行命令，并获取输出和错误信息
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	// switch args[0] {
	// case "init":

	// }
	return nil
}
