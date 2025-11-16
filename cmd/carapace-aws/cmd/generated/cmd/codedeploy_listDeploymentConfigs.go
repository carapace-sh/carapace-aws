package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listDeploymentConfigsCmd = &cobra.Command{
	Use:   "list-deployment-configs",
	Short: "Lists the deployment configurations with the user or Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listDeploymentConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_listDeploymentConfigsCmd).Standalone()

		codedeploy_listDeploymentConfigsCmd.Flags().String("next-token", "", "An identifier returned from the previous `ListDeploymentConfigs` call.")
	})
	codedeployCmd.AddCommand(codedeploy_listDeploymentConfigsCmd)
}
