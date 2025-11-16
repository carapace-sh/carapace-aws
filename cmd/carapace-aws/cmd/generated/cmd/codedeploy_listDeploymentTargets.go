package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listDeploymentTargetsCmd = &cobra.Command{
	Use:   "list-deployment-targets",
	Short: "Returns an array of target IDs that are associated a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listDeploymentTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_listDeploymentTargetsCmd).Standalone()

		codedeploy_listDeploymentTargetsCmd.Flags().String("deployment-id", "", "The unique ID of a deployment.")
		codedeploy_listDeploymentTargetsCmd.Flags().String("next-token", "", "A token identifier returned from the previous `ListDeploymentTargets` call.")
		codedeploy_listDeploymentTargetsCmd.Flags().String("target-filters", "", "A key used to filter the returned targets.")
		codedeploy_listDeploymentTargetsCmd.MarkFlagRequired("deployment-id")
	})
	codedeployCmd.AddCommand(codedeploy_listDeploymentTargetsCmd)
}
