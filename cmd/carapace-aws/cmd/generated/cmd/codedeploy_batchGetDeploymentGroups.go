package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_batchGetDeploymentGroupsCmd = &cobra.Command{
	Use:   "batch-get-deployment-groups",
	Short: "Gets information about one or more deployment groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_batchGetDeploymentGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_batchGetDeploymentGroupsCmd).Standalone()

		codedeploy_batchGetDeploymentGroupsCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the applicable user or Amazon Web Services account.")
		codedeploy_batchGetDeploymentGroupsCmd.Flags().String("deployment-group-names", "", "The names of the deployment groups.")
		codedeploy_batchGetDeploymentGroupsCmd.MarkFlagRequired("application-name")
		codedeploy_batchGetDeploymentGroupsCmd.MarkFlagRequired("deployment-group-names")
	})
	codedeployCmd.AddCommand(codedeploy_batchGetDeploymentGroupsCmd)
}
