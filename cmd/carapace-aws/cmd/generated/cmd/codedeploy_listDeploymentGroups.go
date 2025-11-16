package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listDeploymentGroupsCmd = &cobra.Command{
	Use:   "list-deployment-groups",
	Short: "Lists the deployment groups for an application registered with the Amazon Web Services user or Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listDeploymentGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_listDeploymentGroupsCmd).Standalone()

		codedeploy_listDeploymentGroupsCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
		codedeploy_listDeploymentGroupsCmd.Flags().String("next-token", "", "An identifier returned from the previous list deployment groups call.")
		codedeploy_listDeploymentGroupsCmd.MarkFlagRequired("application-name")
	})
	codedeployCmd.AddCommand(codedeploy_listDeploymentGroupsCmd)
}
