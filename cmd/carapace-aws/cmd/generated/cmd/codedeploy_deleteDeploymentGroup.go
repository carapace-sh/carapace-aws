package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_deleteDeploymentGroupCmd = &cobra.Command{
	Use:   "delete-deployment-group",
	Short: "Deletes a deployment group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_deleteDeploymentGroupCmd).Standalone()

	codedeploy_deleteDeploymentGroupCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
	codedeploy_deleteDeploymentGroupCmd.Flags().String("deployment-group-name", "", "The name of a deployment group for the specified application.")
	codedeploy_deleteDeploymentGroupCmd.MarkFlagRequired("application-name")
	codedeploy_deleteDeploymentGroupCmd.MarkFlagRequired("deployment-group-name")
	codedeployCmd.AddCommand(codedeploy_deleteDeploymentGroupCmd)
}
