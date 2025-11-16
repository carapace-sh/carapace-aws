package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_deleteDeploymentConfigCmd = &cobra.Command{
	Use:   "delete-deployment-config",
	Short: "Deletes a deployment configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_deleteDeploymentConfigCmd).Standalone()

	codedeploy_deleteDeploymentConfigCmd.Flags().String("deployment-config-name", "", "The name of a deployment configuration associated with the user or Amazon Web Services account.")
	codedeploy_deleteDeploymentConfigCmd.MarkFlagRequired("deployment-config-name")
	codedeployCmd.AddCommand(codedeploy_deleteDeploymentConfigCmd)
}
