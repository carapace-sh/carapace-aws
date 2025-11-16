package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_getDeploymentConfigCmd = &cobra.Command{
	Use:   "get-deployment-config",
	Short: "Gets information about a deployment configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_getDeploymentConfigCmd).Standalone()

	codedeploy_getDeploymentConfigCmd.Flags().String("deployment-config-name", "", "The name of a deployment configuration associated with the user or Amazon Web Services account.")
	codedeploy_getDeploymentConfigCmd.MarkFlagRequired("deployment-config-name")
	codedeployCmd.AddCommand(codedeploy_getDeploymentConfigCmd)
}
