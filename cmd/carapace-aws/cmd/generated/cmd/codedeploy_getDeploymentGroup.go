package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_getDeploymentGroupCmd = &cobra.Command{
	Use:   "get-deployment-group",
	Short: "Gets information about a deployment group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_getDeploymentGroupCmd).Standalone()

	codedeploy_getDeploymentGroupCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
	codedeploy_getDeploymentGroupCmd.Flags().String("deployment-group-name", "", "The name of a deployment group for the specified application.")
	codedeploy_getDeploymentGroupCmd.MarkFlagRequired("application-name")
	codedeploy_getDeploymentGroupCmd.MarkFlagRequired("deployment-group-name")
	codedeployCmd.AddCommand(codedeploy_getDeploymentGroupCmd)
}
